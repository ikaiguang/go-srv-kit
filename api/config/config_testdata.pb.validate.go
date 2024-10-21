// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/config/config_testdata.proto

package configpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TestingConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TestingConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestingConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TestingConfigMultiError, or
// nil if none found.
func (m *TestingConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *TestingConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTestdata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestingConfigValidationError{
					field:  "Testdata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestingConfigValidationError{
					field:  "Testdata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestdata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestingConfigValidationError{
				field:  "Testdata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TestingConfigMultiError(errors)
	}

	return nil
}

// TestingConfigMultiError is an error wrapping multiple validation errors
// returned by TestingConfig.ValidateAll() if the designated constraints
// aren't met.
type TestingConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestingConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestingConfigMultiError) AllErrors() []error { return m }

// TestingConfigValidationError is the validation error returned by
// TestingConfig.Validate if the designated constraints aren't met.
type TestingConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestingConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestingConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestingConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestingConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestingConfigValidationError) ErrorName() string { return "TestingConfigValidationError" }

// Error satisfies the builtin error interface
func (e TestingConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestingConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestingConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestingConfigValidationError{}

// Validate checks the field values on TestingConfig_Testdata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TestingConfig_Testdata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestingConfig_Testdata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TestingConfig_TestdataMultiError, or nil if none found.
func (m *TestingConfig_Testdata) ValidateAll() error {
	return m.validate(true)
}

func (m *TestingConfig_Testdata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if len(errors) > 0 {
		return TestingConfig_TestdataMultiError(errors)
	}

	return nil
}

// TestingConfig_TestdataMultiError is an error wrapping multiple validation
// errors returned by TestingConfig_Testdata.ValidateAll() if the designated
// constraints aren't met.
type TestingConfig_TestdataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestingConfig_TestdataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestingConfig_TestdataMultiError) AllErrors() []error { return m }

// TestingConfig_TestdataValidationError is the validation error returned by
// TestingConfig_Testdata.Validate if the designated constraints aren't met.
type TestingConfig_TestdataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestingConfig_TestdataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestingConfig_TestdataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestingConfig_TestdataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestingConfig_TestdataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestingConfig_TestdataValidationError) ErrorName() string {
	return "TestingConfig_TestdataValidationError"
}

// Error satisfies the builtin error interface
func (e TestingConfig_TestdataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestingConfig_Testdata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestingConfig_TestdataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestingConfig_TestdataValidationError{}
