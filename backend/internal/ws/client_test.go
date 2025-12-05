package ws

import (
	"testing"
	"time"
)


func TestClientCreation(t *testing.T) {
	client := &Client{
		//hub:     nil,
		conn:    nil, 
		send:    make(chan []byte, 256),
		userID:  "user123",
		classID: "class456",
		role:    "student",
	}

	if client.userID != "user123" {
		t.Errorf("Expected userID 'user123', got '%s'", client.userID)
	}

	if client.classID != "class456" {
		t.Errorf("Expected classID 'class456', got '%s'", client.classID)
	}

	if client.role != "student" {
		t.Errorf("Expected role 'student', got '%s'", client.role)
	}

	if cap(client.send) != 256 {
		t.Errorf("Expected send channel capacity 256, got %d", cap(client.send))
	}
}


func TestClientSendChannel(t *testing.T) {
	client := &Client{
		send: make(chan []byte, 256),
	}

	testMessage := []byte("test message")

	
	select {
	case client.send <- testMessage:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout sending to channel - should not block")
	}

	select {
	case msg := <-client.send:
		if string(msg) != string(testMessage) {
			t.Errorf("Expected '%s', got '%s'", testMessage, msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout receiving from channel")
	}
}

func TestClientSendChannelBuffering(t *testing.T) {
	client := &Client{
		send: make(chan []byte, 3), 
	}

	
	for i := 0; i < 3; i++ {
		select {
		case client.send <- []byte("message"):
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("Message %d should not block with buffer of 3", i+1)
		}
	}

	select {
	case client.send <- []byte("message 4"):
		t.Fatal("4th message should block when buffer is full")
	case <-time.After(100 * time.Millisecond):
	}

	<-client.send

	select {
	case client.send <- []byte("message 5"):
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Should be able to send after reading from channel")
	}
}