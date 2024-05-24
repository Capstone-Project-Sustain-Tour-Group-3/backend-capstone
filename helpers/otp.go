package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateReferenceId() string {
	const randChars = "abcdefghijklmnopqrstufwxzyABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	ref := make([]byte, 16)
	for i := range ref {
		ref[i] = randChars[rand.Intn(len(randChars))]
	}
	return string(ref)
}

func GenerateOTP() string {
	const otpChars = "1234567890"
	otp := make([]byte, 4)
	for i := range otp {
		otp[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(otp)
}
