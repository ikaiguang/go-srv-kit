package consulpkg

import (
	"github.com/hashicorp/consul/api"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Config consul config
type Config struct {
	Scheme             string
	Address            string
	PathPrefix         string
	Datacenter         string
	WaitTime           *durationpb.Duration
	Token              string
	Namespace          string
	Partition          string
	WithHttpBasicAuth  bool
	AuthUsername       string
	AuthPassword       string
	InsecureSkipVerify bool
	TlsAddress         string
	TlsCaPem           string
	TlsCertPem         string
	TlsKeyPem          string
}

// NewConsulClient .
func NewConsulClient(conf *Config, opts ...Option) (*api.Client, error) {
	return NewClient(conf, opts...)
}

// NewClient ...
func NewClient(conf *Config, opts ...Option) (*api.Client, error) {
	defConfig := api.DefaultConfig()
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

	// new client
	consulCC, err := api.NewClient(defConfig)
	if err != nil {
		return consulCC, err
	}

	// ping
	kv := &api.KVPair{Key: "ping", Value: []byte("pong")}
	_, err = consulCC.KV().Put(kv, nil)
	if err != nil {
		return consulCC, err
	}
	//newKv, _, err := consulCC.KV().Get(kv.Key, nil)
	//if err != nil {
	//	return consulCC, err
	//}
	//_ = newKv

	return consulCC, err
}
