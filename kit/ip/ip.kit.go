package ippkg

import (
	"net"
	"sync"
)

const (
	_ip = "127.0.0.1"
)

var (
	_IP           = net.ParseIP(_ip)
	_localIP      = "127.0.0.1"
	_localIpMutex = sync.Once{}
)

// LocalIP 本地IP
func LocalIP() string {
	_localIpMutex.Do(func() {
		_localIP = PrivateIPv4().String()
	})
	return _localIP
}

// NewLocalIP ...
func NewLocalIP() string {
	return PrivateIPv4().String()
}

// IsValidIP 有效的ip
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// PrivateIPv4 ...
func PrivateIPv4() net.IP {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return _IP
	}

	for _, a := range as {
		ipNet, ok := a.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}

		ip := ipNet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip
		}
	}
	return _IP
}

// isPrivateIPv4 ...
func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}
