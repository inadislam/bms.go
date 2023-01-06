package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JWTClaim struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(uid, email string) (token string, err error) {
	et := time.Now().Add(15 * time.Minute).Unix()
	claims := &JWTClaim{
		ID:    uid,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(et)),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = at.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}
	return
}
