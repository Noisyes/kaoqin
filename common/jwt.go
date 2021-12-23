package common

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("noisyes_will_be_better")

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

func ReleaseJWT(name string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &MyClaims{
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return ss, nil
}

func ParseJWT(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	return token, claims, err
}
