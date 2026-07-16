package emailpkg

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
	"gopkg.in/gomail.v2"
)

type Client interface {
	Send(message *Message) error
	SendCode(message *CodeMessage) error
	Close() error
}

type client struct {
	dialer *gomail.Dialer
	conf   *Sender
}

func NewClient(sender Sender) (Client, error) {
	if err := sender.Validate(); err != nil {
		return nil, err
	}
	d := gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password)
	return &client{
		dialer: d,
		conf:   &sender,
	}, nil
}

func (s *client) Send(message *Message) error {
	if s == nil || s.dialer == nil {
		return fmt.Errorf("email client is nil")
	}
	if err := message.Validate(); err != nil {
		return err
	}
	return s.dialer.DialAndSend(message.EmailMessage())
}

func (s *client) SendCode(message *CodeMessage) error {
	if s == nil || s.conf == nil || message == nil || message.Message == nil {
		return fmt.Errorf("code message or email client is nil")
	}
	if message.Code == "" {
		return fmt.Errorf("code is required")
	}

	var buf = bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	err := emailCodeTemplate.Execute(buf, struct {
		VerificationCode string
		Issuer           string
	}{
		VerificationCode: message.Code,
		Issuer:           s.conf.Issuer,
	})
	if err != nil {
		return err
	}
	message.Body = buf.String()
	return s.Send(message.Message)
}

// Close 当前实现为 no-op。每次调用 Send 时使用 DialAndSend 建立独立连接，无需手动关闭。
func (s *client) Close() error {
	return nil
}

type defaultClient struct {
	sender *Sender
}

func DefaultClient(sender Sender) (Client, error) {
	return &defaultClient{sender: &sender}, nil
}

func (s *defaultClient) Send(message *Message) error {
	if s == nil {
		return fmt.Errorf("email client is nil")
	}
	return Send(s.sender, message)
}

func (s *defaultClient) SendCode(message *CodeMessage) error {
	if s == nil {
		return fmt.Errorf("email client is nil")
	}
	return SendCode(s.sender, message)
}

// Close 当前实现为 no-op。每次调用 Send 时使用 DialAndSend 建立独立连接，无需手动关闭。
func (s *defaultClient) Close() error {
	return nil
}
