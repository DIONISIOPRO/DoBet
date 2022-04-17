package service

import "golang.org/x/crypto/bcrypt"

type PasswordHandler struct{}

func (p PasswordHandler) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
