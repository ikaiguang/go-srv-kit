package sortutil

import "sort"

// pkg func
// sort.Ints([]int)
// sort.Float64s(a []float64)
// sort.Strings(a []string)

// Int32s same as sort.Ints
func Int32s(s []int32) {
	sort.Sort(Int32Slice(s))
}

// Int64s same as sort.Ints
func Int64s(s []int64) {
	sort.Sort(Int64Slice(s))
}

type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
