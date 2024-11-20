package idpkg

import (
	"fmt"
	ippkg "github.com/ikaiguang/go-srv-kit/kit/ip"
	"net"
)

var (
	// Node 生成ID的节点
	// 为了帮助保证唯一性
	// - 确保您的系统保持准确的系统时间
	// - 确保您永远不会有多个节点以相同的节点 ID 运行
	Node Snowflake
)

// ===== Benchmark =====
// BenchmarkNew_SonySonyflake-8			31080             38894 ns/op               0 B/op          0 allocs/op
// BenchmarkNew_BwmarrinSnowflake-8		76981             15611 ns/op               0 B/op          0 allocs/op
// ===== Benchmark =====
func init() {
	nodeID, err := GenNodeID()
	if err != nil {
		nodeID = 1
	}
	Node, err = NewBwmarrinSnowflake(int64(nodeID))
	if err != nil {
		panic(err)
	}
}

func SetNode(node Snowflake) {
	Node = node
}

// NextID ...
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func NextID() (uint64, error) {
	return Node.NextID()
}

func GenNodeID() (uint16, error) {
	return IPV4ToNodeID(ippkg.LocalIP())
}

func IPV4ToNodeID(ip string) (uint16, error) {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ip)
	}

	var lastTwoBytes [2]byte
	copy(lastTwoBytes[:], ipAddr.To4()[2:])

	nodeId := uint16(lastTwoBytes[0])<<8 | uint16(lastTwoBytes[1])
	return nodeId, nil
}
