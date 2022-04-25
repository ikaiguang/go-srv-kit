package iputil

import (
	"testing"
)

// go test -v -count=1 ./kit/ip -test.run=TestLocalIP
func TestLocalIP(t *testing.T) {
	LocalIP()
}
