/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package Auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

)

const (
	AppName = "TestAPI"
	ExpiresTime = 60 * 60 * 24 * 15
)

type TokenInfo struct {
	UserId string `json:"userid"`
	ISS string `json:"iss"`
}



type JSONToken struct{
	AccessToken string `json:"access_token"`
}

var hmacSampleSecret = []byte("asdhguczx1412313214jifh")

func CreateToken (userId string) string{
	claims := jwt.MapClaims{"userId" : userId, "nbf": time.Now().Unix(), "exp":time.Now().Unix() + ExpiresTime,"iss": AppName}
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
		return TokenInfo {ISS: claims["iss"].(string), UserId: claims["userId"].(string)}
	} else {
		panic(err)
	}
}





