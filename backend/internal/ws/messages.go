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

type JoinPayload struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Role string `json:"role"`
	ClassID string `json:"class_id"`
}

type LeavePayload struct {
	UserID string `json:"user_id"`
	ClassID string `json:"class_id"`
	Reason string `json:"reason"`
}

type EngagmentUpdatePayload struct {
	UserID string `json:"user_id"`
	AttentionLevel float64 `json:"attention_level"`
	ConfusionLevel float64 `json:"confusion_level"`
	ParticipationRate float64 `json:"participation_rate"`
	Timestamp time.Time `json:"timestamp"`
}

type WhiteboardUpdatePayload struct {
	Action string `json:"action"`
	Data json.RawMessage `json:"data"`
	UserID string `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

type DrawData struct {
	Tool string `json:"tool"`
	Color string `json:"color"`
	Width int `json:"width"`
	Points []Point `json:"points"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ErrorPayload struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type SuccessPayload struct {
	Action string `json:"action"`
	Message string `json:"message"`
}
