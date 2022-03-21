package utils

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func CompareHashedPassword(hashedPasword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}