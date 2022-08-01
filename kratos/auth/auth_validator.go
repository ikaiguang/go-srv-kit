package authutil

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"

	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// DefaultValidateFunc 默认验证方法
func DefaultValidateFunc() ValidateFunc {
	return func(ctx context.Context, token *jwt.Token) error {
		// 转换Claim
		myClaims, ok := token.Claims.(*Claims)
		if !ok || myClaims.Payload == nil || myClaims.Payload.Key == "" {
			logutil.WarnwWithContext(ctx,
				"error", ErrInvalidRedisKey,
				"token.Claims.(*Claims):OK", ok,
				"token.Claims.(*Claims):Content", fmt.Sprintf("%#v", myClaims),
			)
			err := errorutil.WithStack(ErrInvalidKeyFunc)
			return err
		}

		// 验证信息
		authInfo, ok := FromRedisContext(ctx)
		if !ok || authInfo.Payload == nil {
			err := errorutil.WithStack(ErrInvalidAuthInfo)
			return err
		}
		// 无限制
		switch authInfo.Payload.Limit {
		case authv1.LimitTypeEnum_UNKNOWN, authv1.LimitTypeEnum_UNLIMITED:
			return nil
		case authv1.LimitTypeEnum_ONLY_ONE:
			// 仅一个
			if myClaims.Payload.Time.AsTime().UnixNano() != authInfo.Payload.Time.AsTime().UnixNano() {
				err := errorutil.WithStack(ErrLoginLimit)
				return err
			}
		case authv1.LimitTypeEnum_SAME_PLATFORM:
			// 平台限制一个
			if myClaims.Payload.Platform == authInfo.Payload.Platform &&
				myClaims.Payload.Time.AsTime().UnixNano() != authInfo.Payload.Time.AsTime().UnixNano() {
				err := errorutil.WithStack(ErrLoginLimit)
				return err
			}
		}
		return nil
	}
}
