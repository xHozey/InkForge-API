package pkg

import (
	"fmt"
	"net/mail"

	"inkforge/database/models"
	"inkforge/setup"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hash), err
}

func ValidateSignup(credential Credential) bool {
	var user models.User
	setup.DB.Where("email = ? OR username = ?", credential.Email, credential.Username).First(&user)
	fmt.Println(user)
	if _, err := mail.ParseAddress(credential.Email); err != nil {
		return false
	}
	if !validateLength(credential.Username) || !validateLength(credential.FirstName) || !validateLength(credential.LastName) {
		return false
	}
	return true
}
