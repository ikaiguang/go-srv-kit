package etcdpkg

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const defaultProbeTimeout = 5 * time.Second

func NewEtcdClient(conf *Config) (*clientv3.Client, error) {
	etcdConfig, err := buildClientConfig(conf)
	if err != nil {
		return nil, err
	}
	return NewClient(etcdConfig)
}

func buildClientConfig(conf *Config) (*clientv3.Config, error) {
	if conf == nil {
		return nil, errors.New("etcd config is nil")
	}

	etcdConfig := &clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.Username,
		Password:    conf.Password,
		DialTimeout: conf.DialTimeout.AsDuration(),
	}

	if len(conf.CaCert) > 0 || conf.InsecureSkipVerify {
		// InsecureSkipVerify is an explicit caller configuration option.
		tlsConfig := &tls.Config{InsecureSkipVerify: conf.InsecureSkipVerify} //nolint:gosec
		if len(conf.CaCert) > 0 {
			caCertPool := x509.NewCertPool()
			if ok := caCertPool.AppendCertsFromPEM(conf.CaCert); !ok {
				return nil, errors.New("etcd CA certificate is not valid PEM")
			}
			tlsConfig.RootCAs = caCertPool
		}
		etcdConfig.TLS = tlsConfig
	}
	return etcdConfig, nil
}

// NewClient ...
func NewClient(config *clientv3.Config) (*clientv3.Client, error) {
	if config == nil {
		return nil, errors.New("etcd client config is nil")
	}

	etcdCC, err := clientv3.New(*config)
	if err != nil {
		return nil, err
	}

	probeTimeout := config.DialTimeout
	if probeTimeout <= 0 {
		probeTimeout = defaultProbeTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), probeTimeout)
	defer cancel()

	const pingKey = "/ping"
	_, err = etcdCC.Get(ctx, pingKey, clientv3.WithLimit(1))
	if err != nil {
		if closeErr := etcdCC.Close(); closeErr != nil {
			return nil, fmt.Errorf("etcd ping failed: %w; close client: %v", err, closeErr)
		}
		return nil, fmt.Errorf("etcd ping failed: %w", err)
	}
	return etcdCC, nil
}
