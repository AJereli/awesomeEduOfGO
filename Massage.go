package main

import (
	"awesomeProject/Auth"
	"database/sql"
	"encoding/json"
	_"fmt"
	_"log"
	"net/http"
	"time"
)

type Massage struct {
	Sender string `json:"sender"`
	Reciver string `json:"reciver"`
	SendDate string `json:"send_date"`
	Body string `json:"body"`
}

func (msg * Massage) PopulateFromRow(rows *sql.Rows, senderId string, reciverId string) {

	err := rows.Scan(&msg.Body, &msg.SendDate)
	checkErr(err)
	msg.Reciver = reciverId
	msg.Sender = senderId
}

func GetMassageFromRows(rows * sql.Rows, senderId string, reciverId string) []Massage{
	var msgs []Massage

	for rows.Next(){
		var msg Massage
		msg.PopulateFromRow(rows, senderId, reciverId)
		msgs = append(msgs, msg)
	}
	return msgs
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
		unprocessableEntityApiErr.send(w)
	}
	token := Auth.ParseToken(msgInfo.Access_token)

	if !token.CheckExpTime() {
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


func GetMassagesFromUser (w http.ResponseWriter, r* http.Request){
	type RequestParams struct {
		AccessToken string `json:"access_token"`
		SenderUser string `json:"sender_id"`
	}

	var reqParams RequestParams

	if UnmarshalRequest(r, &reqParams) != nil{
		unprocessableEntityApiErr.send(w)
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

	rows, err := db.Query("SELECT massage_body, create_date FROM massage WHERE creator_id = ? AND " +
		"id IN (SELECT message_id FROM DBForGO.messange_reciver WHERE reciver_id = ?)", reqParams.SenderUser, token.UserId)
	checkErr(err)




	msgs := GetMassageFromRows(rows, reqParams.SenderUser, token.UserId)
	defer rows.Close()

	SendJson(w, msgs)

}

func GetMassagesToUser (w http.ResponseWriter, r * http.Request){
	type RequestParams struct{
		AccessToken string `json:"access_token"`
		ReciverUser string `json:"reciver_id"`

	}

	var reqParams RequestParams

	if UnmarshalRequest(r, &reqParams) != nil{
		unprocessableEntityApiErr.send(w)
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

	rows, err := db.Query("SELECT massage_body, create_date FROM massage WHERE creator_id = ? AND " +
		"id IN (SELECT message_id FROM DBForGO.messange_reciver WHERE reciver_id = ?)", token.UserId, reqParams.ReciverUser)
	checkErr(err)


	defer rows.Close()

	msgs := GetMassageFromRows(rows, token.UserId, reqParams.ReciverUser)


	SendJson(w, msgs)


}