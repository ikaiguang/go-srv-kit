package etcdpkg

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// NewEtcdClient .
func NewEtcdClient(config *clientv3.Config) (*clientv3.Client, error) {
	return NewClient(config)
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
