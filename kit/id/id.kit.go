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
func New() int64 {
	return _idNode.Generate().Int64()
}

// NewID ...
func NewID() uint64 {
	return uint64(_idNode.Generate().Int64())
}
