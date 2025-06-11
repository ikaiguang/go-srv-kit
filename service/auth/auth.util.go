package authutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	stdlog "log"
	"sync"

	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/redis/go-redis/v9"
)

type authInstance struct {
	conf          *configpb.Encrypt_TokenEncrypt
	redisCC       redis.UniversalClient
	loggerManager loggerutil.LoggerManager

	// 不要直接使用 s.tokenXxx, 请使用 GetAuthCollection()
	tokenManager     authpkg.TokenManger
	tokenAuthRepo    authpkg.AuthRepo
	tokenManagerOnce sync.Once
}

type AuthCollection struct {
	TokenManager authpkg.TokenManger
	AuthManager  authpkg.AuthRepo
}

type AuthInstance interface {
	GetTokenManger() (authpkg.TokenManger, error)
	GetAuthManger() (authpkg.AuthRepo, error)
}

func NewAuthInstance(conf *configpb.Encrypt_TokenEncrypt, redisCC redis.UniversalClient, loggerManager loggerutil.LoggerManager) (AuthInstance, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = encrypt.token_encrypt")
		return nil, errorpkg.WithStack(e)
	}
	if conf.GetSignKey() == "" {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = encrypt.token_encrypt.sign_key")
		return nil, errorpkg.WithStack(e)
	}
	if conf.GetRefreshKey() == "" {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = encrypt.token_encrypt.refresh_key")
		return nil, errorpkg.WithStack(e)
	}
	return &authInstance{
		conf:          conf,
		redisCC:       redisCC,
		loggerManager: loggerManager,
	}, nil
}

func (s *authInstance) GetAuthCollection() (*AuthCollection, error) {
	err := s.loadingTokenManagerOnce()
	if err != nil {
		return nil, err
	}
	return &AuthCollection{
		TokenManager: s.tokenManager,
		AuthManager:  s.tokenAuthRepo,
	}, nil
}

func (s *authInstance) GetTokenManger() (authpkg.TokenManger, error) {
	manager, err := s.GetAuthCollection()
	if err != nil {
		return nil, err
	}
	return manager.TokenManager, nil
}
func (s *authInstance) GetAuthManger() (authpkg.AuthRepo, error) {
	manager, err := s.GetAuthCollection()
	if err != nil {
		return nil, err
	}
	return manager.AuthManager, nil
}

func (s *authInstance) loadingTokenManagerOnce() error {
	var err error
	s.tokenManagerOnce.Do(func() {
		s.tokenManager, s.tokenAuthRepo, err = s.loadingTokenManager()
	})
	if err != nil {
		s.tokenManagerOnce = sync.Once{}
	}
	return err
}

func (s *authInstance) loadingTokenManager() (authpkg.TokenManger, authpkg.AuthRepo, error) {
	stdlog.Println("|*** LOADING: TokenManger: ...")
	logger, err := s.loggerManager.GetLoggerForMiddleware()
	if err != nil {
		return nil, nil, err
	}
	tokenManger := authpkg.NewTokenManger(logger, s.redisCC, authpkg.CheckAuthCacheKeyPrefix(nil))
	config := &authpkg.Config{
		SignCrypto:          authpkg.NewSignEncryptor(s.conf.GetSignKey()),
		RefreshCrypto:       authpkg.NewCBCCipher(s.conf.GetRefreshKey()),
		AccessTokenExpire:   s.conf.GetAccessTokenExpire().AsDuration(),
		RefreshTokenExpire:  s.conf.GetRefreshTokenExpire().AsDuration(),
		PreviousTokenExpire: s.conf.GetPreviousTokenExpire().AsDuration(),
	}
	stdlog.Println("|*** LOADING: AuthManger: ...")
	authRepo, err := authpkg.NewAuthRepo(*config, logger, tokenManger)
	if err != nil {
		return nil, nil, err
	}
	return tokenManger, authRepo, nil
}
