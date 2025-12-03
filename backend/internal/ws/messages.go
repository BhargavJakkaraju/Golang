package ws

import (
	"encoding/json"
	"time"
)

type MessageType string

const (
	//connection message
	MessageTypeJoin MessageType = "join"
	MessageTypeLeave MessageType = "leave"

	//engagment message
	MessageTypeEngagmentUpdate MessageType = "engagment_update"
	MessageTypeWhiteboardUpdate MessageType = "whiteboard_update"

	//status message
	MessageTypeError MessageType = "error"
	MessageTypeSuccess MessageType = "success"
)

type Message struct {
	//determines message type
	Type MessageType `json:"type"`

	//contains actual data
	Payload json.RawMessage `json:"payload"`

	//Metadata
	UserID string `json:"user_id"`
	ClassID string `json:"class_id"`
	Timestamp time.Time `json:"timestamp"`

}

