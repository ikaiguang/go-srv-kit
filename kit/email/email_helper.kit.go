package emailpkg

import (
	"fmt"
	bufferpkg "github.com/ikaiguang/go-srv-kit/kit/buffer"
	"gopkg.in/gomail.v2"
)

type Client interface {
	Send(message *Message) error
	SendCode(message *CodeMessage) error
	Close() error
}

type client struct {
	conn gomail.SendCloser
	conf *Sender
}

func NewClient(sender Sender) (Client, error) {
	if err := sender.Validate(); err != nil {
		return nil, err
	}
	d := gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password)
	cc, err := d.Dial()
	if err != nil {
		return nil, err
	}
	return &client{
		conn: cc,
		conf: &sender,
	}, nil
}

func (s *client) Send(message *Message) error {
	if err := message.Validate(); err != nil {
		return err
	}
	return gomail.Send(s.conn, message.EmailMessage())
}

func (s *client) SendCode(message *CodeMessage) error {
	if message.Code == "" {
		return fmt.Errorf("code is required")
	}

	var buf = bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buf)

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

func (s *client) Close() error {
	return s.conn.Close()
}
