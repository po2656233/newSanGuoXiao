package base

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestCreateTokenHs256(t *testing.T) {
	type args struct {
		claims CustomClaims
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "tokeTest",
			args: args{
				claims: CustomClaims{
					ID:      11,
					PlatId:  1,
					Account: "ddd",
					StandardClaims: &jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTokenHs256(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTokenHs256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTokenHs256() got = %v, want %v", got, tt.want)
			}
		})
	}
}
