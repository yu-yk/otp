//  Duplicate the following results from RFC 4226

//  Appendix D - HOTP Algorithm: Test Values

//    The following test data uses the ASCII string
//    "12345678901234567890" for the secret:

//    Secret = 0x3132333435363738393031323334353637383930

//                      Truncated
//    Count    Hexadecimal    Decimal        HOTP
//    0        4c93cf18       1284755224     755224
//    1        41397eea       1094287082     287082
//    2         82fef30        137359152     359152
//    3        66ef7655       1726969429     969429
//    4        61c5938a       1640338314     338314
//    5        33c083d4        868254676     254676
//    6        7256c032       1918287922     287922
//    7         4e5b397         82162583     162583
//    8        2823443f        673399871     399871
//    9        2679dc69        645520489     520489
package otp

import (
	"testing"
)

var secret = "12345678901234567890"

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
		{
			"Default HOTP 0",
			args{secret, 0},
			"755224",
		},
		{
			"Default HOTP 1",
			args{secret, 1},
			"287082",
		},
		{
			"Default HOTP 2",
			args{secret, 2},
			"359152",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateHOTP(tt.args.secret, tt.args.counter); got != tt.want {
				t.Errorf("GenerateHOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateHOTP(t *testing.T) {
	type args struct {
		code    string
		secret  string
		counter int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Validate HOTP 0",
			args{"755224", secret, 0},
			true,
		},
		{
			"Validate HOTP 1",
			args{"287082", secret, 1},
			true,
		},
		{
			"Validate HOTP 2",
			args{"359152", secret, 2},
			true,
		},
		{
			"Incorrect HOTP",
			args{"123456", secret, 12345},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateHOTP(tt.args.code, tt.args.secret, tt.args.counter); got != tt.want {
				t.Errorf("ValidateHOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
