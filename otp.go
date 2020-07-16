package otp

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"math"
	"time"
)

// Option represent otp option
type Option struct {
	HashAlgo func() hash.Hash
	Digits   int
}

// DefaultOption represent the default otp option.
var DefaultOption = &Option{
	HashAlgo: sha256.New,
	Digits:   6,
}

// NewOption return a custom otp option
func NewOption(h func() hash.Hash, d int) *Option {
	return &Option{h, d}
}

// GenerateHOTP using default option
func GenerateHOTP(secret string, counter int64) string {
	return DefaultOption.GenerateHOTP(secret, counter)
}

// GenerateTOTP using default option
func GenerateTOTP(secret string, interval int64) string {
	return DefaultOption.GenerateTOTP(secret, interval)
}

// GenerateHOTP with custom option
func (option *Option) GenerateHOTP(secret string, counter int64) string {
	// http://tools.ietf.org/html/rfc4226#section-5.4

	// 8-byte counter value, the moving factor.
	// put counter value into text byte array
	byteSecret := make([]byte, 8)
	for i := len(byteSecret) - 1; i >= 0; i-- {
		byteSecret[i] = byte(counter & 0xff)
		counter = counter >> 8
	}

	// Step 1: Generate an HMAC-SHA-1 value Let HS = HMAC-SHA-1(K,C)
	h := hmac.New(option.HashAlgo, []byte(secret))
	h.Write(byteSecret)
	hmacHash := h.Sum(nil)

	// Step 2: Generate a 4-byte string (Dynamic Truncation)
	// Let Sbits = DT(HS)   //  DT, defined below,
	//                      //  returns a 31-bit string
	offset := int(hmacHash[len(hmacHash)-1] & 0xf)
	binCode := ((int(hmacHash[offset]) & 0x7f) << 24) |
		((int(hmacHash[offset+1] & 0xff)) << 16) |
		((int(hmacHash[offset+2] & 0xff)) << 8) |
		(int(hmacHash[offset+3]) & 0xff)

	// Step 3: Compute an HOTP value
	// Let Snum  = StToNum(Sbits)   // Convert S to a number in
	// 								 0...2^{31}-1
	// Return D = Snum mod 10^Digit //  D is a number in the range
	// 								 0...10^{Digit}-1
	binCode %= int(math.Pow10(option.Digits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", option.Digits), binCode)
}

// GenerateTOTP with custom option
func (option *Option) GenerateTOTP(secret string, interval int64) string {
	if interval == 0 {
		interval = 30
	}

	T := time.Now().Unix() / interval
	return option.GenerateHOTP(secret, T)
}
