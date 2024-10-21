package authutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"github.com/redis/go-redis/v9"
	"sync"
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

func GetTokenManger(authInstance AuthInstance) (authpkg.TokenManger, error) {
	return authInstance.GetTokenManger()
}
