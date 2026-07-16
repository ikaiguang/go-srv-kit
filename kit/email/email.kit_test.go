package emailpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fakeSender() *Sender {
	return &Sender{
		Issuer:   "UUFFF",
		Host:     "smtp.qiye.aliyun.com",
		Port:     465,
		Username: "",
		Password: "",
	}
}

func fakeCodeMessage() *CodeMessage {
	return &CodeMessage{
		Message: &Message{
			From:    "noreply@uufff.com",
			To:      []string{"example1@gmail.com"},
			Cc:      "example2@gmail.com",
			Subject: "UUUFFF.COM",
			Body:    "2233",
		},
		Code: "6789",
	}
}

// go test -v -count 1 ./kit/email -run TestSendCode
func TestSendCode(t *testing.T) {
	type args struct {
		sender *Sender
		msg    *CodeMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#missing SMTP host",
			args: args{
				sender: &Sender{},
				msg:    fakeCodeMessage(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendCode(tt.args.sender, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailNilSafety(t *testing.T) {
	message := fakeCodeMessage().Message
	sender := fakeSender()

	assert.NotPanics(t, func() {
		require.Error(t, Send(nil, message))
		require.Error(t, Send(sender, nil))
		require.Error(t, SendCode(sender, nil))
	})

	client, err := NewClient(*sender)
	require.NoError(t, err)
	assert.NotPanics(t, func() {
		require.Error(t, client.Send(nil))
		require.Error(t, client.SendCode(nil))
	})
}
