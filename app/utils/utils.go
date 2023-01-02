package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HashPassword(password string) (string, error) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	CheckError(err)
	return string(hpass), err
}

func ComparePass(pass, hashedPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	CheckError(err)
	return err
}
