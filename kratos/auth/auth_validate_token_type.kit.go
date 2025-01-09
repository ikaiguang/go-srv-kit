package authpkg

import (
	"context"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_multiTokenTypes = make(map[TokenTypeEnum_TokenType]bool)
)

func RegisterMultiTokenType(t TokenTypeEnum_TokenType) {
	_multiTokenTypes[t] = true
}

func MustInTokenTypes(_ context.Context, claims *Claims) error {
	in, ok := _multiTokenTypes[claims.Payload.TokenType]
	if !ok || !in {
		e := ErrInvalidClaims()
		return errorpkg.WithStack(e)
	}
	return nil
}

func MustAdminToken(_ context.Context, claims *Claims) error {
	if claims.Payload.TokenType != TokenTypeEnum_ADMIN {
		e := ErrInvalidClaims()
		return errorpkg.WithStack(e)
	}
	return nil
}

func MustUserToken(_ context.Context, claims *Claims) error {
	if claims.Payload.TokenType != TokenTypeEnum_USER {
		e := ErrInvalidClaims()
		return errorpkg.WithStack(e)
	}
	return nil
}

func MustEmployeeToken(_ context.Context, claims *Claims) error {
	if claims.Payload.TokenType != TokenTypeEnum_EMPLOYEE {
		e := ErrInvalidClaims()
		return errorpkg.WithStack(e)
	}
	return nil
}
