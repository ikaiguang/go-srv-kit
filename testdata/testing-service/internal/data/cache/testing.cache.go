package caches

import (
	"github.com/go-kratos/kratos/v2/log"
	datarepos "github.com/ikaiguang/go-srv-kit/testdata/testing-service/internal/data/repo"
)

type testingCache struct {
	log *log.Helper
}

func NewTestingCache(
	logger log.Logger,
) datarepos.TestingCacheRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/data/cache"))

	return &testingCache{
		log: logHelper,
	}
}
