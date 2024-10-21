package consulutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	stdlog "log"
	"sync"

	consulapi "github.com/hashicorp/consul/api"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type consulManager struct {
	conf *configpb.Consul

	consulOnce   sync.Once
	consulClient *consulapi.Client
}

type ConsulManager interface {
	Enable() bool
	GetClient() (*consulapi.Client, error)
	Close() error
}

func NewConsulManager(conf *configpb.Consul) (ConsulManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = consul")
		return nil, errorpkg.WithStack(e)
	}
	return &consulManager{
		conf: conf,
	}, nil
}

func (s *consulManager) GetClient() (*consulapi.Client, error) {
	var err error
	s.consulOnce.Do(func() {
		s.consulClient, err = s.loadingConsulClient()
	})
	if err != nil {
		s.consulOnce = sync.Once{}
	}
	return s.consulClient, err
}

func (s *consulManager) Close() error {
	if s.consulClient != nil {
		//stdlog.Println("|*** STOP: close: consulClient")
		//err := s.consulClient.Agent().ServiceDeregister(s.conf.ServiceName)
		//if err != nil {
		//	stdlog.Println("|*** STOP: close: consulClient failed: ", err.Error())
		//	return err
		//}
	}
	return nil
}

func (s *consulManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *consulManager) loadingConsulClient() (*consulapi.Client, error) {
	stdlog.Println("|*** LOADING: Consul client: ...")
	cc, err := consulpkg.NewConsulClient(ToConsulConfig(s.conf))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return cc, nil
}

// ToConsulConfig ...
func ToConsulConfig(cfg *configpb.Consul) *consulpkg.Config {
	return &consulpkg.Config{
		Scheme:             cfg.Scheme,
		Address:            cfg.Address,
		PathPrefix:         cfg.PathPrefix,
		Datacenter:         cfg.Datacenter,
		WaitTime:           cfg.WaitTime,
		Token:              cfg.Token,
		Namespace:          cfg.Namespace,
		Partition:          cfg.Partition,
		WithHttpBasicAuth:  cfg.WithHttpBasicAuth,
		AuthUsername:       cfg.AuthUsername,
		AuthPassword:       cfg.AuthPassword,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		TlsAddress:         cfg.TlsAddress,
		TlsCaPem:           cfg.TlsCaPem,
		TlsCertPem:         cfg.TlsCertPem,
		TlsKeyPem:          cfg.TlsKeyPem,
	}
}
