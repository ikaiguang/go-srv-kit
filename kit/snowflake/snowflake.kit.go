package snowflake

import (
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

const (
	_workerID = 1
)

var (
	// _idNode 生成ID的节点
	// 为了帮助保证唯一性
	// - 确保您的系统保持准确的系统时间
	// - 确保您永远不会有多个节点以相同的节点 ID 运行
	_idNode     *sonyflake.Sonyflake
	_idNodeOnce sync.Once
	_startTime  = time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
)

func init() {
	var err error
	_idNodeOnce.Do(func() {
		_idNode, err = sonyflake.New(sonyflake.Settings{
			StartTime: _startTime,
			MachineID: func() (uint16, error) {
				return _workerID, nil
			},
			CheckMachineID: func(id uint16) bool {
				return true
			},
		})
	})
	if err != nil {
		_idNode = nil
		_idNodeOnce = sync.Once{}
	}
}

func NewNode(node uint16) (*sonyflake.Sonyflake, error) {
	return sonyflake.New(sonyflake.Settings{
		StartTime: _startTime,
		MachineID: func() (uint16, error) {
			return node, nil
		},
		CheckMachineID: func(id uint16) bool {
			return true
		},
	})
}

func NewNodeWithSetting(st *sonyflake.Settings) (*sonyflake.Sonyflake, error) {
	return sonyflake.New(*st)
}

// SetNode 设置节点
func SetNode(node *sonyflake.Sonyflake) {
	_idNode = node
}

func ID() uint64 {
	u, _ := _idNode.NextID()
	return u
}

func NextID() (uint64, error) {
	return _idNode.NextID()
}
