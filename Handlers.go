package main

import (
	"awesomeProject/Auth"
	"awesomeProject/Trash"
	"encoding/json"
	"fmt"
	"database/sql"
	"io"
	"io/ioutil"
	_"log"
	"net/http"
	_ "reflect"
	_ "log"


	_ "github.com/go-sql-driver/mysql"
	_"github.com/gorilla/mux"

)




func QueryTest (w http.ResponseWriter, r * http.Request){
	url := r.URL
	params := url.Query()

	for k,v := range params{
		fmt.Fprintln(w, "Your param (k v): ", k, v)
	}

}


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

	var userNameIsExists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE userid = ?)", uID).Scan(&userNameIsExists)

	if userNameIsExists {
		notExistUserNameApiErr.send(w)
		db.Close()
		return
	}

	accessToken := Auth.CreateToken(uID)

	stmt, err := db.Prepare("INSERT users SET userid=?, password=?, access_token=?")
	checkErr(err)

	res, err := stmt.Exec(uID, userPass, accessToken)
	checkErr(err)

	fmt.Println(res)

	db.Close()

	jsonToken := Auth.JSONToken{AccessToken: accessToken}

	SendJson(w, jsonToken)
}

func Login (w http.ResponseWriter, r * http.Request){
	var user User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitJSONRead))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}


	if err := json.Unmarshal(body, &user); err != nil {
		unprocessableEntityApiErr.send(w)
	}

	var truePassword, trueToken string

	db, err := sql.Open("mysql", DBForGoInfo.GetDataSourceName())
	checkErr(err)

	db.QueryRow("SELECT password FROM users WHERE userid = ?", user.UserID).Scan(&truePassword)
	defer db.Close()

	if user.Password == truePassword{
		trueToken = Auth.CreateToken(user.UserID)
		SendJson(w, Auth.JSONToken{AccessToken: trueToken})
		db.Query("UPDATE users SET access_token = ? WHERE userid = ?", trueToken, user.UserID)
	} else {
		loginApiErr.send(w)
	}

}


func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}



func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Trash.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitJSONRead))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		unprocessableEntityApiErr.send(w)
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