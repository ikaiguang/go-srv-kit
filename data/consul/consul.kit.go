package consulpkg

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-hclog"
)

const defaultProbeTimeout = 5 * time.Second

// NewConsulClient .
func NewConsulClient(conf *Config, opts ...Option) (*api.Client, error) {
	return NewClient(conf, opts...)
}

// NewClient ...
func NewClient(conf *Config, opts ...Option) (*api.Client, error) {
	defConfig, err := buildConfig(conf, opts...)
	if err != nil {
		return nil, err
	}

	consulCC, err := api.NewClient(defConfig)
	if err != nil {
		return nil, err
	}

	if err = probeClient(consulCC); err != nil {
		return nil, err
	}

	return consulCC, nil
}

func probeClient(consulCC *api.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultProbeTimeout)
	defer cancel()
	queryOpts := (&api.QueryOptions{}).WithContext(ctx)
	_, _, err := consulCC.KV().Get("ping", queryOpts)
	return err
}

func buildConfig(conf *Config, opts ...Option) (*api.Config, error) {
	if conf == nil {
		return nil, errors.New("consul config is nil")
	}

	option := &options{}
	for _, opt := range opts {
		if opt != nil {
			opt(option)
		}
	}
	defConfig := api.DefaultConfig()
	if option.writer != nil {
		defConfig = api.DefaultConfigWithLogger(hclog.New(&hclog.LoggerOptions{
			Name:   "consul-api",
			Output: option.writer,
		}))
	}
	// basic
	if conf.Scheme != "" {
		defConfig.Scheme = conf.Scheme
	}
	if conf.Address != "" {
		defConfig.Address = conf.Address
	}
	if conf.PathPrefix != "" {
		defConfig.PathPrefix = conf.PathPrefix
	}
	if conf.Datacenter != "" {
		defConfig.Datacenter = conf.Datacenter
	}
	if conf.WaitTime.AsDuration() > 0 {
		defConfig.WaitTime = conf.WaitTime.AsDuration()
	}
	if conf.Token != "" {
		defConfig.Token = conf.Token
	}
	if conf.Namespace != "" {
		defConfig.Namespace = conf.Namespace
	}
	if conf.Partition != "" {
		defConfig.Partition = conf.Partition
	}

	// auth
	if conf.WithHttpBasicAuth {
		defConfig.HttpAuth = &api.HttpBasicAuth{
			Username: conf.AuthUsername,
			Password: conf.AuthPassword,
		}
	}

	// tls
	defConfig.TLSConfig.InsecureSkipVerify = conf.InsecureSkipVerify
	if conf.TlsAddress != "" {
		defConfig.TLSConfig.Address = conf.TlsAddress
	}
	if conf.TlsCaPem != "" {
		defConfig.TLSConfig.CAPem = []byte(conf.TlsCaPem)
	}
	if conf.TlsCertPem != "" {
		defConfig.TLSConfig.CertPEM = []byte(conf.TlsCertPem)
	}
	if conf.TlsKeyPem != "" {
		defConfig.TLSConfig.KeyPEM = []byte(conf.TlsKeyPem)
	}

	return defConfig, nil
}
