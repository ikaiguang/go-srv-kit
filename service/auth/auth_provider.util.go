package authutil

import (
	"sync"

	authpkg "github.com/ikaiguang/go-auth-kit/auth"
	configpb "github.com/ikaiguang/go-service-kit/api/config"
	loggerutil "github.com/ikaiguang/go-service-kit/logger"
)

var (
	singletonMutex        sync.Once
	singletonAuthInstance AuthInstance
)

func NewSingletonAuthInstance(
	conf *configpb.Encrypt_TokenEncrypt,
	redisCC redis.UniversalClient,
	loggerManager loggerutil.LoggerManager,
) (AuthInstance, error) {
	var err error
	singletonMutex.Do(func() {
		singletonAuthInstance, err = NewAuthInstance(conf, redisCC, loggerManager)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonAuthInstance, err
}

func GetAuthManager(authInstance AuthInstance) (authpkg.AuthRepo, error) {
	return authInstance.GetAuthManger()
}

func GetTokenManger(authInstance AuthInstance) (authpkg.TokenManager, error) {
	return authInstance.GetTokenManger()
}
