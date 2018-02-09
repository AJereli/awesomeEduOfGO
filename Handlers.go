package main

import (
	"awesomeProject/Auth"
	"awesomeProject/Trash"

	"encoding/json"
	"fmt"
	"database/sql"
	"io"
	"io/ioutil"
	"net/http"
	_ "reflect"
	_ "log"

	_ "github.com/go-sql-driver/mysql"
	_"github.com/gorilla/mux"

)

func Registration (w http.ResponseWriter, r * http.Request){
	//var regInfo Auth.RegistrationInfo

	params := r.URL.Query()

	if !checkParams(params){
		wrongParamsApiErr.send(w)
		return
	}

	uID, userPass := params["userid"][0], params["password"][0]

	db, err := sql.Open("mysql", DBForGoInfo.GetDataSourceName())
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	var userExists bool
	db.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE userid = '%s')	", uID)).Scan(&userExists)

	if userExists{
		notExistUserName.send(w)
		return
	}

	accessToken := Auth.CreateToken(uID)

	stmt, err := db.Prepare("INSERT users SET userid=?, password=?, access_token=?")
	checkErr(err)

	res, err := stmt.Exec(uID, userPass, accessToken)
	checkErr(err)

	fmt.Println(res)

	jsonToken := Auth.JSONToken{AccessToken: accessToken}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(jsonToken); err != nil {
		panic(err)
	}
}

func QueryTest (w http.ResponseWriter, r * http.Request){
	url := r.URL
	params := url.Query()

	for k,v := range params{
		fmt.Fprintln(w, "Your param (k v): ", k, v)
	}

}


func Login (w http.ResponseWriter, r * http.Request){
	var user User


	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(user.UserID)


	//for _, u := range testUsers{
	//	if u.UserID == user.UserID && u.Password == user.Password{
	//		accessToken := Auth.CreateToken(user.UserID)
	//		loginInfo := LoginInfo{AccessToken:accessToken}
	//
	//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//		w.WriteHeader(http.StatusCreated)
	//		if err := json.NewEncoder(w).Encode(loginInfo); err != nil {
	//			panic(err)
	//		}
	//		return
	//	}
	//}
	//
	//apiErr := ApiError{ErrorCode: 403, Message:"Unautorizated\nWrong login or password :(\n"}
	//apiErr.send(w)
	}


func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}



func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Trash.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(string(body))
	fmt.Println(todo)
	t := Trash.RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}