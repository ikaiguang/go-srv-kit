package authpkg

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	aespkg "github.com/ikaiguang/go-srv-kit/kit/aes"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
)

// Config ...
type Config struct {
	SignCrypto    SignEncryptor
	RefreshCrypto RefreshEncryptor
}

// TokenResponse ...
type TokenResponse struct {
	AccessToken  string
	RefreshToken string

	AccessTokenItem  *TokenItem
	RefreshTokenItem *TokenItem
}

// SignEncryptor ...
type SignEncryptor interface {
	JWTSigningKeyFunc(ctx context.Context) jwt.Keyfunc
	JWTSigningMethod() jwt.SigningMethod
	JWTSigningClaims() jwt.Claims

	EncryptToken(ctx context.Context, authClaims *Claims) (string, error)
	DecryptToken(ctx context.Context, accessToken string) (*Claims, error)
}

type signEncryptor struct {
	key           []byte
	signingMethod *jwt.SigningMethodHMAC
}

// NewSignEncryptor ...
func NewSignEncryptor(key string) SignEncryptor {
	return &signEncryptor{
		key:           []byte(key),
		signingMethod: jwt.SigningMethodHS256,
	}
}

// JWTSigningKeyFunc 密钥 jwt.Keyfunc
func (s *signEncryptor) JWTSigningKeyFunc(ctx context.Context) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return s.key, nil
	}
}

// JWTSigningMethod 签名方法
func (s *signEncryptor) JWTSigningMethod() jwt.SigningMethod {
	return s.signingMethod
}

// JWTSigningClaims 签名载体
func (s *signEncryptor) JWTSigningClaims() jwt.Claims {
	return &Claims{}
}

func (s *signEncryptor) EncryptToken(ctx context.Context, authClaims *Claims) (string, error) {
	token, err := jwt.NewWithClaims(s.signingMethod, authClaims).SignedString(s.key)
	if err != nil {
		e := errorpkg.ErrorBadRequest("sign token failed")
		err = errorpkg.Wrap(e, err)
		return "", err
	}
	return token, nil
}

func (s *signEncryptor) DecryptToken(ctx context.Context, accessToken string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, s.JWTSigningKeyFunc(ctx))
	if err != nil {
		e := errorpkg.ErrorBadRequest("decrypt token failed")
		err = errorpkg.Wrap(e, err)
		return nil, err
	}
	return claims, err
}

// RefreshEncryptor ...
type RefreshEncryptor interface {
	EncryptToken(ctx context.Context, refreshClaims *Claims) (string, error)
	DecryptToken(ctx context.Context, refreshToken string) (*Claims, error)
}

// cbcEncryptor ...
type cbcEncryptor struct{ key []byte }

// NewCBCCipher ...
func NewCBCCipher(key string) RefreshEncryptor {
	return &cbcEncryptor{key: []byte(key)}
}

func (s *cbcEncryptor) EncryptToken(ctx context.Context, refreshClaims *Claims) (string, error) {
	refreshClaimsStr, err := refreshClaims.EncodeToString()
	if err != nil {
		return "", err
	}
	token, err := aespkg.EncryptCBC([]byte(refreshClaimsStr), s.key)
	if err != nil {
		e := errorpkg.ErrorBadRequest("crypto refresh claims failed")
		err = errorpkg.Wrap(e, err)
		return "", err
	}
	return token, nil
}

func (s *cbcEncryptor) DecryptToken(ctx context.Context, refreshToken string) (*Claims, error) {
	plaintext, err := aespkg.DecryptCBC(refreshToken, s.key)
	if err != nil {
		e := errorpkg.ErrorBadRequest("decode refresh token claims failed")
		err = errorpkg.Wrap(e, err)
		return nil, err
	}
	claims := &Claims{}
	err = claims.DecodeString(plaintext)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

var _ AuthRepo = (*authRepo)(nil)

// AuthRepo ...
type AuthRepo interface {
	JWTSigningKeyFunc(ctx context.Context) jwt.Keyfunc
	JWTSigningMethod() jwt.SigningMethod
	JWTSigningClaims() jwt.Claims

	// SignToken 签证Token
	// @Param signKey 拼接在原来的signKey上
	SignToken(ctx context.Context, authClaims *Claims) (*TokenResponse, error)
	DecodeAccessToken(ctx context.Context, accessToken string) (*Claims, error)
	DecodeRefreshToken(ctx context.Context, refreshToken string) (*Claims, error)

	VerifyToken(ctx context.Context, jwtToken *jwt.Token) error
}

// authRepo ...
type authRepo struct {
	logHandler    *log.Helper
	signEncryptor SignEncryptor
	refreshCrypto RefreshEncryptor
	tokenManger   TokenManger
}

// NewAuthRepo ...
func NewAuthRepo(config Config, logger log.Logger, tokenManger TokenManger) (AuthRepo, error) {
	if config.SignCrypto == nil {
		e := errorpkg.ErrorBadRequest("invalid SignCrypto")
		err := errorpkg.WithStack(e)
		return nil, err
	}
	if config.RefreshCrypto == nil {
		e := errorpkg.ErrorBadRequest("invalid RefreshCrypto")
		err := errorpkg.WithStack(e)
		return nil, err
	}
	// authCacheKeyPrefix := CheckAuthCacheKeyPrefix(config.AuthCacheKeyPrefix)
	return &authRepo{
		signEncryptor: config.SignCrypto,
		refreshCrypto: config.RefreshCrypto,
		logHandler:    log.NewHelper(log.With(logger, "module", "auth/repo")),
		tokenManger:   tokenManger,
		// tokenManger:   NewTokenManger(redisCC, authCacheKeyPrefix),
	}, nil
}

// JWTSigningKeyFunc 密钥 jwt.Keyfunc
func (s *authRepo) JWTSigningKeyFunc(ctx context.Context) jwt.Keyfunc {
	return s.signEncryptor.JWTSigningKeyFunc(ctx)
}

// JWTSigningMethod 签名方法
func (s *authRepo) JWTSigningMethod() jwt.SigningMethod {
	return s.signEncryptor.JWTSigningMethod()
}

// JWTSigningClaims 签名载体
func (s *authRepo) JWTSigningClaims() jwt.Claims {
	return s.signEncryptor.JWTSigningClaims()
}

// SignToken ...
func (s *authRepo) SignToken(ctx context.Context, authClaims *Claims) (*TokenResponse, error) {
	// token
	if authClaims.ID == "" {
		authClaims.ID = uuidpkg.NewUUID()
	}
	if authClaims.ExpiresAt == nil {
		authClaims.ExpiresAt = DefaultExpireTime()
	}
	tokenString, err := s.signEncryptor.EncryptToken(ctx, authClaims)
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshClaims := DefaultRefreshClaims(authClaims)
	refreshToken, err := s.refreshCrypto.EncryptToken(ctx, refreshClaims)
	if err != nil {
		return nil, err
	}

	// 存储
	var (
		userIdentifier  = authClaims.Payload.UserIdentifier()
		accessTokenItem = &TokenItem{
			TokenID:        authClaims.ID,
			RefreshTokenID: refreshClaims.ID,
			ExpiredAt:      authClaims.ExpiresAt.Time.Unix(),
			IsRefreshToken: false,
			Payload:        authClaims.Payload,
		}
		refreshTokenItem = &TokenItem{
			TokenID:        authClaims.ID,
			RefreshTokenID: refreshClaims.ID,
			ExpiredAt:      refreshClaims.ExpiresAt.Time.Unix(),
			IsRefreshToken: true,
			Payload:        refreshClaims.Payload,
		}
		tokenItems = []*TokenItem{accessTokenItem, refreshTokenItem}
	)

	// save token
	if s.tokenManger != nil {
		err = s.tokenManger.SaveTokens(ctx, userIdentifier, tokenItems)
		if err != nil {
			return nil, err
		}
	}

	// 登录限制
	threadpkg.GoSafe(func() {
		if s.tokenManger != nil {
			s.checkLimitAndDeleteExpireTokens(ctx, authClaims)
		}
	})

	res := &TokenResponse{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,

		AccessTokenItem:  accessTokenItem,
		RefreshTokenItem: refreshTokenItem,
	}
	return res, nil
}

// checkLimitAndDeleteExpireTokens ...
func (s *authRepo) checkLimitAndDeleteExpireTokens(ctx context.Context, authClaims *Claims) {
	checkErr := s.checkLimitAndLogoutOtherAccount(ctx, authClaims)
	if checkErr != nil {
		s.logHandler.WithContext(ctx).Errorw("msg", "checkLoginLimit failed", "err", checkErr)
	}
	deleteErr := s.deleteExpireTokens(ctx, authClaims)
	if deleteErr != nil {
		s.logHandler.WithContext(ctx).Errorw("msg", "deleteExpireTokens failed", "err", deleteErr)
	}
}

// deleteExpireTokens 检查登录限制
func (s *authRepo) deleteExpireTokens(ctx context.Context, authClaims *Claims) error {
	var (
		userIdentifier = authClaims.Payload.UserIdentifier()
		nowUnix        = time.Now().Unix()
		expireList     []*TokenItem
	)

	allTokens, err := s.tokenManger.GetAllTokens(ctx, userIdentifier)
	if err != nil {
		e := errorpkg.ErrorBadRequest("GetAllTokens failed")
		err = errorpkg.Wrap(e, err)
		return err
	}
	for i := range allTokens {
		if allTokens[i].ExpiredAt > nowUnix {
			continue
		}
		expireList = append(expireList, allTokens[i])
	}

	// 删除过期
	if err = s.tokenManger.DeleteTokens(ctx, userIdentifier, expireList); err != nil {
		e := errorpkg.ErrorBadRequest("DeleteTokens failed")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// checkLimitAndLogoutOtherAccount 检查登录限制
func (s *authRepo) checkLimitAndLogoutOtherAccount(ctx context.Context, authClaims *Claims) error {
	if authClaims.Payload.LoginLimit == LoginLimitEnum_UNLIMITED {
		return nil
	}
	userIdentifier := authClaims.Payload.UserIdentifier()
	allTokens, err := s.tokenManger.GetAllTokens(ctx, userIdentifier)
	if err != nil {
		e := errorpkg.ErrorBadRequest("GetAllTokens failed")
		err = errorpkg.Wrap(e, err)
		return err
	}

	var (
		blacklist []*TokenItem
		limitList []*TokenItem
	)
	for iKey := range allTokens {
		// 不检查刷新token
		if allTokens[iKey].IsRefreshToken {
			continue
		}
		// 跳过自己
		if allTokens[iKey].TokenID == authClaims.ID {
			continue
		}

		isLimit := false
		switch authClaims.Payload.LoginLimit {
		case LoginLimitEnum_ONLY_ONE:
			// 同一账户仅允许登录一次
			isLimit = true
		case LoginLimitEnum_PLATFORM_ONE:
			// 同一账户每个平台都可登录一次
			if authClaims.Payload.LoginPlatform == allTokens[iKey].Payload.LoginPlatform {
				isLimit = true
			}
		}
		if isLimit {
			blacklist = append(blacklist, allTokens[iKey])
			limitList = append(limitList, allTokens[iKey])
			if item, ok := allTokens[allTokens[iKey].RefreshTokenID]; ok {
				blacklist = append(blacklist, item)
			}
		}
	}

	// 添加黑名单
	if err = s.tokenManger.AddBlacklist(ctx, userIdentifier, blacklist); err != nil {
		e := errorpkg.ErrorBadRequest("AddBlacklist failed")
		err = errorpkg.Wrap(e, err)
		return err
	}
	// 添加登录限制
	if err = s.tokenManger.AddLoginLimit(ctx, limitList); err != nil {
		e := errorpkg.ErrorBadRequest("AddLoginLimit failed")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// DecodeAccessToken ...
func (s *authRepo) DecodeAccessToken(ctx context.Context, accessToken string) (*Claims, error) {
	claims, err := s.signEncryptor.DecryptToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	// 验证有效性
	if err = claims.Valid(); err != nil {
		e := ErrorTokenExpired("access token expired")
		err = errorpkg.Wrap(e, err)
		return nil, err
	}
	return claims, err
}

// DecodeRefreshToken ...
func (s *authRepo) DecodeRefreshToken(ctx context.Context, refreshToken string) (*Claims, error) {
	claims, err := s.refreshCrypto.DecryptToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	// 验证有效性
	if err = claims.Valid(); err != nil {
		e := ErrorTokenExpired("refresh token expired")
		err = errorpkg.Wrap(e, err)
		return nil, err
	}
	return claims, err
}

// VerifyToken 验证令牌
func (s *authRepo) VerifyToken(ctx context.Context, jwtToken *jwt.Token) error {
	authClaims, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return ErrTokenInvalid()
	}

	// 检查 黑名单 & 白名单
	if s.tokenManger != nil {
		if err := s.checkTokenBlackAndWhite(ctx, authClaims); err != nil {
			return err
		}
	}
	return nil
}

// checkTokenBlackAndWhite 检查黑白名单
func (s *authRepo) checkTokenBlackAndWhite(ctx context.Context, authClaims *Claims) error {
	// 黑名单
	isBlacklist, err := s.tokenManger.IsBlacklist(ctx, authClaims.ID)
	if err != nil {
		e := ErrInvalidClaims()
		e.Metadata = map[string]string{"err": err.Error()}
		return e
	}
	if isBlacklist {
		return ErrBlacklist()
	}

	// 白名单
	isExist, err := s.tokenManger.IsExistToken(ctx, authClaims.Payload.UserIdentifier(), authClaims.ID)
	if err != nil {
		e := ErrInvalidClaims()
		e.Metadata = map[string]string{"err": err.Error()}
		return e
	}
	if !isExist {
		return ErrWhitelist()
	}
	return nil
}
