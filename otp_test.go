package otp

import (
	"hash"
	"reflect"
	"testing"
)

func TestNewOption(t *testing.T) {
	type args struct {
		h func() hash.Hash
		d int
	}
	tests := []struct {
		name string
		args args
		want *Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOption(tt.args.h, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateHOTP(t *testing.T) {
	type args struct {
		secret  string
		counter int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateHOTP(tt.args.secret, tt.args.counter); got != tt.want {
				t.Errorf("GenerateHOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GenerateHOTP(t *testing.T) {
	type fields struct {
		HashAlgo func() hash.Hash
		Digits   int
	}
	type args struct {
		secret  string
		counter int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := &Option{
				HashAlgo: tt.fields.HashAlgo,
				Digits:   tt.fields.Digits,
			}
			if got := option.GenerateHOTP(tt.args.secret, tt.args.counter); got != tt.want {
				t.Errorf("Option.GenerateHOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateTOTP(t *testing.T) {
	type args struct {
		secret   string
		interval int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateTOTP(tt.args.secret, tt.args.interval); got != tt.want {
				t.Errorf("GenerateTOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GenerateTOTP(t *testing.T) {
	type fields struct {
		HashAlgo func() hash.Hash
		Digits   int
	}
	type args struct {
		secret   string
		interval int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := &Option{
				HashAlgo: tt.fields.HashAlgo,
				Digits:   tt.fields.Digits,
			}
			if got := option.GenerateTOTP(tt.args.secret, tt.args.interval); got != tt.want {
				t.Errorf("Option.GenerateTOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
