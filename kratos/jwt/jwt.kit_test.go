package jwtutil

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./internal/pkg/jwt -run=TestGenUserToken
func TestGenUserToken(t *testing.T) {
	var (
		authId        uint64 = 0
		authUuid             = ""
		authOpenid           = ""
		backendOpenid        = ""
		apiKey               = "wfWA5G93z0bR8sdonyeQpIElHLqYX4P2"
	)

	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: GenExpireTime(),
		},
		AuthId:        authId,
		AuthUuid:      authUuid,
		AuthOpenId:    authOpenid,
		BackendOpenId: backendOpenid,
	})
	signedString, err := claims.SignedString([]byte(apiKey))
	require.Nil(t, err)
	t.Log("==> signedString : ", signedString)
}
