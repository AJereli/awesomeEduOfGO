/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main
import (
	"encoding/json"
	"log"
	"net/http"
)

/// API error codes
const (
	wrongRegParmsCode = 100
	wrongParams = "Wrong params"

	userNameNotExistsCode = 101

	loginErrCode = 102



	tokenTimeOutCode = 160


	unprocessableEntityCode = 422
)

var (
	wrongParamsApiErr         = ApiError{wrongRegParmsCode, wrongParams}
	notExistUserNameApiErr    = ApiError{userNameNotExistsCode, "Sorry, this user name is not available"}
	loginApiErr               = ApiError{loginErrCode, "Wrong login or password"}

	tokenTimeOutApiErr	      = ApiError{tokenTimeOutCode, "Access token time out"}

	unprocessableEntityApiErr = ApiError{unprocessableEntityCode, "Unprocessablee entity"}
)

type ApiError struct {
	ErrorCode int `json:"error_code"`
	Message string `json:"message"`
}

func (apiErr *ApiError) send (w http.ResponseWriter){
	log.Println(apiErr.Message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(apiErr); err != nil {
		log.Println(err)
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}