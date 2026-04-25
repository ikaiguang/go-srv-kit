package idpkg

import (
	"fmt"
	"log/slog"
	"net"
	"sync"

	ippkg "github.com/ikaiguang/go-srv-kit/kit/ip"
)

var (
	// Node 生成ID的节点
	// 为了帮助保证唯一性
	// - 确保您的系统保持准确的系统时间
	// - 确保您永远不会有多个节点以相同的节点 ID 运行
	Node Snowflake

	// nodeOnce 用于延迟初始化，确保只初始化一次
	nodeOnce sync.Once
	// nodeErr 记录初始化错误，延迟到 NextID 时返回
	nodeErr error
)

// ===== Benchmark =====
// BenchmarkNew_SonySonyflake-8			31080             38894 ns/op               0 B/op          0 allocs/op
// BenchmarkNew_BwmarrinSnowflake-8		76981             15611 ns/op               0 B/op          0 allocs/op
// ===== Benchmark =====
// SonySonyflake :
// The lifetime (174 years) is longer than that of Snowflake (69 years)
// It can work in more distributed machines (2^16) than Snowflake (2^10)
// It can generate 2^8 IDs per 10 msec at most in a single machine/thread (slower than Snowflake)
// =====
// BwmarrinSnowflake :
// You can alter the number of bits used for the node id and step number (sequence) by
// setting the snowflake.NodeBits and snowflake.StepBits values.
// Remember that There is a maximum of 22 bits available that
// can be shared between these two values. You do not have to use all 22 bits.
func init() {
	// 不再 panic，仅尝试初始化
	nodeID, err := GenNodeID()
	if err != nil {
		nodeID = 1
	}
	Node, err = NewBwmarrinSnowflake(int64(nodeID))
	if err != nil {
		// 记录错误，延迟到 NextID 时处理
		nodeErr = err
		slog.Warn("BwmarrinSnowflake init failed", "err", err)
		Node = nil
	}
}

// SetNode 设置自定义 Node 实例，同时清除初始化错误
func SetNode(node Snowflake) {
	if node == nil {
		return
	}
	Node = node
	nodeErr = nil
}

// NextID 生成下一个唯一 ID
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func NextID() (uint64, error) {
	if Node == nil {
		// 延迟初始化：init() 失败时，在首次调用 NextID 时重试
		nodeOnce.Do(func() {
			if Node != nil {
				return
			}
			nodeID, err := GenNodeID()
			if err != nil {
				nodeID = 1
			}
			Node, nodeErr = NewBwmarrinSnowflake(int64(nodeID))
		})
		if nodeErr != nil {
			slog.Warn("BwmarrinSnowflake init failed", "err", nodeErr)
			nodeOnce = sync.Once{}
			return 0, nodeErr
		}
	}
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
