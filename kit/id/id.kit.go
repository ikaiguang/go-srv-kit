// Package idutil
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
package idutil

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

const (
	_nodeID = 1
)

var (
	// _idNode 生成ID的节点
	// 为了帮助保证唯一性
	// - 确保您的系统保持准确的系统时间
	// - 确保您永远不会有多个节点以相同的节点 ID 运行
	_idNode     *snowflake.Node
	_idNodeOnce sync.Once
)

func init() {
	_idNodeOnce.Do(func() {
		var err error
		_idNode, err = snowflake.NewNode(_nodeID)
		if err != nil {
			_idNodeOnce = sync.Once{}
		}
	})
}

// SetNode 设置节点
func SetNode(node *snowflake.Node) {
	_idNode = node
}

// New ...
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func New() int64 {
	return _idNode.Generate().Int64()
}

// NewID ...
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func NewID() uint64 {
	return uint64(_idNode.Generate().Int64())
}
