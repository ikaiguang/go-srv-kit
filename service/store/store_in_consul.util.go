package storeutil

import (
	"context"
	consulapi "github.com/hashicorp/consul/api"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdlog "log"
)

// StoreInConsul 存储文件到consul
// @param sourceDir 需要存储的文件所在目录；例：/path/to/dir
// @param storeDir 存储到consul的路径；例：go-micro-saas/general-config
func StoreInConsul(consulClient *consulapi.Client, sourceDir, storeDir string) error {
	if consulClient == nil {
		e := errorpkg.ErrorBadRequest("请配置consul客户端：consulClient")
		return errorpkg.WithStack(e)
	}
	if sourceDir == "" {
		e := errorpkg.ErrorBadRequest("请配置资源目录：sourceDir")
		return errorpkg.WithStack(e)
	}
	if storeDir == "" {
		e := errorpkg.ErrorBadRequest("请配置存储路径：storeDir")
		return errorpkg.WithStack(e)
	}

	// 开始配置
	stdlog.Println("|==================== 更新配置到Consul 开始 ====================|")
	defer stdlog.Println("|==================== 更新配置到Consul 结束 ====================|")
	stdlog.Println("|*** 资源路径：	", sourceDir)
	stdlog.Println("|*** 存储路径：	", storeDir)

	handler := NewStoreInConsulHandler(consulClient)

	return handler.Save(sourceDir, storeDir)
}

type storeInConsulHandler struct {
	consulClient *consulapi.Client
}

func NewStoreInConsulHandler(consulClient *consulapi.Client) StoreManager {
	return &storeInConsulHandler{
		consulClient: consulClient,
	}
}

// Save 存储到consul
func (s *storeInConsulHandler) Save(sourceDir, storeDir string) error {
	if s.consulClient == nil {
		e := errorpkg.ErrorBadRequest("请配置consul客户端：consulClient")
		return errorpkg.WithStack(e)
	}

	storeDataM, err := ReadStoreFiles(sourceDir, storeDir)
	if err != nil {
		return err
	}
	ctx := context.Background()
	opt := &consulapi.WriteOptions{}
	opt = opt.WithContext(ctx)
	for key := range storeDataM {
		stdlog.Println("|*** 存储文件：", key)
		kv := &consulapi.KVPair{
			Key:   key,
			Value: storeDataM[key],
		}
		_, err := s.consulClient.KV().Put(kv, opt)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return errorpkg.WithStack(e)
		}
	}
	return nil
}
