package authpkg

import (
	"context"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

const (
	AccessTokenExpire   = time.Hour * 24 * 2
	RefreshTokenExpire  = time.Hour * 24 * 7
	PreviousTokenExpire = time.Minute

	AuthorizationKey = "Authorization"
	BearerWord       = "Bearer"
	BearerFormat     = "Bearer %s"

	PayloadIdentifierPrefixDefault  = "default_"
	PayloadIdentifierPrefixUser     = "user_"
	PayloadIdentifierPrefixAdmin    = "admin_"
	PayloadIdentifierPrefixEmployee = "employee_"
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
	case TokenTypeEnum_EMPLOYEE:
		identifier = PayloadIdentifierPrefixEmployee + identifier
	}
	return identifier
}

// DefaultExpireTime 令牌过期时间
func DefaultExpireTime() *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(AccessTokenExpire))
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

// GenExpireAt ...
func GenExpireAt(duration time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(duration))
}

// GenAuthClaimsByAuthPayload ...
func GenAuthClaimsByAuthPayload(payload *Payload, accessTokenExpire time.Duration) *Claims {
	if payload.TokenID == "" {
		payload.TokenID = uuidpkg.NewUUID()
	}
	authClaims := &Claims{Payload: payload}
	authClaims.ID = payload.TokenID
	authClaims.ExpiresAt = GenExpireAt(AccessTokenExpire)
	if accessTokenExpire > 0 {
		authClaims.ExpiresAt = GenExpireAt(accessTokenExpire)
	}
	return authClaims
}

// GenAuthClaimsByAuthClaims ...
func GenAuthClaimsByAuthClaims(authClaims *Claims, accessTokenExpire time.Duration) *Claims {
	payload := *authClaims.Payload
	payload.TokenID = uuidpkg.NewUUID()
	regClaims := authClaims.RegisteredClaims
	regClaims.ID = payload.TokenID
	regClaims.ExpiresAt = GenExpireAt(AccessTokenExpire)
	if accessTokenExpire > 0 {
		regClaims.ExpiresAt = GenExpireAt(accessTokenExpire)
	}
	return &Claims{
		RegisteredClaims: regClaims,
		Payload:          &payload,
	}
}

// GenRefreshClaimsByAuthClaims ...
func GenRefreshClaimsByAuthClaims(authClaims *Claims, refreshTokenExpire time.Duration) *Claims {
	payload := *authClaims.Payload
	payload.TokenID = uuidpkg.NewUUID()
	regClaims := authClaims.RegisteredClaims
	regClaims.ID = payload.TokenID
	regClaims.ExpiresAt = GenExpireAt(RefreshTokenExpire)
	if refreshTokenExpire > 0 {
		regClaims.ExpiresAt = GenExpireAt(refreshTokenExpire)
	}
	return &Claims{
		RegisteredClaims: regClaims,
		Payload:          &payload,
	}
}

// CheckAndCorrectAuthClaims ...
func CheckAndCorrectAuthClaims(authClaims *Claims) {
	if authClaims.Payload.TokenID == "" {
		authClaims.Payload.TokenID = uuidpkg.NewUUID()
	}
	authClaims.ID = authClaims.Payload.TokenID
	if authClaims.ExpiresAt == nil {
		authClaims.ExpiresAt = DefaultExpireTime()
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

func (s *TokenItem) ID() string {
	if s.IsRefreshToken {
		return s.RefreshTokenID
	}
	return s.TokenID
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
			if transporter, ok := transport.FromServerContext(ctx); ok {
				jwtToken := transporter.RequestHeader().Get(AuthorizationKey)
				tokenInfo, err := validateAuthorizationToken(ctx, jwtToken, signKeyFunc, o)
				if err != nil {
					return nil, err
				}

				// values
				ctx = PutAuthClaimsIntoContext(ctx, tokenInfo.Claims)
				return handler(ctx, req)
			}
			e := ErrWrongContext()
			return nil, errorpkg.WithStack(e)
		}
	}
}

func validateAuthorizationToken(
	ctx context.Context,
	tokenStr string,
	signKeyFunc KeyFunc,
	o *options,
) (*jwt.Token, error) {
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
	if auths := strings.SplitN(tokenStr, " ", 2); len(auths) == 2 && strings.EqualFold(auths[0], BearerWord) {
		tokenStr = auths[1]
	}
	if tokenStr == "" {
		e := ErrMissingToken()
		return nil, errorpkg.WithStack(e)
	}
	var (
		tokenInfo *jwt.Token
		err       error
	)
	if o.claims != nil {
		tokenInfo, err = jwt.ParseWithClaims(tokenStr, o.claims(), keyFunc)
	} else {
		tokenInfo, err = jwt.Parse(tokenStr, keyFunc)
	}
	if err != nil {
		if stderrors.Is(err, jwt.ErrTokenMalformed) || stderrors.Is(err, jwt.ErrTokenUnverifiable) {
			e := ErrInvalidAuthToken()
			return nil, errorpkg.WithStack(e)
		}
		if stderrors.Is(err, jwt.ErrTokenNotValidYet) || stderrors.Is(err, jwt.ErrTokenExpired) {
			e := ErrTokenExpired()
			return nil, errorpkg.WithStack(e)
		}
		e := ErrInvalidAuthToken()
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
	if len(o.accessTokenValidators) > 0 {
		authClaims, ok := tokenInfo.Claims.(*Claims)
		if !ok {
			e := ErrTokenInvalid()
			return nil, errorpkg.WithStack(e)
		}
		for atvIndex := range o.accessTokenValidators {
			if err = o.accessTokenValidators[atvIndex](ctx, authClaims); err != nil {
				return nil, err
			}
		}
	}
	return tokenInfo, nil
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
			if o.accessTokenHeader != nil {
				for k, v := range o.accessTokenHeader {
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
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				clientContext.RequestHeader().Set(AuthorizationKey, fmt.Sprintf(BearerFormat, tokenStr))
				//clientContext.RequestHeader().Set(AuthorizationKey, tokenStr)
				return handler(ctx, req)
			}
			e := ErrWrongContext()
			return nil, errorpkg.WithStack(e)
		}
	}
}
