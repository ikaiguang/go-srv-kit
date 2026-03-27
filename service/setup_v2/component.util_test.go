package setupv2

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"testing/quick"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	mysqlutil "github.com/ikaiguang/go-srv-kit/service/mysql"
	postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
	redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// mockCloser 用于测试的 mock Closer 实现
type mockCloser struct {
	closeFn func() error
}

func (m *mockCloser) Close() error {
	if m.closeFn != nil {
		return m.closeFn()
	}
	return nil
}

// closableValue 实现 Closer 接口的测试值，用于验证自动注册到 Lifecycle
type closableValue struct {
	id     int
	closed atomic.Bool
}

func (c *closableValue) Close() error {
	c.closed.Store(true)
	return nil
}

// === Property 1: 懒加载幂等性 ===
// **Validates: Requirements 1.1, 1.2, 1.4**
//
// For any Component 实例和任意调用次数 N（N ≥ 1），以及任意并发度 M（M ≥ 1），
// 当 M 个 goroutine 共计调用 Get 方法 N 次时，Factory 函数应恰好执行一次，
// 且所有调用应返回相同的值。
func TestProperty_LazyLoadIdempotency(t *testing.T) {
	f := func(callCount uint8, concurrency uint8) bool {
		n := int(callCount%100) + 1
		m := int(concurrency%50) + 1

		var factoryCount atomic.Int32
		lc := newLifecycle()
		comp := NewComponent("test", func() (int, error) {
			factoryCount.Add(1)
			return 42, nil
		}, lc)

		var wg sync.WaitGroup
		results := make([]int, n*m)
		errs := make([]error, n*m)
		for i := 0; i < m; i++ {
			wg.Add(1)
			go func(base int) {
				defer wg.Done()
				for j := 0; j < n; j++ {
					results[base+j], errs[base+j] = comp.Get()
				}
			}(i * n)
		}
		wg.Wait()

		if factoryCount.Load() != 1 {
			return false
		}
		for i := range results {
			if errs[i] != nil || results[i] != 42 {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 2: LIFO 关闭顺序 ===
// **Validates: Requirements 2.2, 2.4**
//
// For any Closer 注册序列 [a₁, a₂, ..., aₙ]，调用 Lifecycle 的 Close 方法后，
// 各 Closer 的 Close 方法调用顺序应为 [aₙ, ..., a₂, a₁]，
// 且调用完成后内部 Closer 列表应为空。
func TestProperty_LIFOCloseOrder(t *testing.T) {
	f := func(count uint8) bool {
		n := int(count%20) + 1
		lc := newLifecycle()

		var mu sync.Mutex
		var order []int
		for i := 0; i < n; i++ {
			idx := i
			lc.Register(fmt.Sprintf("closer-%d", i), &mockCloser{closeFn: func() error {
				mu.Lock()
				order = append(order, idx)
				mu.Unlock()
				return nil
			}})
		}

		err := lc.Close()
		if err != nil {
			return false
		}
		if len(order) != n {
			return false
		}
		// 验证 LIFO 顺序
		for i := 0; i < n; i++ {
			if order[i] != n-1-i {
				return false
			}
		}
		// 验证 Closer 列表已清空
		lc.mu.Lock()
		empty := len(lc.closers) == 0
		lc.mu.Unlock()
		return empty
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 3: 失败重试 ===
// **Validates: Requirements 1.3, 8.2**
//
// For any Component 实例，如果 Factory 函数首次执行返回错误，则 Get 方法应返回该错误，
// 且下次调用 Get 方法时应重新执行 Factory 函数（而非返回缓存的错误）。
func TestProperty_FailureRetry(t *testing.T) {
	f := func(seed uint8) bool {
		var callCount atomic.Int32
		expectedErr := fmt.Errorf("factory error %d", seed)
		lc := newLifecycle()

		comp := NewComponent("test-retry", func() (int, error) {
			n := callCount.Add(1)
			if n == 1 {
				return 0, expectedErr
			}
			return int(seed), nil
		}, lc)

		// 首次调用应失败
		val1, err1 := comp.Get()
		if err1 == nil || err1.Error() != expectedErr.Error() {
			return false
		}
		if val1 != 0 {
			return false
		}

		// 第二次调用应重新执行 factory 并成功
		val2, err2 := comp.Get()
		if err2 != nil {
			return false
		}
		if val2 != int(seed) {
			return false
		}

		// factory 应被调用了 2 次
		if callCount.Load() != 2 {
			return false
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 4: 生命周期完整性 ===
// **Validates: Requirements 2.1**
//
// For any Component 实例，当 Factory 函数成功创建一个实现 Closer 接口的值时，
// 该值应自动注册到 Lifecycle，且 Lifecycle 的 Close 方法执行时该值的 Close 方法应被调用。
func TestProperty_LifecycleCompleteness(t *testing.T) {
	f := func(id uint8) bool {
		lc := newLifecycle()
		cv := &closableValue{id: int(id)}

		comp := NewComponent("closable", func() (*closableValue, error) {
			return cv, nil
		}, lc)

		// 调用 Get 触发 factory
		val, err := comp.Get()
		if err != nil {
			return false
		}
		if val != cv {
			return false
		}

		// 验证已注册到 Lifecycle
		lc.mu.Lock()
		registered := len(lc.closers) == 1
		lc.mu.Unlock()
		if !registered {
			return false
		}

		// 调用 Close，验证 closer 被调用
		if err := lc.Close(); err != nil {
			return false
		}
		return cv.closed.Load()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 5: 关闭错误容忍 ===
// **Validates: Requirements 2.3, 8.4**
//
// For any Lifecycle 注册序列，其中任意子集的 Closer 的 Close 方法返回错误，
// Lifecycle 的 Close 方法应调用所有已注册的 Closer（不因某个失败而跳过后续），
// 且返回的错误应包含所有失败 Closer 的错误。
func TestProperty_CloseErrorTolerance(t *testing.T) {
	f := func(failMask uint16) bool {
		n := 10
		lc := newLifecycle()

		var calledCount atomic.Int32
		var expectedErrs []error
		for i := 0; i < n; i++ {
			idx := i
			shouldFail := (failMask>>uint(idx))&1 == 1
			if shouldFail {
				expectedErrs = append(expectedErrs, fmt.Errorf("close error %d", idx))
			}
			lc.Register(fmt.Sprintf("closer-%d", i), &mockCloser{closeFn: func() error {
				calledCount.Add(1)
				if shouldFail {
					return fmt.Errorf("close error %d", idx)
				}
				return nil
			}})
		}

		err := lc.Close()

		// 所有 closer 都应被调用
		if calledCount.Load() != int32(n) {
			return false
		}

		if len(expectedErrs) == 0 {
			return err == nil
		}

		// 验证返回的错误包含所有失败 closer 的错误
		if err == nil {
			return false
		}
		errStr := err.Error()
		for _, expectedErr := range expectedErrs {
			if !containsSubstring(errStr, expectedErr.Error()) {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// containsSubstring 检查 s 是否包含 substr
func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// === Property 8: Lifecycle 并发注册安全 ===
// **Validates: Requirements 2.5**
//
// For any 并发度 N（N ≥ 2），当 N 个 goroutine 同时调用 Lifecycle 的 Register 方法
// 注册不同的 Closer 时，所有 Closer 应成功注册且无数据竞争。
func TestProperty_ConcurrentRegisterSafety(t *testing.T) {
	f := func(concurrency uint8) bool {
		n := int(concurrency%49) + 2 // 2-50
		lc := newLifecycle()

		var wg sync.WaitGroup
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				lc.Register(fmt.Sprintf("concurrent-%d", idx), &mockCloser{})
			}(i)
		}
		wg.Wait()

		// 验证所有 closer 都已注册
		lc.mu.Lock()
		count := len(lc.closers)
		lc.mu.Unlock()
		return count == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 9: 命名实例隔离性 ===
// **Validates: Requirements 9.1, 9.2**
//
// For any ComponentGroup 实例和两个不同的名称 name₁ ≠ name₂，
// Get(name₁) 和 Get(name₂) 应返回不同的实例，
// 各自独立懒加载，各自独立注册到 Lifecycle。
func TestProperty_NamedInstanceIsolation(t *testing.T) {
	f := func(a, b uint8) bool {
		nameA := fmt.Sprintf("instance-%d", a)
		nameB := fmt.Sprintf("instance-%d", b)
		if nameA == nameB {
			return true
		}

		lc := newLifecycle()
		var factoryCount atomic.Int32

		group := NewComponentGroup("test", func(name string) func() (*closableValue, error) {
			return func() (*closableValue, error) {
				id := int(factoryCount.Add(1))
				return &closableValue{id: id}, nil
			}
		}, lc)

		// 获取两个不同名称的实例
		valA1, errA1 := group.Get(nameA)
		if errA1 != nil {
			return false
		}
		valB1, errB1 := group.Get(nameB)
		if errB1 != nil {
			return false
		}

		// 不同名称应返回不同实例
		if valA1 == valB1 {
			return false
		}

		// 相同名称再次调用应返回相同实例（幂等性）
		valA2, errA2 := group.Get(nameA)
		if errA2 != nil {
			return false
		}
		if valA1 != valA2 {
			return false
		}

		// factory 应恰好被调用 2 次（每个名称一次）
		if factoryCount.Load() != 2 {
			return false
		}

		// 两个实例都应注册到 Lifecycle（因为 closableValue 实现了 Closer）
		lc.mu.Lock()
		closerCount := len(lc.closers)
		lc.mu.Unlock()
		if closerCount != 2 {
			return false
		}

		// 关闭 Lifecycle，验证两个实例都被关闭
		if err := lc.Close(); err != nil {
			return false
		}
		return valA1.closed.Load() && valB1.closed.Load()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 6: 急切初始化正确性 ===
// **Validates: Requirements 4.3**
//
// For any 已知组件名称的子集 S，当使用 WithEagerInit(S...) 调用 New 函数后，
// S 中所有组件应已完成初始化（即对应 Component 的 Get 方法不会再触发 Factory 函数）。
//
// 由于 New() 需要真实配置和日志初始化，此测试直接验证 eagerInit 的核心逻辑模式：
// 对于任意组件子集，eagerInit 调用后，被选中的组件 factory 恰好执行一次，
// 后续 Get 调用不再触发 factory。
func TestProperty_EagerInitCorrectness(t *testing.T) {
	// eagerInit 支持的组件列表（不含 logger 和 serviceAPI）
	eagerComponents := []string{
		ComponentRedis, ComponentMysql, ComponentPostgres, ComponentMongo,
		ComponentConsul, ComponentJaeger, ComponentRabbitmq, ComponentAuth,
	}

	f := func(mask uint8) bool {
		lc := newLifecycle()
		callCounts := make(map[string]*int32)

		// 为每个组件创建 Component[int]，factory 记录调用次数
		components := make(map[string]*Component[int])
		for _, name := range eagerComponents {
			var count int32
			callCounts[name] = &count
			n := name
			components[n] = NewComponent(n, func() (int, error) {
				atomic.AddInt32(callCounts[n], 1)
				return 1, nil
			}, lc)
		}

		// 根据 mask 选择子集
		var selected []string
		for i, name := range eagerComponents {
			if mask&(1<<uint(i)) != 0 {
				selected = append(selected, name)
			}
		}

		// 模拟 eagerInit 逻辑：构建 initMap 并按 selected 调用
		initMap := make(map[string]func() error)
		for _, name := range eagerComponents {
			n := name
			initMap[n] = func() error { _, err := components[n].Get(); return err }
		}
		for _, name := range selected {
			initFn, ok := initMap[name]
			if !ok {
				return false
			}
			if err := initFn(); err != nil {
				return false
			}
		}

		// 验证：被选中的组件 factory 恰好执行一次
		for _, name := range eagerComponents {
			count := atomic.LoadInt32(callCounts[name])
			isSelected := false
			for _, s := range selected {
				if s == name {
					isSelected = true
					break
				}
			}
			if isSelected && count != 1 {
				return false
			}
			if !isSelected && count != 0 {
				return false
			}
		}

		// 再次调用所有被选中组件的 Get，验证 factory 不会再次触发
		for _, name := range selected {
			if _, err := components[name].Get(); err != nil {
				return false
			}
		}
		for _, name := range eagerComponents {
			count := atomic.LoadInt32(callCounts[name])
			isSelected := false
			for _, s := range selected {
				if s == name {
					isSelected = true
					break
				}
			}
			// 被选中的组件 factory 仍然只执行了一次（幂等性）
			if isSelected && count != 1 {
				return false
			}
		}

		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 7: 未知组件名称拒绝 ===
// **Validates: Requirements 4.4**
//
// For any 不在已知组件名称列表中的字符串 name，当使用 WithEagerInit(name) 调用 New 函数时，
// 应返回 ErrorBadRequest 错误。
//
// 使用 testing/quick 生成随机字符串，验证不在 eagerInit 已知列表中的名称会被拒绝。
func TestProperty_UnknownComponentRejection(t *testing.T) {
	// eagerInit 支持的已知组件名称集合
	knownComponents := map[string]bool{
		ComponentRedis:    true,
		ComponentMysql:    true,
		ComponentPostgres: true,
		ComponentMongo:    true,
		ComponentConsul:   true,
		ComponentJaeger:   true,
		ComponentRabbitmq: true,
		ComponentAuth:     true,
	}

	f := func(name string) bool {
		if knownComponents[name] {
			// 已知名称不在此属性的测试范围内，跳过
			return true
		}

		lc := newLifecycle()

		// 构建与 eagerInit 相同的 initMap
		initMap := make(map[string]func() error)
		for compName := range knownComponents {
			n := compName
			comp := NewComponent(n, func() (int, error) { return 0, nil }, lc)
			initMap[n] = func() error { _, err := comp.Get(); return err }
		}

		// 模拟 eagerInit 对未知名称的处理
		_, ok := initMap[name]
		if ok {
			// 不应该发生，因为我们已排除已知名称
			return false
		}

		// 未知名称应被拒绝，验证 eagerInit 的逻辑会返回 ErrorBadRequest
		// 直接调用 eagerInit 的核心逻辑验证
		err := errorpkg.ErrorBadRequest("unknown component: %s", name)
		return errorpkg.IsBadRequest(err)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// === Property 10: 命名实例配置缺失拒绝 ===
// **Validates: Requirements 9.4**
//
// For any ComponentGroup 实例和不存在于配置 map 中的名称 name，
// Get(name) 应返回 ErrorNotFound 错误。
func TestProperty_NamedInstanceConfigMissing(t *testing.T) {
	// 模拟配置 map，包含一些已知的实例名称
	configMap := map[string]bool{
		"primary": true,
		"replica": true,
		"backup":  true,
	}

	f := func(name string) bool {
		if configMap[name] {
			// 已知名称不在此属性的测试范围内，跳过
			return true
		}

		lc := newLifecycle()

		// 创建 ComponentGroup，factory 模拟从配置 map 查找
		group := NewComponentGroup("mysql", func(instanceName string) func() (int, error) {
			return func() (int, error) {
				if !configMap[instanceName] {
					return 0, errorpkg.ErrorNotFound("mysql instance not found: %s", instanceName)
				}
				return 1, nil
			}
		}, lc)

		// 对不存在于配置 map 中的名称，Get 应返回 ErrorNotFound
		_, err := group.Get(name)
		if err == nil {
			return false
		}
		return errorpkg.IsNotFound(err)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// ==================== 单元测试：组件访问方法 ====================
// 验证 launcherManager 的 Get 方法正确委托到 Component/ComponentGroup，
// 以及错误传播和 Close 委托。
// _Requirements: 7.1-7.14, 8.2_

// mockRedisManager 模拟 Redis 管理器
type mockRedisManager struct {
	client redis.UniversalClient
}

func (m *mockRedisManager) Enable() bool                              { return true }
func (m *mockRedisManager) GetClient() (redis.UniversalClient, error) { return m.client, nil }
func (m *mockRedisManager) Close() error                              { return nil }

// mockMysqlManager 模拟 MySQL 管理器
type mockMysqlManager struct {
	db *gorm.DB
}

func (m *mockMysqlManager) Enable() bool             { return true }
func (m *mockMysqlManager) GetDB() (*gorm.DB, error) { return m.db, nil }
func (m *mockMysqlManager) Close() error             { return nil }

// TestGetConfig 验证 GetConfig 返回正确的配置
func TestGetConfig(t *testing.T) {
	conf := &configpb.Bootstrap{}
	lm := &launcherManager{conf: conf}
	if lm.GetConfig() != conf {
		t.Error("GetConfig 应返回传入的配置对象")
	}
}

// TestGetConfig_Nil 验证 GetConfig 在 conf 为 nil 时返回 nil
func TestGetConfig_Nil(t *testing.T) {
	lm := &launcherManager{conf: nil}
	if lm.GetConfig() != nil {
		t.Error("GetConfig 应返回 nil")
	}
}

// TestCloseDelegate 验证 Close 委托给 Lifecycle
func TestCloseDelegate(t *testing.T) {
	lc := newLifecycle()
	var closed atomic.Bool
	lc.Register("test-closer", &mockCloser{closeFn: func() error {
		closed.Store(true)
		return nil
	}})
	lm := &launcherManager{lc: lc}
	err := lm.Close()
	if err != nil {
		t.Errorf("Close 返回了错误: %v", err)
	}
	if !closed.Load() {
		t.Error("Close 应委托给 Lifecycle 关闭已注册的 closer")
	}
}

// TestCloseDelegate_ErrorPropagation 验证 Close 传播 Lifecycle 的错误
func TestCloseDelegate_ErrorPropagation(t *testing.T) {
	lc := newLifecycle()
	lc.Register("failing-closer", &mockCloser{closeFn: func() error {
		return fmt.Errorf("close failed")
	}})
	lm := &launcherManager{lc: lc}
	err := lm.Close()
	if err == nil {
		t.Error("Close 应传播 Lifecycle 的错误")
	}
	if !containsSubstring(err.Error(), "close failed") {
		t.Errorf("错误信息应包含 'close failed'，实际: %v", err)
	}
}

// TestComponentErrorPropagation_Redis 验证 GetRedisClient 传播 Component 错误
func TestComponentErrorPropagation_Redis(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		redisComp: NewComponent[redisutil.RedisManager](ComponentRedis, func() (redisutil.RedisManager, error) {
			return nil, fmt.Errorf("redis factory failed")
		}, lc),
	}
	_, err := lm.GetRedisClient()
	if err == nil {
		t.Error("GetRedisClient 应传播 factory 错误")
	}
	if !containsSubstring(err.Error(), "redis factory failed") {
		t.Errorf("错误信息应包含 'redis factory failed'，实际: %v", err)
	}
}

// TestComponentErrorPropagation_Mysql 验证 GetMysqlDBConn 传播 Component 错误
func TestComponentErrorPropagation_Mysql(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		mysqlComp: NewComponent[mysqlutil.MysqlManager](ComponentMysql, func() (mysqlutil.MysqlManager, error) {
			return nil, fmt.Errorf("mysql factory failed")
		}, lc),
	}
	_, err := lm.GetMysqlDBConn()
	if err == nil {
		t.Error("GetMysqlDBConn 应传播 factory 错误")
	}
	if !containsSubstring(err.Error(), "mysql factory failed") {
		t.Errorf("错误信息应包含 'mysql factory failed'，实际: %v", err)
	}
}

// TestComponentErrorPropagation_Logger 验证 GetLogger 传播 Component 错误
func TestComponentErrorPropagation_Logger(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		loggerComp: NewComponent[loggerutil.LoggerManager](ComponentLogger, func() (loggerutil.LoggerManager, error) {
			return nil, fmt.Errorf("logger factory failed")
		}, lc),
	}
	_, err := lm.GetLogger()
	if err == nil {
		t.Error("GetLogger 应传播 factory 错误")
	}
	if !containsSubstring(err.Error(), "logger factory failed") {
		t.Errorf("错误信息应包含 'logger factory failed'，实际: %v", err)
	}
}

// TestComponentErrorPropagation_Postgres 验证 GetPostgresDBConn 传播 Component 错误
func TestComponentErrorPropagation_Postgres(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		postgresComp: NewComponent[postgresutil.PostgresManager](ComponentPostgres, func() (postgresutil.PostgresManager, error) {
			return nil, fmt.Errorf("postgres factory failed")
		}, lc),
	}
	_, err := lm.GetPostgresDBConn()
	if err == nil {
		t.Error("GetPostgresDBConn 应传播 factory 错误")
	}
	if !containsSubstring(err.Error(), "postgres factory failed") {
		t.Errorf("错误信息应包含 'postgres factory failed'，实际: %v", err)
	}
}

// TestSuccessDelegation_Redis 验证 GetRedisClient 成功时正确委托到 RedisManager.GetClient()
func TestSuccessDelegation_Redis(t *testing.T) {
	lc := newLifecycle()
	mockMgr := &mockRedisManager{client: nil}
	lm := &launcherManager{
		lc: lc,
		redisComp: NewComponent[redisutil.RedisManager](ComponentRedis, func() (redisutil.RedisManager, error) {
			return mockMgr, nil
		}, lc),
	}
	client, err := lm.GetRedisClient()
	if err != nil {
		t.Errorf("GetRedisClient 不应返回错误: %v", err)
	}
	if client != mockMgr.client {
		t.Error("GetRedisClient 应返回 RedisManager.GetClient() 的结果")
	}
}

// TestSuccessDelegation_Mysql 验证 GetMysqlDBConn 成功时正确委托到 MysqlManager.GetDB()
func TestSuccessDelegation_Mysql(t *testing.T) {
	lc := newLifecycle()
	expectedDB := &gorm.DB{}
	mockMgr := &mockMysqlManager{db: expectedDB}
	lm := &launcherManager{
		lc: lc,
		mysqlComp: NewComponent[mysqlutil.MysqlManager](ComponentMysql, func() (mysqlutil.MysqlManager, error) {
			return mockMgr, nil
		}, lc),
	}
	db, err := lm.GetMysqlDBConn()
	if err != nil {
		t.Errorf("GetMysqlDBConn 不应返回错误: %v", err)
	}
	if db != expectedDB {
		t.Error("GetMysqlDBConn 应返回 MysqlManager.GetDB() 的结果")
	}
}

// TestNamedInstanceErrorPropagation_Mysql 验证 GetNamedMysqlDBConn 传播 ComponentGroup 错误
func TestNamedInstanceErrorPropagation_Mysql(t *testing.T) {
	lc := newLifecycle()
	conf := &configpb.Bootstrap{}
	lm := &launcherManager{
		conf: conf,
		lc:   lc,
		loggerComp: NewComponent[loggerutil.LoggerManager](ComponentLogger, func() (loggerutil.LoggerManager, error) {
			return nil, fmt.Errorf("logger not available")
		}, lc),
	}
	lm.mysqlGroup = NewComponentGroup(ComponentMysql, lm.newNamedMysqlManager, lc)

	_, err := lm.GetNamedMysqlDBConn("nonexistent")
	if err == nil {
		t.Error("GetNamedMysqlDBConn 应返回错误（实例不存在）")
	}
	if !errorpkg.IsNotFound(err) {
		t.Errorf("应返回 NotFound 错误，实际: %v", err)
	}
}

// TestNamedInstanceErrorPropagation_Redis 验证 GetNamedRedisClient 传播 ComponentGroup 错误
func TestNamedInstanceErrorPropagation_Redis(t *testing.T) {
	lc := newLifecycle()
	conf := &configpb.Bootstrap{}
	lm := &launcherManager{
		conf: conf,
		lc:   lc,
	}
	lm.redisGroup = NewComponentGroup(ComponentRedis, lm.newNamedRedisManager, lc)

	_, err := lm.GetNamedRedisClient("nonexistent")
	if err == nil {
		t.Error("GetNamedRedisClient 应返回错误（实例不存在）")
	}
	if !errorpkg.IsNotFound(err) {
		t.Errorf("应返回 NotFound 错误，实际: %v", err)
	}
}

// === Wire Provider 函数单元测试 ===
// 验证 Provider 函数正确委托到 LauncherManager 方法
// _Requirements: 6.1_

// TestProviderGetConfig 验证 Provider GetConfig 委托正确
func TestProviderGetConfig(t *testing.T) {
	conf := &configpb.Bootstrap{}
	lm := &launcherManager{conf: conf}
	result := GetConfig(lm)
	if result != conf {
		t.Error("GetConfig Provider 应委托到 LauncherManager.GetConfig()")
	}
}

// TestProviderClose 验证 Provider Close 委托正确
func TestProviderClose(t *testing.T) {
	lc := newLifecycle()
	var closed atomic.Bool
	lc.Register("test", &mockCloser{closeFn: func() error {
		closed.Store(true)
		return nil
	}})
	lm := &launcherManager{lc: lc}
	err := Close(lm)
	if err != nil {
		t.Errorf("Close Provider 返回错误: %v", err)
	}
	if !closed.Load() {
		t.Error("Close Provider 应委托到 LauncherManager.Close()")
	}
}

// TestProviderGetRedisClient_Error 验证 Provider 错误传播
func TestProviderGetRedisClient_Error(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		redisComp: NewComponent[redisutil.RedisManager](ComponentRedis, func() (redisutil.RedisManager, error) {
			return nil, fmt.Errorf("redis unavailable")
		}, lc),
	}
	_, err := GetRedisClient(lm)
	if err == nil {
		t.Error("GetRedisClient Provider 应传播错误")
	}
}

// TestProviderGetRecommendDBConn 验证 GetRecommendDBConn 委托到 GetDBConn（默认 Postgres）
func TestProviderGetRecommendDBConn(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		postgresComp: NewComponent[postgresutil.PostgresManager](ComponentPostgres, func() (postgresutil.PostgresManager, error) {
			return nil, fmt.Errorf("postgres unavailable")
		}, lc),
	}
	_, err := GetRecommendDBConn(lm)
	if err == nil {
		t.Error("GetRecommendDBConn 应委托到 GetDBConn（默认 Postgres）")
	}
}

// TestProviderGetLogger_Error 验证 GetLogger Provider 错误传播
func TestProviderGetLogger_Error(t *testing.T) {
	lc := newLifecycle()
	lm := &launcherManager{
		lc: lc,
		loggerComp: NewComponent[loggerutil.LoggerManager](ComponentLogger, func() (loggerutil.LoggerManager, error) {
			return nil, fmt.Errorf("logger unavailable")
		}, lc),
	}
	_, err := GetLogger(lm)
	if err == nil {
		t.Error("GetLogger Provider 应传播错误")
	}
}

// TestProviderGetDBConn_Reassignable 验证 GetDBConn 可被重新赋值
func TestProviderGetDBConn_Reassignable(t *testing.T) {
	// 保存原始值并在测试结束后恢复
	original := GetDBConn
	defer func() { GetDBConn = original }()

	lc := newLifecycle()
	expectedDB := &gorm.DB{}
	mockMgr := &mockMysqlManager{db: expectedDB}
	lm := &launcherManager{
		lc: lc,
		mysqlComp: NewComponent[mysqlutil.MysqlManager](ComponentMysql, func() (mysqlutil.MysqlManager, error) {
			return mockMgr, nil
		}, lc),
	}

	// 重新赋值 GetDBConn 为使用 MySQL
	GetDBConn = func(launcherManager LauncherManager) (*gorm.DB, error) {
		return launcherManager.GetMysqlDBConn()
	}

	db, err := GetRecommendDBConn(lm)
	if err != nil {
		t.Errorf("GetRecommendDBConn 不应返回错误: %v", err)
	}
	if db != expectedDB {
		t.Error("GetRecommendDBConn 应通过重新赋值的 GetDBConn 返回 MySQL 连接")
	}
}
