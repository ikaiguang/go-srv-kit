package uuidpkg

import (
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

// New ...
func New() string {
	return NewXID()
}

// NewXID returns a globally unique XID string.
func NewXID() string {
	return xid.New().String()
}

// Deprecated: NewUUID returned an XID, not a UUID. Use NewXID instead.
func NewUUID() string { return NewXID() }

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

// NewRandomUUID returns a random RFC 4122 UUID string.
func NewRandomUUID() string {
	return uuid.NewString()
}

// Deprecated: use NewRandomUUID instead.
func UUID() string { return NewRandomUUID() }
