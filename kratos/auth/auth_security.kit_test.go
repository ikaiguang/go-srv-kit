package authpkg

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/error"
	"github.com/stretchr/testify/require"
)

func TestSignEncryptorRejectsUnexpectedAlgorithm(t *testing.T) {
	const key = "signing-key"
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		Payload: &Payload{TokenID: "token-1"},
	}).SignedString([]byte(key))
	require.NoError(t, err)

	_, err = NewSignEncryptor(key).DecryptToken(context.Background(), token)
	require.Error(t, err)
}

func TestSignEncryptorClassifiesExpiredToken(t *testing.T) {
	const key = "signing-key"
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
		},
		Payload: &Payload{TokenID: "token-1"},
	}).SignedString([]byte(key))
	require.NoError(t, err)

	_, err = NewSignEncryptor(key).DecryptToken(context.Background(), token)
	require.True(t, IsTokenExpired(err))
}

func TestAuthRepoRejectsMissingClaims(t *testing.T) {
	repo, err := NewAuthRepo(Config{
		SignCrypto:    NewSignEncryptor("signing-key"),
		RefreshCrypto: NewCBCCipher("0123456789abcdef"),
	}, slog.Default(), nil)
	require.NoError(t, err)

	_, err = repo.SignToken(context.Background(), nil)
	require.Error(t, err)
	require.Error(t, repo.VerifyAccessToken(context.Background(), &Claims{}))
}

func TestSignEncryptorRejectsEmptyKey(t *testing.T) {
	encryptor := NewSignEncryptor("")
	_, err := encryptor.EncryptToken(context.Background(), &Claims{})
	require.Error(t, err)
}

func TestClaimGeneratorsRejectMissingPayload(t *testing.T) {
	require.Nil(t, GenAuthClaimsByAuthPayload(nil, 0))
	require.Nil(t, GenAuthClaimsByAuthClaims(&Claims{}, 0))
	require.Nil(t, GenRefreshClaimsByAuthClaims(&Claims{}, 0))
}

func TestEncryptorsRejectNilClaims(t *testing.T) {
	_, err := NewSignEncryptor("signing-key").EncryptToken(context.Background(), nil)
	require.Error(t, err)
	_, err = NewCBCCipher("0123456789abcdef").EncryptToken(context.Background(), nil)
	require.Error(t, err)
}

func TestValidateAuthorizationTokenDoesNotExposeParserError(t *testing.T) {
	_, err := validateAuthorizationToken(
		context.Background(),
		"invalid.token.value",
		func(context.Context) jwt.Keyfunc {
			return func(*jwt.Token) (any, error) {
				return []byte("signing-key"), nil
			}
		},
		&options{signingMethod: jwt.SigningMethodHS256},
	)
	require.Error(t, err)
	metadata, _ := errorpkg.Metadata(err)
	require.NotContains(t, metadata, errorpkg.EnumMDErrorKey)
}
