package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), 10)
	if err != nil {
		return "", nil
	}

	return string(hashed), nil
}

func ComparePassword(hashed, plainText string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText)); err != nil {
		return false, err
	}

	return true, nil
}
