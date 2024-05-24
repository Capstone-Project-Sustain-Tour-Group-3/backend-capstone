package helpers

import (
	"gopkg.in/gomail.v2"
	"log"
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

func SendOTP(to, otpCode string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "leafthe78@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", "Your OTP code is: "+otpCode)

	d := gomail.NewDialer("smtp.gmail.com", 587, "leafthe78@gmail.com", "cvuznzxfmnsgjmwa")

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}
