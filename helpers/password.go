package helpers

import "golang.org/x/crypto/bcrypt"

type PasswordHelper interface {
	HashPassword(request string) (string, error)
	VerifyPassword(hashPass, password string) error
}

type passwordHelper struct{}

func NewPasswordHelper() PasswordHelper {
	return &passwordHelper{}
}

func (p *passwordHelper) HashPassword(request string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(request), bcrypt.DefaultCost)
	return string(hashPass), err
}

func (p *passwordHelper) VerifyPassword(hashPass, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	return err
}
