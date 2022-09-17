package bcrypt

import (
	// "fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// fmt.Print("cek err", err)
	return err == nil
}
