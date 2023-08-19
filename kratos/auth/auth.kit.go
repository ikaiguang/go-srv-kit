package authpkg

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

const (
	TokenExpireDuration = time.Hour * 24 * 7
	RefreshTokenExpire  = TokenExpireDuration + time.Hour*24*7

	AuthorizationKey = "Authorization"
	BearerWord       = "Bearer"
	BearerFormat     = "Bearer %s"

	PayloadIdentifierPrefixDefault = "default_"
	PayloadIdentifierPrefixUser    = "user_"
	PayloadIdentifierPrefixAdmin   = "admin_"
)

// Payload 授权信息
type Payload struct {
	// TokenId 令牌唯一id
	TokenID string `json:"ti,omitempty"`
	// uid 用户唯一id
	UserID uint64 `json:"uid,omitempty"`
	// UserUuid 用户唯一id
	UserUuid string `json:"uuid,omitempty"`
	// LoginPlatform 登录平台信息
	LoginPlatform LoginPlatformEnum_LoginPlatform `json:"lp,omitempty"`
	// LoginType 登录类型
	LoginType LoginTypeEnum_LoginType `json:"lt,omitempty"`
	// LoginLimit 登录限制
	LoginLimit LoginLimitEnum_LoginLimit `json:"ll,omitempty"`
	// TokenType 令牌类型
	TokenType TokenTypeEnum_TokenType `json:"tt,omitempty"`
}

// UserIdentifier ...
func (s *Payload) UserIdentifier() string {
	identifier := "0"
	// identifier
	if s.UserUuid != "" {
		identifier = s.UserUuid
	} else if s.UserID > 0 {
		identifier = strconv.FormatUint(s.UserID, 10)
	}
	switch s.TokenType {
	default:
		identifier = PayloadIdentifierPrefixDefault + identifier
	case TokenTypeEnum_USER:
		identifier = PayloadIdentifierPrefixUser + identifier
	case TokenTypeEnum_ADMIN:
		identifier = PayloadIdentifierPrefixAdmin + identifier
	}
	return identifier
}

// DefaultExpireTime 令牌过期时间
func DefaultExpireTime() *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(TokenExpireDuration))
}

// Claims jwt.Claims
type Claims struct {
	jwt.RegisteredClaims

	// payload 授权信息
	Payload *Payload `json:"p,omitempty"`
}

// EncodeToString ...
func (s *Claims) EncodeToString() (string, error) {
	res, err := json.Marshal(s)
	if err != nil {
		e := errorpkg.ErrorBadRequest("encode token claims failed")
		err = errorpkg.Wrap(e, err)
		return "", err
	}
	return string(res), nil
}

// DecodeString ...
func (s *Claims) DecodeString(claimCiphertext string) error {
	err := json.Unmarshal([]byte(claimCiphertext), s)
	if err != nil {
		e := errorpkg.ErrorBadRequest("decode token claims failed")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// DefaultAuthClaims ...
func DefaultAuthClaims(payload Payload) *Claims {
	return &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: DefaultExpireTime(),
			ID:        uuidpkg.NewUUID(),
		},
		Payload: &payload,
	}
}

// DefaultRefreshClaims ...
func DefaultRefreshClaims(authClaims *Claims) *Claims {
	payload := *authClaims.Payload
	regClaims := authClaims.RegisteredClaims
	regClaims.ID = uuidpkg.NewUUID()
	regClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(RefreshTokenExpire))
	return &Claims{
		RegisteredClaims: regClaims,
		Payload:          &payload,
	}
}

// TokenItem 令牌信息
type TokenItem struct {
	TokenID        string `json:"ti,omitempty"`
	RefreshTokenID string `json:"rti,omitempty"`
	ExpiredAt      int64  `json:"ea,omitempty"`
	IsRefreshToken bool   `json:"ift,omitempty"`

	// payload 授权信息
	Payload *Payload `json:"p,omitempty"`
}

// EncodeToString ...
func (s *TokenItem) EncodeToString() (string, error) {
	res, err := json.Marshal(s)
	if err != nil {
		e := errorpkg.ErrorBadRequest("encode token item failed")
		err = errorpkg.Wrap(e, err)
		return "", err
	}
	return string(res), nil
}

// DecodeString ...
func (s *TokenItem) DecodeString(tokenItem string) error {
	err := json.Unmarshal([]byte(tokenItem), s)
	if err != nil {
		e := errorpkg.ErrorBadRequest("decode token item failed")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// contextAuthClaims context.Context key
type contextAuthClaims struct{}

// PutAuthClaimsIntoContext put auth info into context
func PutAuthClaimsIntoContext(ctx context.Context, info jwt.Claims) context.Context {
	return context.WithValue(ctx, contextAuthClaims{}, info)
}

// GetAuthClaimsFromContext extract auth info from context
func GetAuthClaimsFromContext(ctx context.Context) (*Claims, bool) {
	token, ok := ctx.Value(contextAuthClaims{}).(*Claims)
	return token, ok
}

// KeyFunc 自定义 jwt.Keyfunc
type KeyFunc func(context.Context) jwt.Keyfunc

// Server is a server auth middleware. Check the token and extract the info from token.
func Server(signKeyFunc KeyFunc, opts ...Option) middleware.Middleware {
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				var keyFunc jwt.Keyfunc
				if signKeyFunc == nil {
					e := ErrMissingSignKeyFunc()
					return nil, errorpkg.WithStack(e)
				}
				keyFunc = signKeyFunc(ctx)
				if keyFunc == nil {
					e := ErrMissingSignKeyFunc()
					return nil, errorpkg.WithStack(e)
				}
				//auths := strings.SplitN(header.RequestHeader().Get(AuthorizationKey), " ", 2)
				//if len(auths) != 2 || !strings.EqualFold(auths[0], BearerWord) {
				//	e := ErrMissingToken()
				//	return nil, errorpkg.WithStack(e)
				//}
				//jwtToken := auths[1]
				jwtToken := header.RequestHeader().Get(AuthorizationKey)
				if jwtToken == "" {
					e := ErrMissingToken()
					return nil, errorpkg.WithStack(e)
				}
				var (
					tokenInfo *jwt.Token
					err       error
				)
				if o.claims != nil {
					tokenInfo, err = jwt.ParseWithClaims(jwtToken, o.claims(), keyFunc)
				} else {
					tokenInfo, err = jwt.Parse(jwtToken, keyFunc)
				}
				if err != nil {
					ve, ok := err.(*jwt.ValidationError)
					if !ok {
						e := ErrInvalidAuthToken()
						return nil, errorpkg.WithStack(e)
					}
					if ve.Errors&jwt.ValidationErrorMalformed != 0 {
						e := ErrTokenInvalid()
						return nil, errorpkg.WithStack(e)
					}
					if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						e := ErrTokenExpired()
						return nil, errorpkg.WithStack(e)
					}
					e := ErrTokenParseFail()
					e.Metadata = map[string]string{"error": err.Error()}
					return nil, errorpkg.WithStack(e)
				}
				if !tokenInfo.Valid {
					e := ErrTokenInvalid()
					return nil, errorpkg.WithStack(e)
				}
				if tokenInfo.Method != o.signingMethod {
					e := ErrUnSupportSigningMethod()
					return nil, errorpkg.WithStack(e)
				}
				if o.tokenValidatorFunc != nil {
					if err = o.tokenValidatorFunc(ctx, tokenInfo); err != nil {
						return nil, err
					}
				}
				ctx = PutAuthClaimsIntoContext(ctx, tokenInfo.Claims)
				return handler(ctx, req)
			}
			e := ErrWrongContext()
			return nil, errorpkg.WithStack(e)
		}
	}
}

// Client is a client jwt middleware.
func Client(customKeyFunc KeyFunc, opts ...Option) middleware.Middleware {
	claims := jwt.RegisteredClaims{}
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
		claims:        func() jwt.Claims { return claims },
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var keyProvider jwt.Keyfunc
			if customKeyFunc == nil {
				e := ErrMissingSignKeyFunc()
				return nil, errorpkg.WithStack(e)
			}
			keyProvider = customKeyFunc(ctx)
			if keyProvider == nil {
				e := ErrMissingSignKeyFunc()
				return nil, errorpkg.WithStack(e)
			}
			if keyProvider == nil {
				e := ErrNeedTokenProvider()
				return nil, errorpkg.WithStack(e)
			}
			token := jwt.NewWithClaims(o.signingMethod, o.claims())
			if o.tokenHeader != nil {
				for k, v := range o.tokenHeader {
					token.Header[k] = v
				}
			}
			key, err := keyProvider(token)
			if err != nil {
				e := ErrGetKey()
				return nil, errorpkg.WithStack(e)
			}
			tokenStr, err := token.SignedString(key)
			if err != nil {
				e := ErrSignToken()
				return nil, errorpkg.WithStack(e)
			}
			if o.tokenValidatorFunc != nil {
				if err = o.tokenValidatorFunc(ctx, token); err != nil {
					return nil, err
				}
			}
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				//clientContext.RequestHeader().Set(AuthorizationKey, fmt.Sprintf(BearerFormat, tokenStr))
				clientContext.RequestHeader().Set(AuthorizationKey, tokenStr)
				return handler(ctx, req)
			}
			e := ErrWrongContext()
			return nil, errorpkg.WithStack(e)
		}
	}
}
