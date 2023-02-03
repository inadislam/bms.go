package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"github.com/inadislam/bms-go/app/db"
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

func GenerateAT(uid, email string) (access_token string, err error) {
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
		return "", err
	}
	return
}

func GenerateJWT(uid, email string) (access_token string, refresh_token string, err error) {
	// access token
	access_token, err = GenerateAT(uid, email)
	if err != nil {
		return "", "", err
	}

	// refresh token
	rtime := time.Now().Add(365 * 24 * time.Hour).Unix()
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

func VerifyToken(tr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, err
}

func ExtractTokenAuth(token string) (string, error) {
	tr, err := VerifyToken(token)
	if err != nil {
		return "", err
	}
	if claims, ok := tr.Claims.(jwt.MapClaims); ok && tr.Valid {
		userid, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return "", err
		}
		user, err := db.UserById(userid)
		if err != nil {
			return "", err
		}
		if user.Email == claims["email"] {
			newAT, err := GenerateAT(user.ID.String(), user.Email)
			if err != nil {
				return "", err
			}
			return newAT, nil
		}
	}
	return "", nil
}
