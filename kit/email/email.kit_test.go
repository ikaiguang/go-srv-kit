package emailpkg

import "testing"

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

// go test -v -count=1 ./kit/email -test.run=TestLocalIP
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
			name: "#TestSendCode",
			args: args{
				sender: fakeSender(),
				msg:    fakeCodeMessage(),
			},
			wantErr: false,
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
