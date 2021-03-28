package message

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "este es un mensaje secreto"
	messageOne := []string{"este", "", "", "", "secreto"}
	messageTwo := []string{"", "es", "", "mensaje", ""}
	messageThree := []string{"", "", "un", "", ""}
	message := NewService().GetMessage(messageOne, messageTwo, messageThree)
	if message != expected {
		t.Errorf("Expected: %v, got: %v", expected, message)
	}
}
