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
	pingPeriod = (pongWait * 9) / 10
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

		switch msg.Type {
		case MessageTypeJoin:
			log.Printf("User %s joined class %s", c.userID, c.classID)
		case MessageTypeEngagmentUpdate:
			var payload EngagmentUpdatePayload
			if err := ParsePayload(&msg, &payload); err != nil {
				c.sendError("Invalid Payload", "Failed to parse message")
				continue
			}
			log.Printf("Engagment update from user %s, attention=%.1f,  confusion=%.1f", c.userID, payload.AttentionLevel, payload.ConfusionLevel)
			c.hub.broadcast <- &msg
		case MessageTypeWhiteboardUpdate:
			var payload WhiteboardUpdatePayload
			if err := ParsePayload(&msg, &payload); err != nil {
				c.sendError("Invalid Payload", "Message class ID doesn't match connection")
				continue
			}
			log.Printf("Whiteboard Update from user %s, action:%s", c.userID, payload.Action)
			c.hub.broadcast <- &msg
		case MessageTypeLeave:
			log.Printf("User %s left class %s", c.userID, c.classID)
			return
		default:
			log.Printf("Unknown message from %s, class %s", c.userID, c.classID)
			c.sendError("Unknown Type", "Cannot send messages as another user")
		}
	}
}

//write pump, handles outgoing messages
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	} ()

	for {
		select {
		case message, ok := <- c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			//message optimization
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		//Time to send a Ping
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//helper methods
func (c *Client) sendError(code, message string) {
	errorPayload := ErrorPayload{
		Code: code,
		Message: message,
	}

	msg, err := NewMessage(MessageTypeError, errorPayload, "system", c.classID)
	if err != nil {
		log.Printf("Failed to create error message,%v", err)
		return
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal error message, %v", err)
	}

	//non-blocking send
	select {
	case c.send <-msgBytes:
	default:
		log.Printf("Failed to send error to client")
	}
}

func (c *Client) sendSuccess(action, message string) {
	successPayload := SuccessPayload{
		Action:  action,
		Message: message,
	}

	msg, err := NewMessage(MessageTypeSuccess, successPayload, "system", c.classID)
	if err != nil {
		log.Printf("Failed to create success message: %v", err)
		return
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal success message: %v", err)
		return
	}

	select {
	case c.send <- msgBytes:
	default:
		log.Printf("Failed to send success to client %s: send channel full", c.userID)
	}
}
