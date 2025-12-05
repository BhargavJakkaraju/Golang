package ws

import (
	"encoding/json"
	"log"
	"time"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingWait = (pongWait * 9) / 10
	maxMessageSize = 10 * 1024 * 1024
)

type Client struct {
	conn *websocket.Conn
	hub *Hub 
	send chan[] byte

	userID string
	classID string
	role string
}

//readPump handles incoming requests from the websocket
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		log.Printf("Client %s disconnected from class %s", c.userID, c.classID)
	} ()
}

