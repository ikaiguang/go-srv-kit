package setuputil

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"testing/quick"
)

// === Property 11: 失败并发无零值泄露 ===
// **Validates: Requirements 1.3, 8.2 的加强版**
//
// 场景：factory 首次调用失败，并发多个 goroutine 同时调用 Get()。
// 原实现存在 data race / 可能返回 (零值, nil) 的 bug：
//   - Goroutine A 进入 once.Do，factory 失败，value 保持零值，但 once 已标记完成
//   - A 重置 once 之前，Goroutine B 调用 Get 时 once.Do 跳过，直接返回 (零值, nil)
//
// 本测试验证：在失败 + 并发场景下，Get() 永远不返回 (零值, nil) 的组合。
//
// 语义保证：
//   - 成功调用必须返回非零业务值 + nil error
//   - 失败调用必须返回零值 + 非 nil error
//   - 不允许出现 (零值, nil) 或 (非零值, 非 nil) 的组合
func TestProperty_FailureConcurrentNoZeroValue(t *testing.T) {
	f := func(concurrency uint8, failTimes uint8) bool {
		m := int(concurrency%32) + 2  // 2-33 个 goroutine
		fails := int(failTimes%5) + 1 // 1-5 次失败后成功

		var callCount atomic.Int32
		lc := newLifecycle()
		comp := NewComponent("failure-concurrent", func() (int, error) {
			n := callCount.Add(1)
			if int(n) <= fails {
				return 0, fmt.Errorf("factory fail #%d", n)
			}
			return 42, nil // 成功后返回固定非零值
		}, lc)

		var wg sync.WaitGroup
		// 重复多轮并发调用，直到所有调用都成功（或检测到异常）
		for round := 0; round < fails+2; round++ {
			errSeen := make([]error, m)
			valSeen := make([]int, m)
			for i := 0; i < m; i++ {
				wg.Add(1)
				go func(idx int) {
					defer wg.Done()
					valSeen[idx], errSeen[idx] = comp.Get()
				}(i)
			}
			wg.Wait()

			for i := 0; i < m; i++ {
				// 违规 1：成功但返回零值
				if errSeen[i] == nil && valSeen[i] == 0 {
					t.Logf("round %d idx %d: got (0, nil), violates invariant", round, i)
					return false
				}
				// 违规 2：失败但返回非零值
				if errSeen[i] != nil && valSeen[i] != 0 {
					t.Logf("round %d idx %d: got (%d, %v), violates invariant", round, i, valSeen[i], errSeen[i])
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 200}); err != nil {
		t.Error(err)
	}
}

// === Property 12: 成功发布原子性 ===
// **Validates: 方案 A 的新语义**
//
// 验证 Component 一旦通过 factory 成功初始化，
//   - 后续任意并发调用 Get() 都必须返回相同的成功值
//   - factory 不会被再次触发
//   - 不会返回错误
func TestProperty_SuccessPublishAtomicity(t *testing.T) {
	f := func(concurrency uint8) bool {
		m := int(concurrency%48) + 2
		var callCount atomic.Int32
		lc := newLifecycle()

		// 使用指针确保每次 factory 创建不同对象，便于验证单一发布
		comp := NewComponent("success-atomicity", func() (*closableValue, error) {
			n := callCount.Add(1)
			return &closableValue{id: int(n)}, nil
		}, lc)

		var wg sync.WaitGroup
		results := make([]*closableValue, m)
		for i := 0; i < m; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				v, err := comp.Get()
				if err == nil {
					results[idx] = v
				}
			}(i)
		}
		wg.Wait()

		// factory 应只成功执行一次
		if callCount.Load() != 1 {
			return false
		}
		// 所有调用应返回同一对象
		var first *closableValue
		for _, r := range results {
			if r == nil {
				return false
			}
			if first == nil {
				first = r
				continue
			}
			if r != first {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
}

// === Property 13: ComponentGroup 并发命中已存在实例 ===
// **Validates: ComponentGroup 从 Mutex 升级为 RWMutex 的正确性**
//
// 验证：并发多个 goroutine 读取同一已存在命名实例时，
//   - 返回同一对象
//   - factory 只被调用一次
//   - 对不同名称调用不互相阻塞（通过高并发多名称模拟）
func TestProperty_ComponentGroupConcurrentHit(t *testing.T) {
	f := func(nameCount, concurrency uint8) bool {
		nn := int(nameCount%8) + 2   // 2-9 个名称
		m := int(concurrency%32) + 2 // 2-33 个 goroutine per name
		lc := newLifecycle()
		var totalFactory atomic.Int32

		group := NewComponentGroup("test", func(name string) func() (*closableValue, error) {
			return func() (*closableValue, error) {
				totalFactory.Add(1)
				return &closableValue{id: len(name)}, nil
			}
		}, lc)

		// 先顺序初始化所有名称
		firstInstances := make(map[string]*closableValue, nn)
		for i := 0; i < nn; i++ {
			name := fmt.Sprintf("name-%d", i)
			v, err := group.Get(name)
			if err != nil {
				return false
			}
			firstInstances[name] = v
		}
		if int(totalFactory.Load()) != nn {
			return false
		}

		// 并发从已初始化的实例读取
		var wg sync.WaitGroup
		for i := 0; i < nn; i++ {
			name := fmt.Sprintf("name-%d", i)
			expected := firstInstances[name]
			for j := 0; j < m; j++ {
				wg.Add(1)
				go func(nm string, exp *closableValue) {
					defer wg.Done()
					got, err := group.Get(nm)
					if err != nil || got != exp {
						t.Errorf("name %s concurrent hit failed: got=%p expected=%p err=%v", nm, got, exp, err)
					}
				}(name, expected)
			}
		}
		wg.Wait()

		// factory 调用次数不应增加
		return int(totalFactory.Load()) == nn
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 50}); err != nil {
		t.Error(err)
	}
}

// ==================== 基准测试 ====================

// BenchmarkComponent_Get_HotPath 测量已初始化组件的 Get 热路径性能
// 方案 A 使用 atomic.Pointer.Load，预期远快于 sync.Once.Do
func BenchmarkComponent_Get_HotPath(b *testing.B) {
	lc := newLifecycle()
	comp := NewComponent("bench-hotpath", func() (int, error) {
		return 42, nil
	}, lc)
	// 预热，确保已初始化
	if _, err := comp.Get(); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = comp.Get()
		}
	})
}

// BenchmarkComponentGroup_Get_HitExisting 测量命中已存在命名实例的性能
// 方案 A 使用 RWMutex 的 RLock，多读场景下远快于原 Mutex
func BenchmarkComponentGroup_Get_HitExisting(b *testing.B) {
	lc := newLifecycle()
	group := NewComponentGroup("bench", func(name string) func() (int, error) {
		return func() (int, error) { return len(name), nil }
	}, lc)

	// 预创建若干实例
	names := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, n := range names {
		if _, err := group.Get(n); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_, _ = group.Get(names[i%len(names)])
			i++
		}
	})
}
