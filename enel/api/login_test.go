package api

import (
	"log"
	"testing"
)

func TestLogin(t *testing.T) {
	type args struct {
		cpf string
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Request witch params",
			args: args{
				cpf: "2615376657125",
				id:  "64357268352",
			},
			want:    "x",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Login(tt.args.cpf, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(got)
			if got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
