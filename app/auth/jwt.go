package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type AClaim struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type RClaim struct {
	ID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(uid, email string) (access_token string, refresh_token string, err error) {
	// access token
	et := time.Now().Add(15 * time.Minute).Unix()
	claims := &AClaim{
		ID:    uid,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(et)),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err = at.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", "", err
	}

	// refresh token
	rtime := time.Now().Add(30 * 24 * time.Hour).Unix()
	rclaims := &RClaim{
		ID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(rtime)),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)
	refresh_token, err = rt.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", "", err
	}
	return access_token, refresh_token, nil
}
