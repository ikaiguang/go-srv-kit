package filepkg

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanonicalHashNamesAndLegacyWrappers(t *testing.T) {
	const content = "hello"
	const wantMD5 = "5d41402abc4b2a76b9719d911017c592"
	const wantSHA256 = "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

	path := t.TempDir() + "/content.txt"
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))

	md5sum, size, err := MD5(path)
	require.NoError(t, err)
	assert.Equal(t, wantMD5, md5sum)
	assert.Equal(t, int64(len(content)), size)

	md5sum, size, err = MD5FromReader(strings.NewReader(content))
	require.NoError(t, err)
	assert.Equal(t, wantMD5, md5sum)
	assert.Equal(t, int64(len(content)), size)

	sha256sum, size, err := SHA256(path)
	require.NoError(t, err)
	assert.Equal(t, wantSHA256, sha256sum)
	assert.Equal(t, int64(len(content)), size)

	sha256sum, size, err = SHA256FromReader(strings.NewReader(content))
	require.NoError(t, err)
	assert.Equal(t, wantSHA256, sha256sum)
	assert.Equal(t, int64(len(content)), size)

	legacy, _, err := Md5(path)
	require.NoError(t, err)
	assert.Equal(t, wantMD5, legacy)
}
