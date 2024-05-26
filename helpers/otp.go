package helpers

import (
	"bytes"
	"log"
	"math/rand"
	"path/filepath"
	"text/template"
	"time"

	"gopkg.in/gomail.v2"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

type EmailData struct {
	Name    string
	OTPCode []string
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

func SendOTP(to, name, otpCode string) error {
	templatePath := filepath.Join("helpers", "email_template.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Failed to parse template:", err)
		return err
	}

	otpDigits := make([]string, len(otpCode))
	for i, char := range otpCode {
		otpDigits[i] = string(char)
	}

	data := EmailData{
		Name:    name,
		OTPCode: otpDigits,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Println("Failed to execute template:", err)
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "Tourease <tourease03@gmail.com>")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Kode Verifikasi (OTP)")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "tourease03@gmail.com", "psjyhtudsaxrptvd")

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}
