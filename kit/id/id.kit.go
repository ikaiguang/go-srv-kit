package idpkg

// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
import (
	"sync"

	"github.com/bwmarrin/snowflake"
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
	var err error
	_idNodeOnce.Do(func() {
		_idNode, err = snowflake.NewNode(_nodeID)
	})
	if err != nil {
		_idNode = nil
		_idNodeOnce = sync.Once{}
	}
}

// SetNode 设置节点
func SetNode(node *snowflake.Node) {
	_idNode = node
}

func NewNode(node int64) (*snowflake.Node, error) {
	return snowflake.NewNode(node)
}

// ID ...
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func ID() int64 {
	return _idNode.Generate().Int64()
}

func NextID() int64 {
	return _idNode.Generate().Int64()
}
