
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {

	UserID string `json:"userid"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type ApiError struct {
	ErrorCode int `json:"error_code"`
	Message string `json:"message"`
}

func (apiErr *ApiError) send (w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(apiErr); err != nil {
		panic(err)
	}
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

