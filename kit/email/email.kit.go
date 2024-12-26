package emailpkg

import (
	_ "embed"
	"fmt"
	bufferpkg "github.com/ikaiguang/go-srv-kit/kit/buffer"
	"gopkg.in/gomail.v2"
	"html/template"
)

var (
	//go:embed email_code.html
	emailCodeHTML     string
	emailCodeTemplate *template.Template

	DefaultSubject = "New-Email"
	DefaultIssuer  = "issuer"
)

func init() {
	var err error
	emailCodeTemplate, err = template.New("email_code").Parse(emailCodeHTML)
	if err != nil {
		panic(err)
	}
}

type Sender struct {
	Issuer   string // 发行人
	Host     string // 主机
	Port     int    // 端口
	Username string // 用户名
	Password string // 密码
}

func (s *Sender) Validate() error {
	if s.Issuer == "" {
		s.Issuer = DefaultIssuer
	}
	if s.Host == "" {
		return fmt.Errorf("host is required")
	}
	if s.Port < 1 {
		return fmt.Errorf("port is required")
	}
	return nil
}

type Message struct {
	From    string   // 发件人
	To      []string // 收件人
	Cc      string   // 抄送
	Subject string   // 主题
	Body    string   // 内容
}

func (msg *Message) Validate() error {
	if msg.From == "" {
		return fmt.Errorf("from is required")
	}
	if len(msg.To) == 0 {
		return fmt.Errorf("to is required")
	}
	if msg.Subject == "" {
		msg.Subject = DefaultSubject
	}
	if msg.Body == "" {
		return fmt.Errorf("body is required")
	}
	return nil
}

func (msg *Message) EmailMessage() *gomail.Message {
	content := gomail.NewMessage()
	content.SetHeader("From", msg.From)
	content.SetHeader("To", msg.To...)
	if msg.Cc != "" {
		content.SetAddressHeader("Cc", msg.Cc, msg.Cc)
	}
	content.SetHeader("Subject", msg.Subject)
	content.SetBody("text/html", msg.Body)
	return content
}

func Send(sender *Sender, msg *Message) error {
	if err := sender.Validate(); err != nil {
		return err
	}
	if err := msg.Validate(); err != nil {
		return err
	}

	d := gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password)
	if err := d.DialAndSend(msg.EmailMessage()); err != nil {
		return err
	}
	return nil
}

type CodeMessage struct {
	*Message
	Code string
}

func SendCode(sender *Sender, msg *CodeMessage) error {
	if msg.Code == "" {
		return fmt.Errorf("code is required")
	}

	var buf = bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buf)

	err := emailCodeTemplate.Execute(buf, struct {
		VerificationCode string
		Issuer           string
	}{
		VerificationCode: msg.Code,
		Issuer:           sender.Issuer,
	})
	if err != nil {
		return err
	}
	msg.Body = buf.String()
	return Send(sender, msg.Message)
}
