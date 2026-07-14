# email

构造邮件消息并通过 SMTP 发送，支持普通邮件和验证码邮件。

## 基础用法

```go
sender := &emailpkg.Sender{Host: "smtp.example.com", Port: 587, Username: "user", Password: "secret"}
msg := &emailpkg.Message{From: "from@example.com", To: []string{"to@example.com"}, Subject: "subject", Body: "body"}
err := emailpkg.Send(sender, msg)
```

## 注意事项

SMTP 账号、密码和授权码必须来自配置或密钥管理；测试中不要发送真实邮件。

## 验证

```bash
go test ./email
```
