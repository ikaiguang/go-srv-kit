package idutil

import (
	"sync"
)

const (
	_nodeID = 1
)

var (
	_idNode     interface{}
	_idNodeOnce sync.Once
)
