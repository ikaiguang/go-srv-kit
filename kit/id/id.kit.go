package idpkg

var (
	// Node 生成ID的节点
	// 为了帮助保证唯一性
	// - 确保您的系统保持准确的系统时间
	// - 确保您永远不会有多个节点以相同的节点 ID 运行
	Node Snowflake
)

// ===== Benchmark =====
// BenchmarkNew_SonySonyflake-8			31080             38894 ns/op               0 B/op          0 allocs/op
// BenchmarkNew_BwmarrinSnowflake-8		76981             15611 ns/op               0 B/op          0 allocs/op
// ===== Benchmark =====
func init() {
	var err error
	Node, err = NewBwmarrinSnowflake(1)
	if err != nil {
		panic(err)
	}
}

func SetNode(node Snowflake) {
	Node = node
}

// NextID ...
// 为了帮助保证唯一性
// - 确保您的系统保持准确的系统时间
// - 确保您永远不会有多个节点以相同的节点 ID 运行
func NextID() (uint64, error) {
	return Node.NextID()
}
