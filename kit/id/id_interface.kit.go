package idpkg

import (
	stderrors "errors"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
)

const (
	snowflakeNodeBits uint8 = 12
	snowflakeStepBits uint8 = 8
	snowflakeMaxNode        = int64(1<<snowflakeNodeBits - 1)
)

var (
	DefaultEpoch = time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local)
)

type Snowflake interface {
	NextID() (uint64, error)
}

type bwmarrinSnowflake struct {
	node *snowflake.Node
}

func NewBwmarrinSnowflake(nodeid int64) (Snowflake, error) {
	snowflake.Epoch = DefaultEpoch.UnixMilli()
	snowflake.NodeBits = snowflakeNodeBits
	snowflake.StepBits = snowflakeStepBits
	if nodeid < 0 || nodeid > snowflakeMaxNode {
		return nil, stderrors.New("Node number must be between 0 and " + strconv.FormatInt(snowflakeMaxNode, 10))
	}
	node, err := snowflake.NewNode(nodeid)
	if err != nil {
		return nil, err
	}
	return &bwmarrinSnowflake{node: node}, nil
}

func (s *bwmarrinSnowflake) NextID() (uint64, error) {
	si := s.node.Generate()
	return uint64(si.Int64()), nil
}
