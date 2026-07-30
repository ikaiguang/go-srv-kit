// Package rotatelogs is a port of File-RotateLogs from Perl
// (https://metacpan.org/release/File-RotateLogs), and it allows
// you to automatically rotate output files when you write to them
// according to the filename pattern that you can specify.
package rotatelogs

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	strftime "github.com/lestrrat-go/strftime"
)

func (c clockFn) Now() time.Time {
	return c()
}

// New creates a new RotateLogs object. A log filename pattern
// must be passed. Optional `Option` parameters may be passed
func New(p string, options ...Option) (*RotateLogs, error) {
	if p == "" {
		return nil, errors.New("log filename pattern must not be empty")
	}

	globPattern := p
	for _, re := range patternConversionRegexps {
		globPattern = re.ReplaceAllString(globPattern, "*")
	}

	pattern, err := strftime.New(p)
	if err != nil {
		return nil, fmt.Errorf("invalid strftime pattern: %w", err)
	}

	var clock Clock = Local
	rotationTime := 24 * time.Hour
	var rotationSize int64
	var rotationCount uint
	var linkName string
	var maxAge time.Duration
	var handler Handler
	var forceNewFile bool

	for _, o := range options {
		if isNil(o) {
			continue
		}

		name := o.Name()
		value := o.Value()
		switch name {
		case optkeyClock:
			clockValue, ok := value.(Clock)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			if isNil(clockValue) {
				return nil, errors.New("clock option must not be nil")
			}
			clock = clockValue
		case optkeyLinkName:
			linkValue, ok := value.(string)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			linkName = linkValue
		case optkeyMaxAge:
			maxAgeValue, ok := value.(time.Duration)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			maxAge = maxAgeValue
			if maxAge < 0 {
				maxAge = 0
			}
		case optkeyRotationTime:
			rotationTimeValue, ok := value.(time.Duration)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			rotationTime = rotationTimeValue
			if rotationTime < 0 {
				rotationTime = 0
			}
		case optkeyRotationSize:
			rotationSizeValue, ok := value.(int64)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			rotationSize = rotationSizeValue
			if rotationSize < 0 {
				rotationSize = 0
			}
		case optkeyRotationCount:
			rotationCountValue, ok := value.(uint)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			rotationCount = rotationCountValue
		case optkeyHandler:
			if isNil(value) {
				handler = nil
				continue
			}
			handlerValue, ok := value.(Handler)
			if !ok {
				return nil, invalidOptionValue(name, value)
			}
			handler = handlerValue
		case optkeyForceNewFile:
			forceNewFile = true
		}
	}

	if maxAge > 0 && rotationCount > 0 {
		return nil, errors.New("options MaxAge and RotationCount cannot be both set")
	}

	if maxAge == 0 && rotationCount == 0 {
		// if both are 0, give maxAge a sane default
		maxAge = 7 * 24 * time.Hour
	}

	return &RotateLogs{
		clock:         clock,
		eventHandler:  handler,
		globPattern:   globPattern,
		linkName:      linkName,
		maxAge:        maxAge,
		pattern:       pattern,
		rotationTime:  rotationTime,
		rotationSize:  rotationSize,
		rotationCount: rotationCount,
		forceNewFile:  forceNewFile,
	}, nil
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return reflected.IsNil()
	default:
		return false
	}
}

func invalidOptionValue(name string, value any) error {
	return fmt.Errorf("invalid value of type %T for option %q", value, name)
}

func (rl *RotateLogs) genFilename() string {
	now := rl.clock.Now()

	// XXX HACK: Truncate only happens in UTC semantics, apparently.
	// observed values for truncating given time with 86400 secs:
	//
	// before truncation: 2018/06/01 03:54:54 2018-06-01T03:18:00+09:00
	// after  truncation: 2018/06/01 03:54:54 2018-05-31T09:00:00+09:00
	//
	// This is really annoying when we want to truncate in local time
	// so we hack: we take the apparent local time in the local zone,
	// and pretend that it's in UTC. do our math, and put it back to
	// the local zone
	var base time.Time
	if now.Location() != time.UTC {
		base = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
		base = base.Truncate(time.Duration(rl.rotationTime))
		base = time.Date(base.Year(), base.Month(), base.Day(), base.Hour(), base.Minute(), base.Second(), base.Nanosecond(), base.Location())
	} else {
		base = now.Truncate(time.Duration(rl.rotationTime))
	}
	return rl.pattern.FormatString(base)
}

// Write satisfies the io.Writer interface. It writes to the
// appropriate file handle that is currently being used.
// If we have reached rotation time, the target file gets
// automatically rotated, and also purged if necessary.
func (rl *RotateLogs) Write(p []byte) (n int, err error) {
	// Guard against concurrent writes
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	if rl.closed {
		return 0, os.ErrClosed
	}

	out, err := rl.getWriter_nolock(false, false)
	if err != nil {
		return 0, fmt.Errorf("failed to acquire target io.Writer: %w", err)
	}

	return out.Write(p)
}

// must be locked during this operation
func (rl *RotateLogs) getWriter_nolock(bailOnRotateFail, useGenerationalNames bool) (io.Writer, error) {
	generation := rl.generation
	previousFn := rl.curFn
	// This filename contains the name of the "NEW" filename
	// to log to, which may be newer than rl.currentFilename
	baseFn := rl.genFilename()
	filename := baseFn
	var forceNewFile bool

	fi, err := os.Stat(rl.curFn)
	sizeRotation := false
	if err == nil && rl.rotationSize > 0 && rl.rotationSize <= fi.Size() {
		forceNewFile = true
		sizeRotation = true
	}

	if baseFn != rl.curBaseFn {
		generation = 0
		// even though this is the first write after calling New(),
		// check if a new file needs to be created
		if rl.forceNewFile {
			forceNewFile = true
		}
	} else {
		if !useGenerationalNames && !sizeRotation {
			// nothing to do
			return rl.outFh, nil
		}
		forceNewFile = true
		generation++
	}
	if forceNewFile {
		// A new file has been requested. Instead of just using the
		// regular strftime pattern, we create a new file name using
		// generational names such as "foo.1", "foo.2", "foo.3", etc
		var name string
		for {
			if generation == 0 {
				name = filename
			} else {
				name = fmt.Sprintf("%s.%d", filename, generation)
			}
			if _, err := os.Stat(name); err != nil {
				filename = name
				break
			}
			generation++
		}
	}
	// make sure the dir is existed, eg:
	// ./foo/bar/baz/hello.log must make sure ./foo/bar/baz is existed
	dirname := filepath.Dir(filename)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", dirname, err)
	}
	// if we got here, then we need to create a file
	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}

	if err := rl.rotate_nolock(filename); err != nil {
		err = fmt.Errorf("failed to rotate: %w", err)
		if bailOnRotateFail {
			// Failure to rotate is a problem, but it's really not a great
			// idea to stop your application just because you couldn't rename
			// your log.
			//
			// We only return this error when explicitly needed (as specified by bailOnRotateFail)
			//
			// However, we *NEED* to close `fh` here
			if fh != nil { // probably can't happen, but being paranoid
				fh.Close()
			}
			return nil, err
		}
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}

	var closeErr error
	if rl.outFh != nil {
		closeErr = rl.outFh.Close()
	}
	rl.outFh = fh
	rl.curBaseFn = baseFn
	rl.curFn = filename
	rl.generation = generation

	if h := rl.eventHandler; h != nil {
		go h.Handle(&FileRotatedEvent{
			prev:    previousFn,
			current: filename,
		})
	}
	if closeErr != nil {
		err := fmt.Errorf("failed to close previous log file: %w", closeErr)
		if bailOnRotateFail {
			return nil, err
		}
		fmt.Fprintln(os.Stderr, err)
	}
	return fh, nil
}

// CurrentFileName returns the current file name that
// the RotateLogs object is writing to
func (rl *RotateLogs) CurrentFileName() string {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	return rl.curFn
}

var patternConversionRegexps = []*regexp.Regexp{
	regexp.MustCompile(`%[%+A-Za-z]`),
	regexp.MustCompile(`\*+`),
}

// Rotate forcefully rotates the log files. If the generated file name
// clash because file already exists, a numeric suffix of the form
// ".1", ".2", ".3" and so forth are appended to the end of the log file
//
// This method can be used in conjunction with a signal handler to
// emulate servers that generate new log files when they receive a
// SIGHUP
func (rl *RotateLogs) Rotate() error {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	if rl.closed {
		return os.ErrClosed
	}
	if _, err := rl.getWriter_nolock(true, true); err != nil {
		return err
	}
	return nil
}

func (rl *RotateLogs) rotate_nolock(filename string) error {
	lockfn := filename + `_lock`
	fh, err := os.OpenFile(lockfn, os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		// Can't lock, just return
		return err
	}

	defer func() {
		_ = fh.Close()
		_ = os.Remove(lockfn)
	}()

	if rl.linkName != "" {
		// Change how the link name is generated based on where the
		// target location is. if the location is directly underneath
		// the main filename's parent directory, then we create a
		// symlink with a relative path
		linkDest := filename
		linkDir := filepath.Dir(rl.linkName)

		baseDir := filepath.Dir(filename)
		linkDirFromBase, relErr := filepath.Rel(baseDir, linkDir)
		if relErr == nil && linkDirFromBase != ".." && !strings.HasPrefix(linkDirFromBase, ".."+string(os.PathSeparator)) {
			tmp, err := filepath.Rel(linkDir, filename)
			if err != nil {
				return fmt.Errorf("failed to evaluate relative path from %q to %q: %w", baseDir, rl.linkName, err)
			}

			linkDest = tmp
		}

		// the directory where rl.linkName should be created must exist
		_, err := os.Stat(linkDir)
		if err != nil { // Assume err != nil means the directory doesn't exist
			if err := os.MkdirAll(linkDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", linkDir, err)
			}
		}

		// Keep the temporary link in linkDir so the final rename remains on
		// the same filesystem even when logs and the stable link differ.
		tmpLinkName := filepath.Join(linkDir, "."+filepath.Base(rl.linkName)+"_symlink")
		_ = os.Remove(tmpLinkName)
		if err := os.Symlink(linkDest, tmpLinkName); err != nil {
			return fmt.Errorf("failed to create new symlink: %w", err)
		}

		if err := os.Rename(tmpLinkName, rl.linkName); err != nil {
			_ = os.Remove(tmpLinkName)
			return fmt.Errorf("failed to rename new symlink: %w", err)
		}
	}

	if rl.maxAge <= 0 && rl.rotationCount <= 0 {
		return errors.New("maxAge and rotationCount are both disabled")
	}

	matches, err := rl.rotationFiles()
	if err != nil {
		return err
	}

	cutoff := rl.clock.Now().Add(-1 * rl.maxAge)
	type rotationFile struct {
		path    string
		modTime time.Time
	}
	var toUnlink []rotationFile
	for _, path := range matches {
		// Ignore lock files
		if strings.HasSuffix(path, "_lock") || strings.HasSuffix(path, "_symlink") {
			continue
		}

		fi, err := os.Stat(path)
		if err != nil {
			continue
		}

		fl, err := os.Lstat(path)
		if err != nil {
			continue
		}

		if rl.maxAge > 0 && fi.ModTime().After(cutoff) {
			continue
		}

		if rl.rotationCount > 0 && fl.Mode()&os.ModeSymlink == os.ModeSymlink {
			continue
		}
		toUnlink = append(toUnlink, rotationFile{path: path, modTime: fi.ModTime()})
	}
	sort.Slice(toUnlink, func(i, j int) bool {
		if toUnlink[i].modTime.Equal(toUnlink[j].modTime) {
			if toUnlink[i].path == filename {
				return false
			}
			if toUnlink[j].path == filename {
				return true
			}
			return toUnlink[i].path < toUnlink[j].path
		}
		return toUnlink[i].modTime.Before(toUnlink[j].modTime)
	})

	if rl.rotationCount > 0 {
		// Only delete if we have more than rotationCount
		if rl.rotationCount >= uint(len(toUnlink)) {
			return nil
		}

		toUnlink = toUnlink[:len(toUnlink)-int(rl.rotationCount)]
	}

	if len(toUnlink) <= 0 {
		return nil
	}

	var removeErr error
	for _, file := range toUnlink {
		if err := os.Remove(file.path); err != nil && !errors.Is(err, os.ErrNotExist) {
			removeErr = errors.Join(removeErr, fmt.Errorf("remove expired log file %s: %w", file.path, err))
		}
	}
	return removeErr
}

func (rl *RotateLogs) rotationFiles() ([]string, error) {
	baseMatches, err := filepath.Glob(rl.globPattern)
	if err != nil {
		return nil, err
	}

	generationMatches, err := filepath.Glob(rl.globPattern + ".*")
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{}, len(baseMatches)+len(generationMatches))
	matches := make([]string, 0, len(baseMatches)+len(generationMatches))
	for _, path := range baseMatches {
		seen[path] = struct{}{}
		matches = append(matches, path)
	}
	for _, path := range generationMatches {
		if _, ok := seen[path]; ok || !hasGenerationSuffix(path) {
			continue
		}
		seen[path] = struct{}{}
		matches = append(matches, path)
	}
	return matches, nil
}

func hasGenerationSuffix(path string) bool {
	ext := filepath.Ext(path)
	if len(ext) < 2 {
		return false
	}
	generation, err := strconv.ParseUint(ext[1:], 10, 64)
	return err == nil && generation > 0
}

// Close satisfies the io.Closer interface. You must
// call this method if you performed any writes to
// the object.
func (rl *RotateLogs) Close() error {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	if rl.closed {
		return nil
	}
	rl.closed = true

	if rl.outFh == nil {
		return nil
	}

	err := rl.outFh.Close()
	rl.outFh = nil
	return err
}
