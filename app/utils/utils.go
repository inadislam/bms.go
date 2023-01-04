package utils

import (
	"log"
	"crypto/rand"
	"math/big"

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

func GenerateOTP() (int64, error) {
	max := big.NewInt(999999)
	numb, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return numb.Int64(), nil
}