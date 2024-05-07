package password

import (
	err_util "github.com/satryanararya/go-chefbot/utils/error"
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtil interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) error
}

type passwordUtil struct {}

func NewPasswordUtil() *passwordUtil {
	return &passwordUtil{}
}

func (p *passwordUtil) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err_util.ErrFailedHashingPassword
	}
	return string(hashedPassword), nil
}

func (p *passwordUtil) ComparePassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err_util.ErrPasswordMismatch
	}
	return nil
}