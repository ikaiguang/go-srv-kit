package authpkg

import (
	"context"

	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/locker"
)

// TokenManager 令牌管理器接口。
type TokenManager interface {
	SaveAccessTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	ResetPreviousTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	AddBlacklist(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	AddLoginLimit(ctx context.Context, tokenItems []*TokenItem) error

	GetToken(ctx context.Context, userIdentifier string, tokenID string) (item *TokenItem, isNotFound bool, err error)
	GetAllTokens(ctx context.Context, userIdentifier string) (map[string]*TokenItem, error)

	DeleteTokens(ctx context.Context, userIdentifier string, tokenItems []*TokenItem) error
	DeleteExpireTokens(ctx context.Context, userIdentifier string) error

	IsLoginLimit(ctx context.Context, tokenID string) (bool, LoginLimitEnum_LoginLimit, error)
	IsExistToken(ctx context.Context, userIdentifier string, tokenID string) (bool, error)
	IsBlacklist(ctx context.Context, tokenID string) (bool, error)

	// EasyLock 简单锁，等待解锁或者锁定时间过期后自动解锁。
	EasyLock(ctx context.Context, lockName string) (lockerpkg.Unlocker, error)
}
