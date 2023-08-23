package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestParsingUpdateMessageWithText(t *testing.T) {
	var msg = Message{
		Text: "Hello, World!",
		Chat: Chat{Id: 4343},
	}
	var update = Update{
		UpdateId: 1,
		Message:  msg,
	}

	requestBody, err := json.Marshal(update)
	if err != nil {
		t.Errorf("Failed to marshal update in json, got %s", err.Error())
	}
	req := httptest.NewRequest("POST", "http://myTelegramHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	var updateToTest, errPars = parserTelegramRequest(req)
	if err != nil {
		t.Errorf("Expected <nil> error, got %s", errPars.Error())
	}
	if *updateToTest != update {
		t.Errorf("Expected %v got %v", update, updateToTest)
	}

}
