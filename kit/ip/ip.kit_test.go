package ippkg

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./ip -run TestLocalIP
func TestLocalIP(t *testing.T) {
	ip := LocalIP()
	assert.NotEmpty(t, ip, "LocalIP 不应为空")
	assert.True(t, IsValidIP(ip), "LocalIP 返回的应是有效 IP: %s", ip)
}

// go test -v -count 1 ./ip -run TestNewLocalIP
func TestNewLocalIP(t *testing.T) {
	ip := NewLocalIP()
	assert.NotEmpty(t, ip)
	assert.True(t, IsValidIP(ip), "NewLocalIP 返回的应是有效 IP: %s", ip)
}

// go test -v -count 1 ./ip -run TestIsValidIP
func TestIsValidIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{"有效IPv4", "192.168.1.1", true},
		{"有效IPv4_localhost", "127.0.0.1", true},
		{"有效IPv6", "::1", true},
		{"有效IPv6_full", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"空字符串", "", false},
		{"无效IP", "999.999.999.999", false},
		{"非IP字符串", "hello", false},
		{"部分IP", "192.168", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidIP(tt.ip)
			assert.Equal(t, tt.want, got)
		})
	}
}

// go test -v -count 1 ./ip -run TestPrivateIPv4
func TestPrivateIPv4(t *testing.T) {
	ip := PrivateIPv4()
	assert.NotNil(t, ip, "PrivateIPv4 不应返回 nil")
	assert.True(t, IsValidIP(ip.String()), "PrivateIPv4 返回的应是有效 IP: %s", ip.String())
}

// go test -v -count 1 ./ip -run TestIsPrivateIPv4
func TestIsPrivateIPv4(t *testing.T) {
	assert.True(t, isPrivateIPv4(netParseIP(t, "10.0.0.1")))
	assert.True(t, isPrivateIPv4(netParseIP(t, "172.16.0.1")))
	assert.True(t, isPrivateIPv4(netParseIP(t, "192.168.1.1")))
	assert.False(t, isPrivateIPv4(netParseIP(t, "8.8.8.8")))
	assert.False(t, isPrivateIPv4(netParseIP(t, "::1")))
	assert.False(t, isPrivateIPv4(nil))
}

func netParseIP(t *testing.T, raw string) net.IP {
	t.Helper()
	ip := net.ParseIP(raw)
	assert.NotNil(t, ip)
	return ip
}
