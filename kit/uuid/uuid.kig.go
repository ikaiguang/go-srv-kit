package uuidutil

import (
	"github.com/rs/xid"
	"time"
)

// New ...
func New() string {
	return xid.New().String()
}

// NewWithTime ...
func NewWithTime(t time.Time) string {
	return xid.NewWithTime(t).String()
}

// ID ...
func ID() xid.ID {
	return xid.New()
}

// IDWithTime ...
func IDWithTime(t time.Time) xid.ID {
	return xid.NewWithTime(t)
}

// FromString ...
func FromString(id string) (xid.ID, error) {
	return xid.FromString(id)
}

// FromBytes ...
func FromBytes(b []byte) (xid.ID, error) {
	return xid.FromBytes(b)
}

// Sort ...
func Sort(ids []xid.ID) {
	xid.Sort(ids)
}
