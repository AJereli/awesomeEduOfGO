
package main

import (
	_"encoding/json"
	_"net/http"
)

type User struct {

	UserID string `json:"userid"`
	Password string `json:"password"`
	Email string `json:"email"`
}



type LoginInfo struct {
	AccessToken string `json:"access_token"`
}

type Users []User


var testUsers Users

func InitDB(){
	testUsers = append(testUsers, User{UserID: "test",Password: "pass", Email: "email"})
	testUsers = append(testUsers, User{UserID: "Serj Tankian", Password: "pass", Email: "email"})

}

