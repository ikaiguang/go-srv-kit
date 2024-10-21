package clientutil

import (
	"sync"
)

var (
	_apiConnection = sync.Map{}
)

func NewSingletonServiceAPIConnection(serviceAPIManager ServiceAPIManager, serviceName ServiceName) (ServiceAPIConnection, error) {
	cc, ok := _apiConnection.Load(serviceName)
	if ok {
		if conn, ok := cc.(ServiceAPIConnection); ok {
			return conn, nil
		}
	}
	conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	_apiConnection.Store(serviceName, conn)
	return conn, nil
}

func NewServiceAPIConnection(serviceAPIManager ServiceAPIManager, serviceName ServiceName) (ServiceAPIConnection, error) {
	conn, err := serviceAPIManager.NewAPIConnection(serviceName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
