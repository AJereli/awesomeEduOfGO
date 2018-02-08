package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

)

type TokenInfo struct {
	userId string `json:"userid"`
	iss string `json:"iss"`
}

var hmacSampleSecret = []byte("asdhguczx1412313214jifh")

func CreateToken (userId string) string{
	claims := jwt.MapClaims{"userId" : userId, "nbf": time.Now(), "iss": "TestAPI"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims.Valid()
	toketString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		panic (err)
	}
	return toketString

}

func ParseToken (tokenString string)  TokenInfo{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return TokenInfo {iss: claims["iss"].(string), userId: claims["userId"].(string)}
	} else {
		panic(err)
	}
}