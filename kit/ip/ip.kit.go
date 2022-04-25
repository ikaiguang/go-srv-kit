package iputil

import (
	"net"
	"sync"
)

var (
	_localIP      = "127.0.0.1"
	_localIpMutex = &sync.Once{}
)

// LocalIP 本地IP
func LocalIP() string {
	_localIpMutex.Do(func() {
		_localIP = getLocalIP()
	})
	return _localIP
}

// getLocalIP 本地IP
func getLocalIP() string {
	localIp := "127.0.0.1"
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return localIp
	}
	for i := range addr {
		if ipNet, ok := addr[i].(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				localIp = ipNet.IP.String()
				break
			}
		}
	}
	return localIp
}
