package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"

)

type Client struct {
	id int
	webSock * websocket.Conn
	server *Server
	msgChan chan *Message
	closeChan chan bool
}

var maxID int = 0

func makeClient (webSock * websocket.Conn, server *Server) *Client{
	 if webSock == nil {
	 	panic("WebSocket is nil")
	 }
	 if server == nil {
	 	panic("Server is nil")
	 }
	 maxID++

	 var msgChan = make(chan * Message, maxChanBuffSize)
	 var closeChan = make(chan bool)

	 return &Client{maxID, webSock, server, msgChan, closeChan}
}

func (c *Client) Connection () *websocket.Conn{
	return c.webSock
}

func (c *Client) Write (message *Message){
	select {
	case c.msgChan <- message:
	default:
		c.server.DeleteClient(c)
		err := fmt.Errorf("client %d is disconnected.", c.id)
		c.server.Error(err)
	}
}

func (c *Client) Close (){
	c.closeChan <- true
}

func (c *Client) Listen (){
	go c.ListenWrite()
	c.ListenRead()
}

func (c *Client) ListenRead(){
	fmt.Println("Listening read from client")

	for {
		select {
		case <- c.closeChan:
			c.server.DeleteClient(c)
			c.closeChan <- true
			return
		default:
			var messege Message
			err := websocket.JSON.Receive(c.webSock, &messege)
			if err == io.EOF{
				c.closeChan <- true
			}else if err != nil{
				c.server.Error(err)
			}else {
				c.server.SendMsgToAll(&messege)
			}
		}
	}
}

func (c *Client) ListenWrite (){
	fmt.Println("Listening write to client")

	for {
		select {
		case msg := <-c.msgChan:
			fmt.Println("Send ", msg)
			websocket.JSON.Send(c.webSock, msg)

		case <-c.closeChan:
			c.server.DeleteClient(c)
			c.closeChan <-true
			return
		}
	}
}

