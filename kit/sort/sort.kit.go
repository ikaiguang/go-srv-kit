package sortpkg

// pkg func
// sort.Ints([]int)
// sort.Float64s(a []float64)
// sort.Strings(a []string)

// Deprecated: 使用 Sort[int32] 替代
func Int32s(s []int32) {
	Sort(s)
}

// Deprecated: 使用 Sort[int64] 替代
func Int64s(s []int64) {
	Sort(s)
}

// Deprecated: 使用 Sort[int32] 替代
type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Deprecated: 使用 Sort[int64] 替代
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
