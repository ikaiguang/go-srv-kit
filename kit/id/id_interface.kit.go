package idpkg

import (
	stderrors "errors"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

const (
	snowflakeNodeBits uint8 = 12
	snowflakeStepBits uint8 = 8
	snowflakeMaxNode        = int64(1<<snowflakeNodeBits - 1)
)

var (
	// use SetDefaultEpoch and DefaultEpochValue for synchronized access.
	defaultEpoch = time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local)

	snowflakeConfigMu sync.Mutex
)

type Snowflake interface {
	NextID() (uint64, error)
}

type bwmarrinSnowflake struct {
	node *snowflake.Node
}

func NewBwmarrinSnowflake(nodeid int64) (Snowflake, error) {
	if nodeid < 0 || nodeid > snowflakeMaxNode {
		return nil, stderrors.New("Node number must be between 0 and " + strconv.FormatInt(snowflakeMaxNode, 10))
	}
	snowflakeConfigMu.Lock()
	defer snowflakeConfigMu.Unlock()
	snowflake.Epoch = defaultEpoch.UnixMilli()
	snowflake.NodeBits = snowflakeNodeBits
	snowflake.StepBits = snowflakeStepBits
	node, err := snowflake.NewNode(nodeid)
	if err != nil {
		return nil, err
	}
	return &bwmarrinSnowflake{node: node}, nil
}

// SetDefaultEpoch updates the epoch used by subsequently created Snowflake nodes.
func SetDefaultEpoch(epoch time.Time) {
	snowflakeConfigMu.Lock()
	defer snowflakeConfigMu.Unlock()
	defaultEpoch = epoch
}

// DefaultEpochValue returns the epoch used by subsequently created Snowflake nodes.
func DefaultEpochValue() time.Time {
	snowflakeConfigMu.Lock()
	defer snowflakeConfigMu.Unlock()
	return defaultEpoch
}

func (s *bwmarrinSnowflake) NextID() (uint64, error) {
	si := s.node.Generate()
	return uint64(si.Int64()), nil
}
