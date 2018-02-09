package main

import (
	"fmt"

	"log"
	"net/http"

)




func main() {
	fmt.Println("apiTry.main")


	router := InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

