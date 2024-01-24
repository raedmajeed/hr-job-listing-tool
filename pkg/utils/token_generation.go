package utils

import (
	"github.com/raedmajeed/hr-job-tool/config"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateToken(email, role string, cfg *config.ConfigParams) (string, error) {
	expireTime := time.Now().Add(time.Minute * 20).Unix()
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg.SECRETKEY))
	if err != nil {
		log.Printf("unable to generate jwt token for user %v, err: %v", email, err.Error())
		return "", err
	}

	return signedToken, nil
}
