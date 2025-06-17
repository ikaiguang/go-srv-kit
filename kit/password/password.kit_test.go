package passwordpkg

import (
	"testing"
)

// go test -v -count 1 ./kit/password -run TestCompare
func TestCompare(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#比较密码",
			args: args{
				hashedPassword: "$2a$10$Ptx7B.ZifXx2rDgidTrSXeMa.neg8HRnf8jHsuznKsjKZtGaGOqLW",
				password:       "123456",
			},
			wantErr: false,
		},
		{
			name: "#比较密码",
			args: args{
				hashedPassword: "$2a$10$Ptx7B.ZifXx2rDgidTrSXeMa.neg8HRnf8jHsuznKsjKZtGaGOqLW",
				password:       "1234567",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+":"+tt.args.password, func(t *testing.T) {
			if err := Compare(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// go test -v -count 1 ./kit/password -run TestEncrypt
func TestEncrypt(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "#加密密码",
			args: args{
				password: "123456",
			},
			want:    []byte("$2a$10$Z0/4jarnYqNPGJLAVVLiaOX2KuxmzFpYqNEh37THbEKVvAjtlyPsm"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+":"+tt.args.password, func(t *testing.T) {
			got, err := Encrypt(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Encrypt() got = %s, want %s", got, tt.want)
			//}
			t.Log("==> got :", string(got))
		})
	}
}
