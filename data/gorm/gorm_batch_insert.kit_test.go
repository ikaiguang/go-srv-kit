package gormutil

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// go test -v ./data/gorm/ -count=1 -test.run=TestBatchInsert
func TestBatchInsert(t *testing.T) {
	var (
		now        = time.Now().Format(time.RFC3339)
		userTotal  = 10
		userModels = make([]*User, userTotal)
	)
	for i := 0; i < userTotal; i++ {
		userModels[i] = &User{
			Name: "user_" + now + "_" + strconv.Itoa(i),
			Age:  i + 1,
		}
	}

	err := BatchInsert(dbConn, UserSlice(userModels))
	require.Nil(t, err)
}
