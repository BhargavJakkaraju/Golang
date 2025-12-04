package ws

import (
	"testing"
)

func TestMessageCreattion(t *testing.T) {
	joinPayload := JoinPayload{
		UserID: "user123",
		Username: "Bhargav",
		Role: "student",
		ClassID: "class456",

	}

	msg, err := NewMessage(MessageTypeJoin, joinPayload, "user123", "class456")
	if err != nil {
		t.Fatalf("Failed to create message: %v", err)
	}

	if msg.Type != MessageTypeJoin {
		t.Fatalf("Expected type %s, received %s", MessageTypeJoin, msg.Type)
	}

	var parsedPayload JoinPayload
	err = ParsePayload(msg, &parsedPayload)
	if err != nil {
		t.Fatalf("Failed to parse payload %v", err)
	}

	if parsedPayload.Username != "Bhargav" {
		t.Errorf("Not expected Username")
	}
}