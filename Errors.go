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
	"net/http"
)

/// API error codes
const (
	wrongRegParmsCode = 100
	wrongParams = "Wrong params"

	userNameNotExistsCode = 101

	loginErrCode = 102
)

var (
	wrongParamsApiErr = ApiError{wrongRegParmsCode, wrongParams}
	notExistUserName = ApiError{userNameNotExistsCode, "Sorry, this user name not available"}
	loginApiErr = ApiError{loginErrCode, "Wrong login or password"}
)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}