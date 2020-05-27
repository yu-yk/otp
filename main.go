package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"math"
	"time"
)

func main() {
	secret := "secret"
	fmt.Println(TOTP(secret))
}

// TOTP = HOTP(K, T)
// More specifically, T = (Current Unix time - T0) / X,
// where the default floor function is used in the computation.
// For example, with T0 = 0 and Time Step X = 30, T = 1 if the current
// Unix time is 59 seconds, and T = 2 if the current Unix time is
// 60 seconds.
func TOTP(secret string) string {
	X := 30 // Time Step X
	T0 := 0 // T0
	T := (int(time.Now().Unix()) - T0) / X
	K := secret
	return HOTP(K, T)
}

// HOTP(K,C) = Truncate(HMAC-SHA-1(K,C))
// where Truncate represents the function that can convert an HMAC-SHA-1
// value into an HOTP value.  K and C represent the shared secret and
// counter value; see [RFC4226] for detailed definitions.
func HOTP(K string, C int) string {
	// put counter value into text byte array
	text := make([]byte, 8)
	for i := len(text) - 1; i >= 0; i-- {
		text[i] = byte(C & 0xff)
		C = C >> 8
	}

	// using HMAC-SHA-512
	h := hmac.New(sha512.New, []byte(K))
	h.Write(text)
	hmacHash := h.Sum(nil)

	// 64 bytes hmacHash
	offset := int(hmacHash[len(hmacHash)-1] & 0xf)
	binCode := ((int(hmacHash[offset]) & 0x7f) << 24) |
		((int(hmacHash[offset+1] & 0xff)) << 16) |
		((int(hmacHash[offset+2] & 0xff)) << 8) |
		(int(hmacHash[offset+3]) & 0xff)

	// using 10 digits
	binCode %= int(math.Pow10(10))
	return fmt.Sprintf("%010d", binCode)
}
