package test

import (
	"testing"
)

func TestCompareTwoStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		op1  string
		op2  string
	}{
		{
			name: "Example1 Input(str1=ABC, str2=BC)",
			args: args{
				str1: "ABC",
				str2: "BC",
			},
			op1: "A",
			op2: "",
		},
		{
			name: "Example2 (str1=BC, str2=BANGALORE)",
			args: args{
				str1: "BC",
				str2: "BANGALORE",
			},
			op1: "C",
			op2: "ANGALORE",
		},
		{
			name: "Example3 (str1=TECHCHEFS, str2=VIRTUSA)",
			args: args{
				str1: "TECHCHEFS",
				str2: "VIRTUSA",
			},
			op1: "ECHCHEF",
			op2: "VIRUA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := CompareTwoStrings(tt.args.str1, tt.args.str2)
			if got1 != tt.op1 {
				t.Errorf("CompareTwoStrings() got1 = %v, op1 = %v", got1, tt.op1)
			}
			if got2 != tt.op2 {
				t.Errorf("CompareTwoStrings() got2 = %v, op2 =  %v", got2, tt.op2)
			}
		})
	}
}

func TestGetEmailFromTheResponseOfAPI(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get email from (https://reqres.in/api/users/1)",
			args: args{
				endpoint: "https://reqres.in/api/users/1",
			},
			want:    "george.bluth@reqres.in",
			wantErr: false,
		},
		{
			name: "Get email from (https://reqres.in/api/users/3)",
			args: args{
				endpoint: "https://reqres.in/api/users/3",
			},
			want:    "emma.wong@reqres.in",
			wantErr: false,
		},
		{
			name: "Get email from (https://reqres.in/api/users/10)",
			args: args{
				endpoint: "https://reqres.in/api/users/10",
			},
			want:    "byron.fields@reqres.in",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEmailFromTheResponseOfAPI(tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmailFromTheResponseOfAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEmailFromTheResponseOfAPI() got = %v, want %v", got, tt.want)
			}
		})
	}
}
