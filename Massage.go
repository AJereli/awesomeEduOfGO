package main

import (
	"awesomeProject/Auth"
	"database/sql"
	"encoding/json"
	_"fmt"
	"log"
	"net/http"
	"time"
)

type Massage struct {
	Sender string `json:"sender"`
	Reciver string `json:"reciver"`
	SendDate string `json:"send_date"`
	Body string `json:"body"`
}

func SendMassage (w http.ResponseWriter, r * http.Request){
	type MassageInfo struct {
		Access_token string `json:"access_token"`
		Massage string `json:"massage"`
		ReciverId string `json:"reciver_id"`
	}

	var msgInfo MassageInfo

	body := ReadRequestBody(r)

	if err := json.Unmarshal(body, &msgInfo); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	token := Auth.ParseToken(msgInfo.Access_token)

	if token.CheckExpTime() {
		tokenTimeOutApiErr.send(w)
		return
	}

	db, err := sql.Open("mysql",  DBForGoInfo.GetDataSourceName())
	checkErr(err)
	err = db.Ping()
	checkErr(err)

	var maxID int
	db.QueryRow("SELECT max(id) FROM massage").Scan(&maxID)

	currMsgId := maxID + 1

	db.Query("INSERT massage SET id = ?, massage_body = ?, create_date = ?, creator_id = ?", currMsgId, msgInfo.Massage, time.Now(), token.UserId)
	db.Query("INSERT messange_reciver SET reciver_id = ?, message_id = ?", msgInfo.ReciverId, currMsgId)

	db.Close()

}


func GetMassagesToUser (w http.ResponseWriter, r * http.Request){

	type RequestParams struct {
		AccessToken string `json:"access_token"`
		ReciverUser string `json:"reciver_id"`

	}

	var reqParams RequestParams

	body := ReadRequestBody(r)

	if err := json.Unmarshal(body, &reqParams); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	token := Auth.ParseToken(reqParams.AccessToken)

	if !token.CheckExpTime() {
		tokenTimeOutApiErr.send(w)
		return
	}

	db, err := sql.Open("mysql", DBForGoInfo.GetDataSourceName())
	defer db.Close()
	checkErr(err)
	err = db.Ping()
	checkErr(err)

	rows, err := db.Query("SELECT massage_body, create_date FROM massage WHERE creator_id = ? AND id IN (SELECT message_id FROM DBForGO.messange_reciver WHERE reciver_id = ?)", token.UserId, reqParams.ReciverUser)
	defer rows.Close()
	var msgs []Massage
	for rows.Next(){
		var msg Massage
		err = rows.Scan(&msg.Body, &msg.SendDate)
		msg.Reciver = reqParams.ReciverUser
		msg.Sender = token.UserId
		msgs = append(msgs, msg)

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(msgs); err != nil {
		panic(err)
	}

}