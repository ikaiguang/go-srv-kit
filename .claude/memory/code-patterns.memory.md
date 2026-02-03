# 代码模式记忆

## Service 层模式

### 标准 API 实现
```go
func (s *xxxService) XxxMethod(ctx context.Context, req *pb.XxxReq) (*pb.XxxResp, error) {
    // 1. 参数验证
    if req.GetField() == "" {
        return nil, errorpkg.ErrorBadRequest("field is required")
    }

    // 2. DTO → BO
    param := dto.ToBoXxxParam(req)

    // 3. 调用业务逻辑
    result, err := s.xxxBiz.XxxMethod(ctx, param)
    if err != nil {
        logpkg.WithContext(ctx).Errorw("xxx failed", "error", err)
        return nil, err
    }

    // 4. BO → Proto
    return dto.ToProtoXxxResp(result), nil
}
```

## Business 层模式

### 带验证的业务逻辑
```go
func (b *xxxBiz) XxxMethod(ctx context.Context, param *bo.XxxParam) (*bo.XxxResult, error) {
    // 1. 业务验证
    exists, err := b.xxxRepo.CheckExists(ctx, param.Key)
    if err != nil {
        return nil, errorpkg.FormatError(err)
    }
    if exists {
        return nil, customErrors.AlreadyExists()
    }

    // 2. 调用 Data 层
    result, err := b.xxxRepo.XxxMethod(ctx, param)
    if err != nil {
        return nil, errorpkg.WrapWithMetadata(err, nil)
    }

    return result, nil
}
```

## Data 层模式

### 标准查询
```go
func (d *xxxData) Get(ctx context.Context, id uint) (*bo.Result, error) {
    var po po.Xxx
    err := d.db.WithContext(ctx).First(&po, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errorpkg.ErrorNotFound("not found")
        }
        return nil, errorpkg.FormatError(err)
    }
    return toBoResult(&po), nil
}
```

### 事务操作
```go
err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    profile.UserID = user.ID
    return tx.Create(&profile).Error
})
```

## 错误处理模式

### Data 层错误转换
```go
if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errorpkg.ErrorNotFound("xxx not found")
    }
    if errors.Is(err, gorm.ErrDuplicatedKey) {
        return nil, customErrors.AlreadyExists()
    }
    return nil, errorpkg.FormatError(err)
}
```

### Service 层错误包装
```go
if err != nil {
    logpkg.WithContext(ctx).Errorw("operation failed", "error", err)
    return nil, errorpkg.WrapWithMetadata(err, metadata)
}
```

## 第三方 API 调用模式

### 封装第三方客户端
```go
// ✅ 正确：封装第三方 API 客户端
type smsClient struct {
    client    *http.Client
    apiKey    string
    baseURL   string
    timeout   time.Duration
}

func NewSmsClient(config *Config) *smsClient {
    return &smsClient{
        client:  &http.Client{Timeout: config.Timeout},
        apiKey:  config.APIKey,
        baseURL: config.BaseURL,
    }
}

func (c *smsClient) Send(ctx context.Context, phone, message string) error {
    // 统一的请求构建
    req, err := c.buildRequest(ctx, phone, message)
    if err != nil {
        return errorpkg.FormatError(err)
    }

    // 统一的 HTTP 调用
    resp, err := c.client.Do(req)
    if err != nil {
        return errorpkg.WrapWithMetadata(err, nil)
    }
    defer resp.Body.Close()

    // 统一的错误处理
    if resp.StatusCode != http.StatusOK {
        return errorpkg.ErrorInternal("sms send failed")
    }

    return nil
}

// 在 Biz 层使用
func (b *userBiz) SendSMS(ctx context.Context, phone string) error {
    return b.smsClient.Send(ctx, phone, "verification code")
}
```

### 复用 HTTP 客户端
```go
// ❌ 错误：每次调用都创建新的 HTTP 客户端
func callAPI1() error {
    client := &http.Client{Timeout: 30 * time.Second}
    // ...
}

func callAPI2() error {
    client := &http.Client{Timeout: 30 * time.Second}
    // ...
}

// ✅ 正确：复用 HTTP 客户端
type apiClient struct {
    client *http.Client
}

func NewAPIClient() *apiClient {
    return &apiClient{
        client: &http.Client{
            Timeout: 30 * time.Second,
            Transport: &http.Transport{
                MaxIdleConns:        100,
                MaxIdleConnsPerHost: 10,
            },
        },
    }
}
```

### 统一错误转换
```go
// 将第三方错误转换为项目统一格式
func (c *smsClient) handleError(resp *http.Response) error {
    if resp.StatusCode == http.StatusOK {
        return nil
    }

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
```

### 改写适配模式：从复制粘贴到可复用组件

#### 场景：支付接口

**❌ 复制粘贴方式**
```go
// 支付宝支付
func (s *orderService) Alipay(ctx context.Context, order *Order) error {
    client := &http.Client{Timeout: 30 * time.Second}
    body := map[string]string{
        "app_id":  "alipay_app_id",
        "method":  "alipay.trade.create",
        "amount":  order.Amount,
        "subject": order.Subject,
    }
    // ... 30+ 行重复代码
}

// 微信支付（复制粘贴修改）
func (s *orderService) WechatPay(ctx context.Context, order *Order) error {
    client := &http.Client{Timeout: 30 * time.Second}
    body := map[string]string{
        "app_id":  "wechat_app_id",  // 改了
        "method":  "pay.unifiedorder",  // 改了
        "amount":  order.Amount,
        "subject": order.Subject,
    }
    // ... 30+ 行几乎相同的代码
}
```

**✅ 改写适配方式**
```go
// 第一步：定义统一的支付接口
type PaymentProvider interface {
    CreatePayment(ctx context.Context, order *Order) (*PaymentResult, error)
    QueryPayment(ctx context.Context, paymentID string) (*PaymentStatus, error)
    RefundPayment(ctx context.Context, paymentID string, amount float64) error
}

// 第二步：实现具体的支付提供商
type alipayProvider struct {
    client  *http.Client
    appID   string
    secret  string
    baseURL string
}

func NewAlipayProvider(config *Config) PaymentProvider {
    return &alipayProvider{
        client: &http.Client{Timeout: 30 * time.Second},
        appID:  config.AlipayAppID,
        secret: config.AlipaySecret,
        baseURL: "https://openapi.alipay.com",
    }
}

func (p *alipayProvider) CreatePayment(ctx context.Context, order *Order) (*PaymentResult, error) {
    // 支付宝特定的实现
    body := p.buildAlipayRequest(order)
    resp, err := p.doRequest(ctx, "/gateway.do", body)
    return p.parseAlipayResponse(resp)
}

type wechatPayProvider struct {
    client  *http.Client
    appID   string
    mchID   string
    key     string
    baseURL string
}

func NewWechatPayProvider(config *Config) PaymentProvider {
    return &wechatPayProvider{
        client: &http.Client{Timeout: 30 * time.Second},
        appID:  config.WechatAppID,
        mchID:  config.WechatMchID,
        key:    config.WechatKey,
        baseURL: "https://api.mch.weixin.qq.com",
    }
}

func (p *wechatPayProvider) CreatePayment(ctx context.Context, order *Order) (*PaymentResult, error) {
    // 微信特定的实现
    body := p.buildWechatRequest(order)
    resp, err := p.doRequest(ctx, "/pay/unifiedorder", body)
    return p.parseWechatResponse(resp)
}

// 第三步：在业务层使用接口
type orderBiz struct {
    alipay   PaymentProvider  // 依赖接口
    wechat   PaymentProvider  // 依赖接口
}

func (b *orderBiz) CreatePayment(ctx context.Context, provider string, order *Order) (*PaymentResult, error) {
    switch provider {
    case "alipay":
        return b.alipay.CreatePayment(ctx, order)
    case "wechat":
        return b.wechat.CreatePayment(ctx, order)
    default:
        return nil, errorpkg.ErrorBadRequest("unsupported payment provider")
    }
}
```

#### 改写适配的收益

| 方面 | 复制粘贴 | 改写适配 |
|------|----------|----------|
| **代码行数** | N × 30 = 90+ 行 | 接口 3 行 + 各实现 30 行 + 调用 1 行 |
| **添加新支付** | 复制 30 行，修改 5 处 | 实现接口（3 个方法） |
| **修改超时** | 修改 N 个地方 | 修改构造函数 1 处 |
| **单元测试** | 需要 mock HTTP | 可 mock 接口 |
| **问题排查** | 逐个检查 | 统一日志和错误处理 |

#### 改写适配检查清单

当发现以下情况时，必须改写适配：

- [ ] 相同的 HTTP 客户端初始化代码出现 2 次以上
- [ ] 相同的请求构建逻辑（Marshal、Header 设置）重复
- [ ] 相同的错误处理逻辑重复
- [ ] 只有少量参数不同的 API 调用（如 phone/message）
- [ ] 同一第三方服务的不同接口调用

#### 改写适配模板

```go
// 1. 定义接口
type XXXClient interface {
    Method1(ctx, params) (result, error)
    Method2(ctx, params) (result, error)
}

// 2. 实现客户端
type xxxClient struct {
    client  *http.Client
    config  *XXXConfig
}

func NewXXXClient(config *Config) XXXClient {
    return &xxxClient{
        client: &http.Client{Timeout: config.Timeout},
        config: config,
    }
}

// 3. 实现通用方法
func (c *xxxClient) doRequest(ctx, path, body) (*Response, error) {
    // 统一的请求逻辑
}

// 4. 实现接口方法
func (c *xxxClient) Method1(...) (...) {
    // 调用通用方法
}
```

## 函数拆分模式

### 识别需要拆分的函数

**函数过长时需要拆分（超过 150 行）**

常见需要拆分的场景：
- 包含多个独立的功能块
- 有大量重复的验证逻辑
- 嵌套层级超过 3 层
- 难以理解或测试

### 拆分模式：提取验证函数

```go
// ❌ 拆分前：验证逻辑混杂在主函数中
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // 验证用户
    if req.GetUserId() == 0 {
        return nil, errorpkg.ErrorBadRequest("user_id is required")
    }
    // 验证商品
    if len(req.GetItems()) == 0 {
        return nil, errorpkg.ErrorBadRequest("items is required")
    }
    for _, item := range req.GetItems() {
        if item.GetProductId() == 0 {
            return nil, errorpkg.ErrorBadRequest("product_id is required")
        }
        if item.GetQuantity() <= 0 {
            return nil, errorpkg.ErrorBadRequest("quantity must be positive")
        }
    }
    // 验证优惠券
    if req.GetCouponCode() != "" {
        if len(req.GetCouponCode()) < 6 {
            return nil, errorpkg.ErrorBadRequest("invalid coupon code")
        }
    }
    // ... 其他业务逻辑
}

// ✅ 拆分后：提取验证函数
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    if err := s.validateCreateOrderReq(req); err != nil {
        return nil, err
    }
    // ... 其他业务逻辑
}

func (s *orderService) validateCreateOrderReq(req *pb.CreateOrderReq) error {
    if req.GetUserId() == 0 {
        return errorpkg.ErrorBadRequest("user_id is required")
    }
    if err := s.validateOrderItems(req.GetItems()); err != nil {
        return err
    }
    return s.validateCouponCode(req.GetCouponCode())
}

func (s *orderService) validateOrderItems(items []*pb.OrderItem) error {
    if len(items) == 0 {
        return errorpkg.ErrorBadRequest("items is required")
    }
    for _, item := range items {
        if item.GetProductId() == 0 {
            return errorpkg.ErrorBadRequest("product_id is required")
        }
        if item.GetQuantity() <= 0 {
            return errorpkg.ErrorBadRequest("quantity must be positive")
        }
    }
    return nil
}

func (s *orderService) validateCouponCode(code string) error {
    if code != "" && len(code) < 6 {
        return errorpkg.ErrorBadRequest("invalid coupon code")
    }
    return nil
}
```

### 拆分模式：提取计算逻辑

```go
// ❌ 拆分前：计算逻辑混杂
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // ... 验证逻辑

    // 计算价格
    var totalAmount float64
    var discount float64
    for _, item := range req.GetItems() {
        product, _ := s.productBiz.GetProduct(ctx, item.GetProductId())
        subtotal := float64(item.GetQuantity()) * product.Price
        totalAmount += subtotal
    }

    // 应用优惠券
    if req.GetCouponCode() != "" {
        coupon, _ := s.couponBiz.GetCoupon(ctx, req.GetCouponCode())
        discount = totalAmount * coupon.Discount
    }

    // 应用会员折扣
    user, _ := s.userBiz.GetUser(ctx, req.GetUserId())
    if user.VipLevel >= 3 {
        discount += totalAmount * 0.1
    }

    finalAmount := totalAmount - discount

    // ... 创建订单
}

// ✅ 拆分后：提取计算函数
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // ... 验证逻辑

    // 计算价格
    finalAmount, err := s.calculateOrderAmount(ctx, req)
    if err != nil {
        return nil, err
    }

    // ... 创建订单
}

func (s *orderService) calculateOrderAmount(ctx context.Context, req *pb.CreateOrderReq) (float64, error) {
    // 1. 计算商品总价
    totalAmount, err := s.calculateItemsTotal(ctx, req.GetItems())
    if err != nil {
        return 0, err
    }

    // 2. 计算优惠
    discount, err := s.calculateDiscount(ctx, req.GetCouponCode(), req.GetUserId(), totalAmount)
    if err != nil {
        return 0, err
    }

    return totalAmount - discount, nil
}

func (s *orderService) calculateItemsTotal(ctx context.Context, items []*pb.OrderItem) (float64, error) {
    var total float64
    for _, item := range items {
        product, err := s.productBiz.GetProduct(ctx, item.GetProductId())
        if err != nil {
            return 0, err
        }
        total += float64(item.GetQuantity()) * product.Price
    }
    return total, nil
}

func (s *orderService) calculateDiscount(ctx context.Context, couponCode string, userID uint, amount float64) (float64, error) {
    var discount float64

    // 优惠券折扣
    if couponCode != "" {
        coupon, err := s.couponBiz.GetCoupon(ctx, couponCode)
        if err != nil {
            return 0, err
        }
        discount += amount * coupon.Discount
    }

    // 会员折扣
    user, err := s.userBiz.GetUser(ctx, userID)
    if err != nil {
        return 0, err
    }
    if user.VipLevel >= 3 {
        discount += amount * 0.1
    }

    return discount, nil
}
```

### 拆分模式：提取异步处理

```go
// ❌ 拆分前：同步等待所有后续操作
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // ... 创建订单

    // 同步扣减库存
    for _, item := range items {
        if err := s.stockBiz.DeductStock(ctx, item.ProductID, item.Quantity); err != nil {
            // 回滚...
            return nil, err
        }
    }

    // 同步发送通知
    if err := s.notificationBiz.SendOrderCreated(ctx, order); err != nil {
        // 记录错误但不影响主流程
    }

    // 同步更新统计
    s.statsBiz.IncrementOrderCount(ctx, user.ID)

    return &pb.CreateOrderResp{OrderId: order.ID}, nil
}

// ✅ 拆分后：异步处理后续操作
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // ... 创建订单

    // 异步处理后续操作
    go s.afterOrderCreated(context.Background(), order, items, user)

    return &pb.CreateOrderResp{OrderId: order.ID}, nil
}

func (s *orderService) afterOrderCreated(ctx context.Context, order *Order, items []*OrderItem, user *User) {
    threadpkg.GoSafe(func() {
        // 扣减库存
        for _, item := range items {
            if err := s.stockBiz.DeductStock(ctx, item.ProductID, item.Quantity); err != nil {
                logpkg.WithContext(ctx).Errorw("deduct stock failed", "order_id", order.ID, "error", err)
            }
        }

        // 发送通知
        if err := s.notificationBiz.SendOrderCreated(ctx, order); err != nil {
            logpkg.WithContext(ctx).Errorw("send notification failed", "order_id", order.ID, "error", err)
        }

        // 更新统计
        s.statsBiz.IncrementOrderCount(ctx, user.ID)
    })
}
```

### 函数命名规范

提取的辅助函数命名应该体现其功能：

| 函数类型 | 命名模式 | 示例 |
|---------|---------|------|
| 验证函数 | `validate{Xxx}` | `validateCreateOrderReq` |
| 计算函数 | `calculate{Xxx}` | `calculateOrderAmount` |
| 构建函数 | `build{Xxx}` | `buildOrderItems` |
| 转换函数 | `convert{Xxx}To{Yyy}` | `convertUserToPO` |
| 处理函数 | `process{Xxx}` | `processPayment` |
| 异步函数 | `after{Xxx}` | `afterOrderCreated` |
| 辅助函数 | `do{Xxx}` | `doDatabaseOperation` |

### 拆分后的好处

```go
// 主函数清晰简洁
func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
    // 1. 验证
    if err := s.validateCreateOrderReq(req); err != nil {
        return nil, err
    }

    // 2. 查询
    user, err := s.userBiz.GetUser(ctx, req.GetUserId())
    if err != nil {
        return nil, err
    }

    // 3. 计算
    items, err := s.validateAndCalculateItems(ctx, req.GetItems())
    if err != nil {
        return nil, err
    }

    // 4. 创建
    order, err := s.createOrder(ctx, user, items)
    if err != nil {
        return nil, err
    }

    // 5. 后续处理
    go s.afterOrderCreated(context.Background(), order)

    return &pb.CreateOrderResp{OrderId: order.ID}, nil
}
```

**一眼就能看懂整个流程！**

### 拆分检查清单

- [ ] 主函数不超过 150 行
- [ ] 每个辅助函数单一职责
- [ ] 辅助函数可独立测试
- [ ] 辅助函数可复用
- [ ] 函数命名清晰准确
- [ ] 避免过度拆分（不要为了拆而拆）
