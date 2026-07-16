package idpkg

import (
	"fmt"
	"log/slog"
	"net"
	"sync"

	ippkg "github.com/ikaiguang/go-srv-kit/kit/v3/ip"
)

var (
	// nodeHandler 生成ID的节点
	// use SetNode for synchronized updates.
	// 为了帮助保证唯一性
	// - 确保您的系统保持准确的系统时间
	// - 确保您永远不会有多个节点以相同的节点 ID 运行
	nodeHandler Snowflake

	// nodeMu protects the package-managed nodeHandler and initialization error.
	nodeMu sync.RWMutex
	// nodeErr 记录初始化错误，延迟到 NextID 时返回
	nodeErr error
)

// 默认 Snowflake 配置为 12 位节点号和 8 位序列号：
// - 最多 4096 个节点
// - 单节点每毫秒最多 256 个 ID
// - 基于 2026-01-01 的 epoch，理论寿命约 279 年
func init() {
	nodeMu.Lock()
	defer nodeMu.Unlock()
	initializeNodeLocked()
}

func initializeNodeLocked() {
	nodeID, err := GenerateNodeID()
	if err != nil {
		nodeID = 1
	}
	nodeHandler, err = NewBwmarrinSnowflake(int64(nodeID))
	if err != nil {
		// 记录错误，延迟到 NextID 时处理
		nodeErr = err
		slog.Warn("BwmarrinSnowflake init failed", "err", err)
		nodeHandler = nil
	}
}

// SetNode 设置自定义 nodeHandler 实例，同时清除初始化错误
func SetNode(node Snowflake) {
	if node == nil {
		return
	}
	nodeMu.Lock()
	defer nodeMu.Unlock()
	nodeHandler = node
	nodeErr = nil
}

// NextID 生成下一个唯一 ID
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func NextID() (uint64, error) {
	nodeMu.RLock()
	node := nodeHandler
	err := nodeErr
	nodeMu.RUnlock()

	if node == nil {
		nodeMu.Lock()
		if nodeHandler == nil {
			initializeNodeLocked()
		}
		node = nodeHandler
		err = nodeErr
		nodeMu.Unlock()
	}
	if err != nil {
		slog.Warn("BwmarrinSnowflake init failed", "err", err)
		return 0, err
	}
	return node.NextID()
}

// GenerateNodeID derives a Snowflake node ID from the local IPv4 address.
func GenerateNodeID() (uint16, error) {
	return IPv4ToNodeID(ippkg.LocalIP())
}

// IPv4ToNodeID derives a Snowflake node ID from the final IPv4 bytes.
func IPv4ToNodeID(ip string) (uint16, error) {
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

// Deprecated: use GenerateNodeID instead.
func GenNodeID() (uint16, error) { return GenerateNodeID() }

// Deprecated: use IPv4ToNodeID instead.
func IPV4ToNodeID(ip string) (uint16, error) { return IPv4ToNodeID(ip) }
