package example

import (
	"google.golang.org/protobuf/runtime/protoimpl"
	"testing"
)

func TestHello_AsPublic(t *testing.T) {
	h := &Hello{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Name:          "Public Name",
		PrivateString: "aaasdf",
		PrivateInt:    123123,
		PrivateMessage: &PrivateMessage{
			Password: "Secret",
		},
		PartialPrivateMessage: &PartialPrivateMessage{
			Password:     "secret",
			PublicString: "public",
		},
	}

	h.AsPublic()

	if h.Name != "Public Name" {
		t.Errorf("Name should be Public Name, got %s", h.Name)
	}

	if h.PrivateString != "" {
		t.Errorf("PrivateString should be empty, got %s", h.PrivateString)
	}

	if h.PrivateInt != 0 {
		t.Errorf("PrivateInt should be 0, got %d", h.PrivateInt)
	}

	if h.PrivateMessage != nil {
		t.Errorf("PrivateMessage should be nil, got %v", h.PrivateMessage)
	}

	if h.PartialPrivateMessage.Password != "" {
		t.Errorf("PartialPrivateMessage.Password should be empty, got %s", h.PartialPrivateMessage.Password)
	}

	if h.PartialPrivateMessage.PublicString != "public" {
		t.Errorf("PartialPrivateMessage.PublicString should be public, got %s", h.PartialPrivateMessage.PublicString)
	}
}

func TestPrivateMessage_AsPublic(t *testing.T) {
	subject := &PrivateMessage{
		Password: "secret",
	}

	subject.AsPublic()
	if subject.Password != "secret" {
		t.Errorf("Password should be secret, got %s", subject.Password)
	}
}
