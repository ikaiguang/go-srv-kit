# Code Review 规范

## 审查原则

1. **友好建设性** - 提出改进建议，而非批评
2. **关注代码本身** - 对事不对人
3. **解释原因** - 说明为什么需要修改
4. **承认自己的无知** - 不确定时提问而非强求
5. **及时响应** - 尽快处理评论

## 通用审查清单

### 代码质量

- [ ] 代码是否符合 `coding-standards.md` 规范
- [ ] 命名是否清晰易懂
- [ ] 函数是否单一职责
- [ ] 函数长度是否超过 150 行（需要拆分）
- [ ] 是否有重复代码
- [ ] 是否有魔法数字或硬编码
- [ ] 注释是否准确且必要
- [ ] 第三方 API 调用是否复用（而非重复实现）

### 错误处理

- [ ] 是否正确处理了所有错误
- [ ] 是否使用 `kratos/error/` 包
- [ ] 错误信息是否清晰
- [ ] 是否有必要的错误日志

### 安全性

- [ ] 是否有 SQL 注入风险
- [ ] 是否有 XSS 风险
- [ ] 敏感信息是否脱敏
- [ ] 权限是否正确校验
- [ ] 输入参数是否验证

### 性能

- [ ] 是否有 N+1 查询
- [ ] 是否有不必要的循环嵌套
- [ ] 是否正确使用缓存
- [ ] 数据库查询是否可优化

### 测试

- [ ] 是否有单元测试
- [ ] 测试覆盖率是否足够
- [ ] 是否有边界情况测试
- [ ] 测试是否可独立运行

## 分层审查重点

### Proto 层

```protobuf
// ✅ 好的实践
syntax = "proto3";
package api.user.service.v1;
option go_package = "github.com/ikaiguang/go-srv-kit/api/user-service/v1;v1";

// 字段注释
message CreateUserReq {
  string username = 1 [(validate.rules).string.min_len = 1]; // 有验证规则
}
```

**审查清单：**
- [ ] 包名是否符合 `api.{service}/v1` 格式
- [ ] `go_package` option 是否正确
- [ ] 字段是否有注释
- [ ] 是否有验证规则（`validate.rules`）
- [ ] 错误定义是否在 `errors/` 目录

### Service Layer

```go
// ✅ 好的实践
func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
    // 1. 参数验证
    if req.GetUsername() == "" {
        return nil, errorpkg.ErrorBadRequest("username is required")
    }

    // 2. DTO → BO
    param := dto.ToBoCreateUserParam(req)

    // 3. 调用业务逻辑
    result, err := s.userBiz.CreateUser(ctx, param)
    if err != nil {
        logpkg.WithContext(ctx).Errorw("create user failed", "error", err)
        return nil, err
    }

    // 4. BO → Proto
    return dto.ToProtoCreateUserResp(result), nil
}
```

**审查清单：**
- [ ] 是否有参数验证
- [ ] DTO 转换是否正确
- [ ] 是否调用了 Business 层而非直接调用 Data 层
- [ ] 错误是否正确处理和记录日志
- [ ] 是否正确使用 Context

### Business Layer

```go
// ✅ 好的实践
func (b *userBiz) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    // 1. 业务验证
    exists, err := b.userRepo.CheckUserExists(ctx, param.Username)
    if err != nil {
        return nil, errorpkg.FormatError(err)
    }
    if exists {
        return nil, customErrors.UserAlreadyExists()
    }

    // 2. 调用 Data 层
    result, err := b.userRepo.CreateUser(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return result, nil
}
```

**审查清单：**
- [ ] 业务逻辑是否清晰
- [ ] 是否有业务规则验证
- [ ] 是否只调用 Repository 接口
- [ ] 事务处理是否正确
- [ ] 复杂逻辑是否有注释

### Data Layer

```go
// ✅ 好的实践
func (d *userData) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*bo.CreateUserResult, error) {
    user := &po.User{
        Username: param.Username,
        Email:    param.Email,
    }

    if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
        // 处理唯一约束冲突
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            return nil, customErrors.UserAlreadyExists()
        }
        return nil, errorpkg.FormatError(err)
    }

    return &bo.CreateUserResult{ID: user.ID}, nil
}
```

**审查清单：**
- [ ] 是否使用 `WithContext(ctx)`
- [ ] PO 模型是否正确定义
- [ ] 是否正确处理 GORM 错误
- [ ] 是否有 SQL 注入风险
- [ ] 批量操作是否使用 Batch
- [ ] 是否正确使用事务

### Wire 依赖注入

```go
// ✅ 好的实践
func exportServices(launcher setupv2.LauncherManager, hs *http.Server, gs *grpc.Server) {
    panic(wire.Build(
        // 基础设施
        setupv2.GetLogger,

        // Data 层
        data.NewUserData,

        // Business 层
        biz.NewUserBiz,

        // Service 层
        service.NewUserService,

        // 注册服务
        service.RegisterServices,
    ))
}
```

**审查清单：**
- [ ] 依赖顺序是否正确（从下到上）
- [ ] 接口是否使用 `wire.Bind`
- [ ] 是否有循环依赖
- [ ] 新增的依赖是否已添加
- [ ] 是否生成了 `wire_gen.go`

## 常见问题

### 架构违规

```go
// ❌ 错误：Service 直接调用 Data
func (s *userService) GetUser(ctx context.Context, id uint) (*bo.User, error) {
    return s.userData.GetUser(ctx, id)  // 违规！应该调用 Biz
}

// ✅ 正确
func (s *userService) GetUser(ctx context.Context, id uint) (*bo.User, error) {
    return s.userBiz.GetUser(ctx, id)
}
```

### 错误处理缺失

```go
// ❌ 错误：忽略错误
user, _ := s.userBiz.GetUser(ctx, id)

// ✅ 正确
user, err := s.userBiz.GetUser(ctx, id)
if err != nil {
    return nil, err
}
```

### Context 未传递

```go
// ❌ 错误：未传递 Context
users, _ := d.db.Find(&users).Error

// ✅ 正确
users, _ := d.db.WithContext(ctx).Find(&users).Error
```

### 硬编码配置

```go
// ❌ 错误：硬编码
timeout := 30 * time.Second

// ✅ 正确：从配置读取
timeout := time.Duration(config.GetTimeout()) * time.Second
```

### 敏感信息日志

```go
// ❌ 错误：记录密码
log.Infow("user login", "password", password)

// ✅ 正确：脱敏
log.Infow("user login", "password", stringutil.MaskPassword(password))
```

### 函数过长

**函数长度不能超过 150 行，超过则必须拆分。**

### 嵌套层级过深

**函数嵌套层级不能超过 3 层，超过则必须重构。**

#### 什么是嵌套层级？

嵌套层级是指代码中 `if`、`for`、`switch` 等控制结构的嵌套深度。

```go
// 0 层 - 平铺代码
func example() {
    statement1()
    statement2()
}

// 1 层
func example() {
    if condition {           // 第 1 层
        statement1()
    }
}

// 2 层
func example() {
    if condition {           // 第 1 层
        if condition2 {      // 第 2 层
            statement1()
        }
    }
}

// 3 层（允许的最大值）
func example() {
    if condition {           // 第 1 层
        if condition2 {      // 第 2 层
            if condition3 {  // 第 3 层
                statement1()
            }
        }
    }
}

// 4 层（超过限制，必须重构）
func example() {
    if condition {           // 第 1 层
        if condition2 {      // 第 2 层
            if condition3 {  // 第 3 层
                if condition4 { // 第 4 层 - 超过限制！
                    statement1()
                }
            }
        }
    }
}
```

#### ❌ 错误示例：嵌套过深

```go
// 嵌套层级：5 层（超过限制）
func (s *orderService) ProcessOrder(ctx context.Context, orderID uint) error {
    // 第 1 层
    if orderID > 0 {
        order, err := s.orderBiz.GetOrder(ctx, orderID)
        if err == nil {
            // 第 2 层
            if order.Status == "pending" {
                // 第 3 层
                for _, item := range order.Items {
                    // 第 4 层
                    if item.ProductID > 0 {
                        stock, err := s.stockBiz.CheckStock(ctx, item.ProductID)
                        if err == nil {
                            // 第 5 层 - 超过限制！
                            if stock >= item.Quantity {
                                // ... 处理逻辑
                            }
                        }
                    }
                }
            }
        }
    }
    return nil
}
```

#### ✅ 正确示例：减少嵌套

```go
// 方法 1：提前返回（Guard Clauses）
func (s *orderService) ProcessOrder(ctx context.Context, orderID uint) error {
    // 提前返回，减少嵌套
    if orderID == 0 {
        return errorpkg.ErrorBadRequest("invalid order_id")
    }

    order, err := s.orderBiz.GetOrder(ctx, orderID)
    if err != nil {
        return err
    }

    if order.Status != "pending" {
        return nil
    }

    return s.processOrderItems(ctx, order.Items)
}

// 辅助函数：处理订单商品（独立函数，不会增加主函数嵌套）
func (s *orderService) processOrderItems(ctx context.Context, items []*OrderItem) error {
    for _, item := range items {
        if err := s.processOrderItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}

// 辅助函数：处理单个商品
func (s *orderService) processOrderItem(ctx context.Context, item *OrderItem) error {
    if item.ProductID == 0 {
        return errorpkg.ErrorBadRequest("invalid product_id")
    }

    stock, err := s.stockBiz.CheckStock(ctx, item.ProductID)
    if err != nil {
        return err
    }

    if stock < item.Quantity {
        return errorpkg.ErrorConflict("insufficient stock")
    }

    return s.stockBiz.DeductStock(ctx, item.ProductID, item.Quantity)
}

// 方法 2：使用 continue 跳过无效数据
func (s *orderService) ProcessOrderV2(ctx context.Context, orderID uint) error {
    if orderID == 0 {
        return errorpkg.ErrorBadRequest("invalid order_id")
    }

    order, err := s.orderBiz.GetOrder(ctx, orderID)
    if err != nil {
        return err
    }

    if order.Status != "pending" {
        return nil
    }

    // 使用 continue 减少嵌套
    for _, item := range order.Items {
        if item.ProductID == 0 {
            continue  // 跳过无效商品
        }

        stock, err := s.stockBiz.CheckStock(ctx, item.ProductID)
        if err != nil {
            return err
        }

        if stock >= item.Quantity {
            s.stockBiz.DeductStock(ctx, item.ProductID, item.Quantity)
        }
    }

    return nil
}
```

#### 减少嵌套的技巧

| 技巧 | 说明 |
|------|------|
| **提前返回** | 先处理错误情况，提前返回 |
| **continue/break** | 在循环中使用 continue 跳过无效数据 |
| **提取函数** | 将深层嵌套的代码提取为独立函数 |
| **卫语句** | 使用 `if !condition { return }` 代替 `if condition { ... }` |

```go
// ❌ 错误：深层嵌套
func process(data []string) {
    for _, s := range data {
        if s != "" {
            if len(s) > 10 {
                if strings.HasPrefix(s, "prefix") {
                    // ... 处理逻辑
                }
            }
        }
    }
}

// ✅ 正确：提前返回 + 提取函数
func process(data []string) {
    for _, s := range data {
        if s == "" {
            continue
        }
        if len(s) <= 10 {
            continue
        }
        if !strings.HasPrefix(s, "prefix") {
            continue
        }
        // ... 处理逻辑
    }
}
```

#### 嵌套层级检查清单

- [ ] 主函数嵌套不超过 3 层
- [ ] 使用提前返回减少嵌套
- [ ] 使用 continue/break 减少循环嵌套
- [ ] 将深层逻辑提取为独立函数
- [ ] 每个函数职责单一，避免复杂嵌套

---

### 函数过长（续）

#### ❌ 错误示例：函数过长

```go
// 这个函数有 200+ 行，违反了长度限制
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // 1. 参数验证 (20 行)
    if req.GetUserId() == 0 {
        return nil, errorpkg.ErrorBadRequest("user_id is required")
    }
    // ... 更多验证

    // 2. 查询用户信息 (15 行)
    user, err := s.userBiz.GetUser(ctx, req.GetUserId())
    // ... 处理逻辑

    // 3. 查询商品信息 (30 行)
    for _, item := range req.GetItems() {
        product, err := s.productBiz.GetProduct(ctx, item.GetProductId())
        // ... 处理逻辑
    }

    // 4. 计算价格 (25 行)
    var totalAmount float64
    for _, item := range req.GetItems() {
        // ... 计算逻辑
    }

    // 5. 检查库存 (20 行)
    for _, item := range req.GetItems() {
        stock, err := s.stockBiz.CheckStock(ctx, item.GetProductId())
        // ... 检查逻辑
    }

    // 6. 应用优惠券 (15 行)
    if req.GetCouponCode() != "" {
        coupon, err := s.couponBiz.GetCoupon(ctx, req.GetCouponCode())
        // ... 处理逻辑
    }

    // 7. 创建订单 (20 行)
    order := &po.Order{
        // ... 订单信息
    }

    // 8. 扣减库存 (15 行)
    for _, item := range req.GetItems() {
        err := s.stockBiz.DeductStock(ctx, item.GetProductId(), item.GetQuantity())
        // ... 处理逻辑
    }

    // 9. 发送通知 (10 行)
    // ... 通知逻辑

    // 10. 返回结果 (10 行)
    return &pb.CreateOrderResp{OrderId: order.ID}, nil
}
```

#### ✅ 正确示例：拆分为多个函数

```go
// 主函数：简洁清晰，只负责协调调用
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // 1. 参数验证
    if err := s.validateCreateOrderReq(req); err != nil {
        return nil, err
    }

    // 2. 查询用户
    user, err := s.userBiz.GetUser(ctx, req.GetUserId())
    if err != nil {
        return nil, err
    }

    // 3. 查询商品并计算
    items, err := s.validateAndCalculateItems(ctx, req.GetItems())
    if err != nil {
        return nil, err
    }

    // 4. 应用优惠券
    totalAmount := s.calculateTotalAmount(items)
    if req.GetCouponCode() != "" {
        totalAmount, err = s.applyCoupon(ctx, req.GetCouponCode(), totalAmount)
        if err != nil {
            return nil, err
        }
    }

    // 5. 创建订单
    order, err := s.createOrderWithItems(ctx, user, items, totalAmount)
    if err != nil {
        return nil, err
    }

    // 6. 异步处理后续操作
    go s.afterOrderCreated(context.Background(), order, items)

    return &pb.CreateOrderResp{OrderId: order.ID}, nil
}

// 辅助函数 1：参数验证
func (s *orderService) validateCreateOrderReq(req *pb.CreateOrderReq) error {
    if req.GetUserId() == 0 {
        return errorpkg.ErrorBadRequest("user_id is required")
    }
    if len(req.GetItems()) == 0 {
        return errorpkg.ErrorBadRequest("items is required")
    }
    return nil
}

// 辅助函数 2：验证商品并计算
func (s *orderService) validateAndCalculateItems(ctx context.Context, reqItems []*pb.OrderItem) ([]*OrderItem, error) {
    items := make([]*OrderItem, 0, len(reqItems))

    for _, reqItem := range reqItems {
        product, err := s.productBiz.GetProduct(ctx, reqItem.GetProductId())
        if err != nil {
            return nil, err
        }

        stock, err := s.stockBiz.CheckStock(ctx, reqItem.GetProductId())
        if stock < reqItem.GetQuantity() {
            return nil, errorpkg.ErrorConflict("insufficient stock")
        }

        items = append(items, &OrderItem{
            ProductID: reqItem.GetProductId(),
            Quantity:  reqItem.GetQuantity(),
            Price:     product.Price,
        })
    }

    return items, nil
}

// 辅助函数 3：计算总金额
func (s *orderService) calculateTotalAmount(items []*OrderItem) float64 {
    var total float64
    for _, item := range items {
        total += float64(item.Quantity) * item.Price
    }
    return total
}

// 辅助函数 4：应用优惠券
func (s *orderService) applyCoupon(ctx context.Context, couponCode string, amount float64) (float64, error) {
    coupon, err := s.couponBiz.GetCoupon(ctx, couponCode)
    if err != nil {
        return 0, err
    }
    return amount * (1 - coupon.Discount), nil
}

// 辅助函数 5：创建订单
func (s *orderService) createOrderWithItems(ctx context.Context, user *User, items []*OrderItem, amount float64) (*Order, error) {
    order := &Order{
        UserID:    user.ID,
        Items:     items,
        Amount:    amount,
        Status:    OrderStatusPending,
    }
    return s.orderBiz.CreateOrder(ctx, order)
}

// 辅助函数 6：订单创建后的异步处理
func (s *orderService) afterOrderCreated(ctx context.Context, order *Order, items []*OrderItem) {
    threadpkg.GoSafe(func() {
        // 扣减库存
        for _, item := range items {
            s.stockBiz.DeductStock(ctx, item.ProductID, item.Quantity)
        }
        // 发送通知
        s.notificationBiz.SendOrderCreated(ctx, order)
    })
}
```

#### 拆分函数的收益

| 方面 | 函数过长 | 拆分函数 |
|------|----------|----------|
| **可读性** | 难以理解，需要滚屏 | 一目了然，逻辑清晰 |
| **可测试性** | 难以单独测试某个逻辑 | 每个函数可独立测试 |
| **可维护性** | 修改一处影响整体 | 修改局部不影响其他 |
| **可复用性** | 无法复用 | 辅助函数可复用 |
| **代码审查** | 难以发现错误 | 容易定位问题 |

#### 函数拆分原则

1. **单一职责**：每个函数只做一件事
2. **长度限制**：不超过 150 行
3. **参数合理**：参数不超过 5 个
4. **命名清晰**：函数名准确描述其功能
5. **合理抽象**：提取通用逻辑为辅助函数

#### 拆分时机

当函数出现以下情况时，应该拆分：

- [ ] 函数超过 150 行
- [ ] 函数有多层嵌套（超过 3 层）
- [ ] 函数有多个独立的功能块
- [ ] 函数有大量重复的验证逻辑
- [ ] 函数难以理解或测试

### 第三方 API 调用重复

**禁止复制粘贴代码！** 发现重复的第三方 API 调用代码时，必须改写适配为可复用的组件。

#### ❌ 错误示例：复制粘贴实现

```go
// 第一次实现：发送验证码短信
func (s *userService) SendVerificationSMS(ctx context.Context, phone string) error {
    client := &http.Client{Timeout: 30 * time.Second}
    body := map[string]string{
        "phone":   phone,
        "message": "Your code is 123456",
        "apikey":  "abc123",
    }
    jsonData, _ := json.Marshal(body)
    req, _ := http.NewRequestWithContext(ctx, "POST", "https://api.sms.com/send", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("sms send failed")
    }
    return nil
}

// 第二次实现：发送通知短信（复制粘贴后修改）
func (s *userService) SendNotificationSMS(ctx context.Context, phone string) error {
    client := &http.Client{Timeout: 30 * time.Second}  // 重复！
    body := map[string]string{
        "phone":   phone,
        "message": "You have a new notification",  // 修改了消息
        "apikey":  "abc123",  // 重复！
    }
    jsonData, _ := json.Marshal(body)  // 重复！
    req, _ := http.NewRequestWithContext(ctx, "POST", "https://api.sms.com/send", bytes.NewBuffer(jsonData))  // 重复！
    req.Header.Set("Content-Type", "application/json")  // 重复！

    resp, err := client.Do(req)  // 重复！
    if err != nil {
        return err  // 重复！
    }
    defer resp.Body.Close()  // 重复！

    if resp.StatusCode != 200 {  // 重复！
        return fmt.Errorf("sms send failed")  // 重复！
    }
    return nil
}
```

#### ✅ 正确示例：改写适配为可复用组件

```go
// 第一步：定义通用的第三方客户端接口
type SMSClient interface {
    Send(ctx context.Context, phone, message string) error
    SendVerification(ctx context.Context, phone string) error
    SendNotification(ctx context.Context, phone string) error
}

// 第二步：实现客户端，封装所有第三方调用细节
type smsClient struct {
    client    *http.Client
    apiKey    string
    baseURL   string
    timeout   time.Duration
}

func NewSmsClient(config *Config) SMSClient {
    return &smsClient{
        client: &http.Client{
            Timeout: config.Timeout,
            Transport: &http.Transport{
                MaxIdleConns:        100,
                MaxIdleConnsPerHost: 10,
                IdleConnTimeout:     90 * time.Second,
            },
        },
        apiKey:  config.SMSAPIKey,
        baseURL: config.SMSBaseURL,
    }
}

// 通用的发送方法
func (c *smsClient) Send(ctx context.Context, phone, message string) error {
    body := map[string]string{
        "phone":   phone,
        "message": message,
        "apikey":  c.apiKey,
    }
    jsonData, err := json.Marshal(body)
    if err != nil {
        return errorpkg.FormatError(err)
    }

    req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/send", bytes.NewBuffer(jsonData))
    if err != nil {
        return errorpkg.FormatError(err)
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := c.client.Do(req)
    if err != nil {
        return errorpkg.WrapWithMetadata(err, nil)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return c.handleErrorResponse(resp)
    }

    return nil
}

// 适配具体业务场景的便捷方法
func (c *smsClient) SendVerification(ctx context.Context, phone string) error {
    return c.Send(ctx, phone, "Your verification code is: {{code}}")
}

func (c *smsClient) SendNotification(ctx context.Context, phone string) error {
    return c.Send(ctx, phone, "You have a new notification")
}

// 统一的错误处理
func (c *smsClient) handleErrorResponse(resp *http.Response) error {
    switch resp.StatusCode {
    case http.StatusBadRequest:
        return errorpkg.ErrorBadRequest("invalid sms request")
    case http.StatusUnauthorized:
        return errorpkg.ErrorInternal("sms api key invalid")
    case http.StatusTooManyRequests:
        return errorpkg.ErrorConflict("sms rate limit exceeded")
    default:
        return errorpkg.ErrorInternal("sms service unavailable")
    }
}

// 第三步：在 Biz 层使用
type userBiz struct {
    smsClient SMSClient  // 依赖接口，便于测试
}

func (b *userBiz) SendVerificationSMS(ctx context.Context, phone string) error {
    return b.smsClient.SendVerification(ctx, phone)
}

func (b *userBiz) SendNotificationSMS(ctx context.Context, phone string) error {
    return b.smsClient.SendNotification(ctx, phone)
}
```

#### 改写适配的关键点

| 方面 | 复制粘贴 | 改写适配 |
|------|----------|----------|
| **代码量** | 每次复制 30+ 行 | 调用 1 行 |
| **维护成本** | 修改需要改 N 处 | 修改 1 处即可 |
| **配置管理** | 硬编码在各处 | 集中配置 |
| **错误处理** | 各不相同 | 统一转换 |
| **测试难度** | 需要 mock HTTP | 可 mock 接口 |
| **扩展性** | 添加功能需复制 | 添加方法即可 |

#### 发现重复代码时的改写步骤

1. **识别共同点**：找出重复的 HTTP 调用、请求构建、错误处理
2. **定义接口**：根据业务场景定义清晰的接口方法
3. **实现客户端**：封装第三方调用细节，统一配置和错误处理
4. **适配业务**：为具体业务场景提供便捷方法
5. **替换调用**：用新接口替换所有复制粘贴的代码
6. **删除重复**：删除所有重复代码

**审查清单：**
- [ ] 是否有复制粘贴的第三方 API 调用代码
- [ ] 是否将第三方服务封装为独立的 Client/Component
- [ ] HTTP 客户端是否复用（而非每次创建新实例）
- [ ] 请求构建逻辑是否抽取为通用函数
- [ ] 错误处理和重试逻辑是否统一
- [ ] 配置（API Key、Timeout）是否集中管理
- [ ] 第三方错误是否转换为项目统一格式

**审查原则：**
1. **一次且仅一次** - 相同的第三方 API 调用逻辑只能出现一次
2. **封装客户端** - 第三方服务必须封装为独立的 Client/Component
3. **统一配置** - API Key、Timeout 等配置必须集中管理
4. **统一错误处理** - 第三方错误必须转换为项目统一的错误格式
5. **接口优先** - 定义接口而非直接依赖具体实现，便于测试和替换
6. **改写适配** - 发现重复代码必须改写，禁止复制粘贴

## 审查评论模板

### 建议修改

```markdown
### 建议：[标题]

**当前代码：**
\`\`\`go
// 粘贴代码
\`\`\`

**问题：**
说明具体问题

**建议修改：**
\`\`\`go
// 建议的代码
\`\`\`

**原因：**
解释为什么这样改更好
```

### 必须修改

```markdown
### 🔴 必须修复：[标题]

**问题：**
[严重问题说明]

**影响：**
[如果不修复会有什么后果]

**建议：**
[修复方案]
```

### 可以优化

```markdown
### 💡 优化建议：[标题]

**当前实现：**
[描述]

**优化方案：**
[描述]

**收益：**
[性能/可读性/维护性提升]
```

## 审查流程

### 1. 自动检查

```bash
# 运行测试
go test ./...

# 格式检查
gofmt -l .

# 静态分析
go vet ./...
golangci-lint run
```

### 2. 人工审查

按照审查清单逐项检查

### 3. 反馈

- 使用友好的语气
- 提供具体的改进建议
- 解释原因

### 4. 跟进

- 作者修改后及时重新审查
- 修改满意后批准（LGTM）
- 有大问题则请求变更（Request Changes）

## 审查响应

### 对于评论者

- 提出问题后关注作者回复
- 讨论达成共识后更新评论状态

### 对于作者

- 积极回应每条评论
- 不同意时说明理由
- 修改后标记已解决

## 审查工具

### GitHub 功能

- Review Changes - 逐文件审查
- Line Comments - 行内评论
- Suggestions - 代码建议
- Approve/Request Changes - 审查决策

### 本地工具

```bash
# 查看 PR 变更
git diff main...feature-branch

# 查看特定文件
git diff main...feature-branch -- path/to/file
```
