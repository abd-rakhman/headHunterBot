package telegram

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Message struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
		} `json:"from"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

type CallbackQuery struct {
	UpdateId      int `json:"update_id"`
	CallbackQuery struct {
		Message struct {
			MessageId int `json:"message_id"`
			From      struct {
				Id        int    `json:"id"`
				FirstName string `json:"first_name"`
				IsBot     bool   `json:"is_bot"`
			} `json:"from"`
			Chat struct {
				Id        int    `json:"id"`
				FirstName string `json:"first_name"`
			} `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
		Data string `json:"data"`
	} `json:"callback_query"`
}

func ParseMessage(data []byte) *Message {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		fmt.Printf("Parse did not work on message by data: %s\n", data)
	}
	if message.Message.MessageId == 0 {
		return nil
	} else {
		return &message
	}
}

func ParseCallbackQuery(data []byte) *CallbackQuery {
	var callbackQuery CallbackQuery
	err := json.Unmarshal(data, &callbackQuery)
	if err != nil {
		fmt.Printf("Parse did not work on message by data: %s\n", data)
	}
	if callbackQuery.UpdateId == 0 {
		return nil
	} else {
		return &callbackQuery
	}
}
