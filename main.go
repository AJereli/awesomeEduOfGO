package main

import (
	"SC"
)

func main(){

	server := SC.MakeServer("server1")
	server.Listen()
}
