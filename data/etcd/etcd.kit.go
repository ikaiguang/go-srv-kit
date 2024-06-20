package etcdpkg

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Config struct {
	Endpoints          []string
	Username           string
	Password           string
	DialTimeout        *durationpb.Duration
	CaCert             []byte
	InsecureSkipVerify bool
}

func NewEtcdClient(conf *Config) (*clientv3.Client, error) {
	etcdConfig := &clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.Username,
		Password:    conf.Password,
		DialTimeout: conf.DialTimeout.AsDuration(),
	}

	// ca cert
	if len(conf.CaCert) > 0 {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(conf.CaCert)
		tlsConfig := &tls.Config{
			InsecureSkipVerify: conf.InsecureSkipVerify,
			RootCAs:            caCertPool,
		}
		etcdConfig.TLS = tlsConfig
	}
	//if conf.InsecureSkipVerify {
	//	etcdConfig.DialOptions = append(etcdConfig.DialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//}
	return NewClient(etcdConfig)
}

// NewClient ...
func NewClient(config *clientv3.Config) (*clientv3.Client, error) {
	etcdCC, err := clientv3.New(*config)
	if err != nil {
		return nil, err
	}

	_, err = etcdCC.Put(context.Background(), "/ping", "pong")
	if err != nil {
		return etcdCC, err
	}
	return etcdCC, nil
}
