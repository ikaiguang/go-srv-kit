package randompkg

import (
	"math/rand"
	"strconv"
	"time"
)

// NewRandHandler ...
func NewRandHandler() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Name random name
func Name() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36) + "_" + Strings(5)
}

// Strings : A-Z a-z 0-9
func Strings(size int) string {

	randHandler := NewRandHandler()

	res := make([]byte, size)

	for i := 0; i < size; i++ {

		t := randHandler.Intn(3)

		if t == 0 {
			// A-Z
			res[i] = byte(randHandler.Intn(26) + 65)
		} else if t == 1 {
			// a-z
			res[i] = byte(randHandler.Intn(26) + 97)
		} else {
			// 0-9
			res[i] = byte(randHandler.Intn(9) + 48)
		}
	}
	return string(res)
}

// Letter : A-Z a-z
func Letter(size int) string {

	randHandler := NewRandHandler()

	res := make([]byte, size)

	for i := 0; i < size; i++ {

		t := randHandler.Intn(2)

		if t == 0 {
			// A-Z
			res[i] = byte(randHandler.Intn(26) + 65)
		} else {
			// a-z
			res[i] = byte(randHandler.Intn(26) + 97)
		}
	}
	return string(res)
}

// Numeric 0-9
func Numeric(size int) string {

	randHandler := NewRandHandler()

	res := make([]byte, size)

	for i := 0; i < size; i++ {
		res[i] = byte(randHandler.Intn(9) + 48)
	}
	return string(res)
}

// Int32Between random number between min-max
func Int32Between(min, max int32) int32 {
	if min == max {
		return min
	}

	if min >= max {
		min, max = max, min
	}

	return NewRandHandler().Int31n(max-min) + min
}

// Int64Between random number between min-max
func Int64Between(min, max int64) int64 {
	if min == max {
		return min
	}

	if min >= max {
		min, max = max, min
	}

	return NewRandHandler().Int63n(max-min) + min
}

// NumericBetween random number between min-max
func NumericBetween(min, max int64) int64 {
	return Int64Between(min, max)
}

// IntBetween random number between min-max
func IntBetween(min, max int) int {
	if min == max {
		return min
	}

	if min >= max {
		min, max = max, min
	}

	return NewRandHandler().Intn(max-min) + min
}
