package authpkg

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	aespkg "github.com/ikaiguang/go-srv-kit/kit/aes"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
)

// Config ...
type Config struct {
	SignCrypto    SignEncryptor
	RefreshCrypto RefreshEncryptor

	AccessTokenExpire   time.Duration
	RefreshTokenExpire  time.Duration
	PreviousTokenExpire time.Duration
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
	RefreshToken(ctx context.Context, originRefreshClaims *Claims) (*TokenResponse, error)

	DecodeAccessToken(ctx context.Context, accessToken string) (*Claims, error)
	DecodeRefreshToken(ctx context.Context, refreshToken string) (*Claims, error)

	VerifyAccessToken(ctx context.Context, authClaims *Claims) error
	VerifyRefreshToken(ctx context.Context, authClaims *Claims) error
}

// authRepo ...
type authRepo struct {
	accessTokenExpire   time.Duration
	refreshTokenExpire  time.Duration
	previousTokenExpire time.Duration

	log           *log.Helper
	signEncryptor SignEncryptor
	refreshCrypto RefreshEncryptor
	tokenManger   TokenManger

	jwtValidator *jwt.Validator
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
	if config.AccessTokenExpire < 1 {
		config.AccessTokenExpire = AccessTokenExpire
	}
	if config.RefreshTokenExpire < 1 {
		config.RefreshTokenExpire = RefreshTokenExpire
	}
	if config.PreviousTokenExpire < 1 {
		config.PreviousTokenExpire = PreviousTokenExpire
	}
	// authCacheKeyPrefix := CheckAuthCacheKeyPrefix(config.AuthCacheKeyPrefix)
	return &authRepo{
		accessTokenExpire:   config.AccessTokenExpire,
		refreshTokenExpire:  config.RefreshTokenExpire,
		previousTokenExpire: config.PreviousTokenExpire,

		signEncryptor: config.SignCrypto,
		refreshCrypto: config.RefreshCrypto,
		log:           log.NewHelper(log.With(logger, "module", "kit.auth.token.repo")),
		tokenManger:   tokenManger,
		// tokenManger:   NewTokenManger(redisCC, authCacheKeyPrefix),

		jwtValidator: jwt.NewValidator(),
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
// Note: CheckAndCorrectAuthClaims
func (s *authRepo) SignToken(ctx context.Context, authClaims *Claims) (*TokenResponse, error) {
	// check CheckAndCorrectAuthClaims
	// authClaims.ID = authClaims.Payload.TokenID
	authClaims.ExpiresAt = GenExpireAt(s.accessTokenExpire)

	// token
	tokenString, err := s.signEncryptor.EncryptToken(ctx, authClaims)
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshClaims := GenRefreshClaimsByAuthClaims(authClaims, s.refreshTokenExpire)
	refreshToken, err := s.refreshCrypto.EncryptToken(ctx, refreshClaims)
	if err != nil {
		return nil, err
	}

	// 存储
	var (
		userIdentifier                    = authClaims.Payload.UserIdentifier()
		accessTokenItem, refreshTokenItem = s.genTokenItems(authClaims, refreshClaims)
		tokenItems                        = []*TokenItem{accessTokenItem, refreshTokenItem}
	)

	// save token
	if s.tokenManger != nil {
		err = s.tokenManger.SaveAccessTokens(ctx, userIdentifier, tokenItems)
		if err != nil {
			return nil, err
		}
	}

	// 登录限制
	threadpkg.GoSafe(func() {
		s.checkLimitAndDeleteExpireTokens(ctx, authClaims)
	})

	res := &TokenResponse{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,

		AccessTokenItem:  accessTokenItem,
		RefreshTokenItem: refreshTokenItem,
	}
	return res, nil
}

// RefreshToken ...
// Note: authClaims = refreshToken.authClaims
func (s *authRepo) RefreshToken(ctx context.Context, originRefreshClaims *Claims) (*TokenResponse, error) {
	// token
	authClaims := GenAuthClaimsByAuthClaims(originRefreshClaims, s.accessTokenExpire)
	tokenString, err := s.signEncryptor.EncryptToken(ctx, authClaims)
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshClaims := GenRefreshClaimsByAuthClaims(authClaims, s.refreshTokenExpire)
	refreshToken, err := s.refreshCrypto.EncryptToken(ctx, refreshClaims)
	if err != nil {
		return nil, err
	}

	// 存储
	var (
		userIdentifier                    = authClaims.Payload.UserIdentifier()
		accessTokenItem, refreshTokenItem = s.genTokenItems(authClaims, refreshClaims)
		tokenItems                        = []*TokenItem{accessTokenItem, refreshTokenItem}
	)

	// save token
	if s.tokenManger != nil {
		// 设置被刷新的token过期时间为 PreviousTokenExpire
		if err = s.setPreviousTokenExpireTime(ctx, originRefreshClaims); err != nil {
			return nil, err
		}
		err = s.tokenManger.SaveAccessTokens(ctx, userIdentifier, tokenItems)
		if err != nil {
			return nil, err
		}
		// 清除过期Token
		threadpkg.GoSafe(func() {
			// ctx many be cancel
			deleteErr := s.tokenManger.DeleteExpireTokens(context.Background(), userIdentifier)
			if deleteErr != nil {
				s.log.WithContext(ctx).Errorw("msg", "deleteExpireTokens failed", "err", deleteErr)
			}
		})
	}

	res := &TokenResponse{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,

		AccessTokenItem:  accessTokenItem,
		RefreshTokenItem: refreshTokenItem,
	}
	return res, nil
}

// genTokenItems ...
func (s *authRepo) genTokenItems(authClaims, refreshClaims *Claims) (accessTokenItem, refreshTokenItem *TokenItem) {
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
	return accessTokenItem, refreshTokenItem
}

// checkLimitAndDeleteExpireTokens ...
func (s *authRepo) checkLimitAndDeleteExpireTokens(ctx context.Context, authClaims *Claims) {
	// ctx many be cancel
	newCtx := context.Background()
	checkErr := s.checkLimitAndLogoutOtherAccount(newCtx, authClaims)
	if checkErr != nil {
		s.log.WithContext(ctx).Errorw("msg", "checkLoginLimit failed", "err", checkErr)
	}
	if s.tokenManger != nil {
		deleteErr := s.tokenManger.DeleteExpireTokens(newCtx, authClaims.Payload.UserIdentifier())
		if deleteErr != nil {
			s.log.WithContext(ctx).Errorw("msg", "deleteExpireTokens failed", "err", deleteErr)
		}
	}
}

// setPreviousTokenExpireTime 设置上一个令牌的过期时间
func (s *authRepo) setPreviousTokenExpireTime(ctx context.Context, originRefreshClaims *Claims) error {
	if s.tokenManger == nil {
		return nil
	}
	var (
		userIdentifier   = originRefreshClaims.Payload.UserIdentifier()
		refreshTokenId   = originRefreshClaims.Payload.TokenID
		previousExpireAt = time.Now().Add(s.previousTokenExpire).Unix()
		tokenItems       []*TokenItem
	)
	refreshTokenItem, isNotFound, err := s.tokenManger.GetToken(ctx, userIdentifier, refreshTokenId)
	if err != nil {
		return err
	}
	if isNotFound {
		return nil
	}
	tokenItems = append(tokenItems, refreshTokenItem)
	if refreshTokenItem.ExpiredAt > previousExpireAt {
		refreshTokenItem.ExpiredAt = previousExpireAt
	}

	var accessTokenId = refreshTokenItem.TokenID
	accessTokenItem, isNotFound, err := s.tokenManger.GetToken(ctx, userIdentifier, accessTokenId)
	if err != nil {
		return err
	}
	if !isNotFound {
		if accessTokenItem.ExpiredAt > previousExpireAt {
			accessTokenItem.ExpiredAt = previousExpireAt
		}
		tokenItems = append(tokenItems, accessTokenItem)
	}
	if err = s.tokenManger.ResetPreviousTokens(ctx, userIdentifier, tokenItems); err != nil {
		return err
	}
	return nil
}

// checkLimitAndLogoutOtherAccount 检查登录限制
func (s *authRepo) checkLimitAndLogoutOtherAccount(ctx context.Context, authClaims *Claims) error {
	if s.tokenManger == nil {
		return nil
	}
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
	//if err = claims.Valid(); err != nil {
	if err = s.jwtValidator.Validate(claims); err != nil {
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
	//if err = claims.Valid(); err != nil {
	if err = s.jwtValidator.Validate(claims); err != nil {
		e := ErrorTokenExpired("refresh token expired")
		err = errorpkg.Wrap(e, err)
		return nil, err
	}
	return claims, err
}

// VerifyAccessToken 验证令牌
func (s *authRepo) VerifyAccessToken(ctx context.Context, authClaims *Claims) error {
	if s.tokenManger == nil {
		return nil
	}
	// 检查 黑名单 & 白名单
	if err := s.checkTokenBlackAndWhite(ctx, authClaims); err != nil {
		return err
	}
	return nil
}

// VerifyRefreshToken 验证令牌
func (s *authRepo) VerifyRefreshToken(ctx context.Context, authClaims *Claims) error {
	if s.tokenManger == nil {
		return nil
	}
	// 检查 黑名单 & 白名单
	if err := s.checkTokenBlackAndWhite(ctx, authClaims); err != nil {
		return err
	}
	return nil
}

// checkTokenBlackAndWhite 检查黑白名单
func (s *authRepo) checkTokenBlackAndWhite(ctx context.Context, authClaims *Claims) error {
	// 黑名单
	isBlacklist, err := s.tokenManger.IsBlacklist(ctx, authClaims.ID)
	if err != nil {
		return err
	}
	if isBlacklist {
		e := ErrBlacklist()
		return errorpkg.WithStack(e)
	}

	// 白名单
	isExist, err := s.tokenManger.IsExistToken(ctx, authClaims.Payload.UserIdentifier(), authClaims.ID)
	if err != nil {
		return err
	}
	if !isExist {
		e := ErrWhitelist()
		return errorpkg.WithStack(e)
	}
	return nil
}
