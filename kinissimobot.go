package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

func parserTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	var update, err = parserTelegramRequest(r)
	if err != nil {
		log.Printf("Error parsing update, %s", err.Error())
		return
	}

	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, update.Message.Text)
	if errTelegram != nil {
		log.Printf("Got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("Message text %s successfuly distributed to chat id %d", update.Message.Text, update.Message.Chat.Id)
	}
}

func sendTextToTelegramChat(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})
	if err != nil {
		log.Printf("Error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("Error parsing telegram answare %s", errRead.Error())
	}

	bodyString := string(bodyBytes)
	log.Printf("Body of telegram response: %s", bodyString)

	return bodyString, nil
}
