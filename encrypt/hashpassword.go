package encrypt

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//encrypts plain text passwords
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil { 
		log.Println(err)
		return "", errors.New("error occurred during password encryption")
	}
	return string(bytes), nil
}