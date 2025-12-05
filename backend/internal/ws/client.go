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
	//handles client disconnect
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		log.Printf("Client %s disconnected from class %s", c.userID, c.classID)
	} ()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	//infinite loop that can handle messages on the same websocket connection
	for {
		_, messageBytes, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Websocket err for user %s: %s", c.userID, err)
			}
			break
		}

		//parsing the message
		var msg Message
		if err := json.Unmarshal(messageBytes, &msg); err != nil {
			continue
		}

		if msg.ClassID != c.classID { continue}
		if msg.UserID != c.userID{ continue}

	}
	
}

