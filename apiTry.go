package main

import (
	"awesomeProject/Auth"
	"fmt"
	"time"

	"log"
	"net/http"

)




func main() {
	fmt.Println("apiTry.main")

	fmt.Println(time.Unix(time.Now().Unix() + Auth.ExpiresTime,0),"\n")

	InitDB()

	tokenString := Auth.CreateToken("Some boy")
	fmt.Println("Created token: ", tokenString)
	fmt.Println(Auth.ParseToken(tokenString))

	router := InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

