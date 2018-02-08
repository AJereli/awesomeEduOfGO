package main

import (
	"fmt"

	"log"
	"net/http"

)




func main() {
	fmt.Println("apiTry.main")

	InitDB()

	tokenString := CreateToken("Some boy")
	fmt.Println("Created token: ", tokenString)
	fmt.Println(ParseToken(tokenString))

	router := InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

