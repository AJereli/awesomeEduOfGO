package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type Server struct {
	pattern string
	messages []*Message
	clients map[int]*Client
	addChan chan *Client
	delChan chan *Client
	sendAllChan chan *Message
	closeChan chan bool
	errorChan chan error
}

func makeServer (pattern string) *Server{
	messages := []*Message{}
	clients := make(map[int]*Client)
	addChan := make(chan *Client)
	delChan := make(chan *Client)
	sendAllChan := make(chan *Message)
	closeChan := make(chan bool)
	errorChan := make(chan error)

	return &Server{
		pattern,
		messages,
		clients,
		addChan,
		delChan,
		sendAllChan,
		closeChan,
		errorChan,
	}
}

func (s *Server) AddClient (cl *Client){
	s.addChan <- cl
}

func (s *Server) DeleteClient (cl *Client){
	s.delChan <- cl
}

func (s *Server) SendMsgToAll (msg *Message){
	s.sendAllChan <- msg
}

func (s *Server) Close (){
	s.closeChan <- true
}

func (s *Server) Error (err error) {
	s.errorChan <- err
}

func (s *Server) SendAllMessages (client *Client){
	for _, msg := range s.messages{
		client.Write(msg)
	}
}

func (s *Server) _SendMsgToAll (msg *Message){
	for _, c := range s.clients{
		c.Write(msg)
	}
}

func (s *Server) Listen (){
	fmt.Println("Start listen")

	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errorChan <- err
			}
		}()

		client := makeClient(ws, s)
		s.AddClient(client)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	fmt.Println("Created handler")

	for {
		select {
		case c := <- s.addChan:
			fmt.Println("New client")
			s.clients[c.id] = c
			s.SendAllMessages(c)
		case c := <- s.delChan:
			fmt.Println("Delete client with id: ", c.id)
			delete(s.clients, c.id)
		case msg := <- s.sendAllChan :
			fmt.Println("Send all: ", msg)
			s.messages = append(s.messages, msg)
			s._SendMsgToAll(msg)
		case err := <-s.errorChan:
			log.Println("Error:", err.Error())
		case <-s.closeChan:
			return
		}
	}
}


