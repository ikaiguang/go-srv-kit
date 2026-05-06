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

// 默认 Snowflake 配置为 12 位节点号和 8 位序列号：
// - 最多 4096 个节点
// - 单节点每毫秒最多 256 个 ID
// - 基于 2026-01-01 的 epoch，理论寿命约 279 年
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
	ipv4 := ipAddr.To4()
	if ipv4 == nil {
		return 0, fmt.Errorf("invalid IPv4 address: %s", ip)
	}

	var lastTwoBytes [2]byte
	copy(lastTwoBytes[:], ipv4[2:])

	nodeID := uint16(lastTwoBytes[0])<<8 | uint16(lastTwoBytes[1])
	return nodeID & uint16(snowflakeMaxNode), nil
}
