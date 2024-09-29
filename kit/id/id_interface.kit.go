package idpkg

import (
	stderrors "errors"
	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
	"math"
	"strconv"
	"time"
)

var (
	DefaultEpoch = time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)
)

type Snowflake interface {
	NextID() (uint64, error)
}

type sonySonyflake struct {
	node *sonyflake.Sonyflake
}

func NewSonySonyflake(nodeid int64) (Snowflake, error) {
	if nodeid > math.MaxUint16 {
		return nil, stderrors.New("Node number must be between 0 and " + strconv.FormatInt(math.MaxUint16, 10))
	}
	st := sonyflake.Settings{
		StartTime: DefaultEpoch,
		MachineID: func() (uint16, error) {
			return uint16(nodeid), nil
		},
	}
	node, err := sonyflake.New(st)
	if err != nil {
		return nil, err
	}
	return &sonySonyflake{node: node}, nil
}

func (s *sonySonyflake) NextID() (uint64, error) {
	return s.node.NextID()
}

type bwmarrinSnowflake struct {
	node *snowflake.Node
}

func NewBwmarrinSnowflake(nodeid int64) (Snowflake, error) {
	snowflake.Epoch = DefaultEpoch.UnixMilli()
	snowflake.NodeBits = 16
	snowflake.StepBits = 6
	if nodeid > math.MaxUint16 {
		return nil, stderrors.New("Node number must be between 0 and " + strconv.FormatInt(math.MaxUint16, 10))
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
