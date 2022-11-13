package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"headHunterBot/myLog"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const telegramURL = "https://api.telegram.org"

type MessageData struct {
	ChatId      int                  `json:"chat_id"`
	MessageId   int                  `json:"message_id,omitempty"`
	Text        string               `json:"text"`
	ParseMode   string               `json:"parse_mode,omitempty"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type PlainMessage struct {
	ChatId    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

func returnInlineKeyboardMarkup(index, items, messageId int) InlineKeyboardMarkup {
	messageIdString := strconv.Itoa(messageId)

	pages := (items + 6) / 7
	index++

	if pages <= 5 {

		keyboard := [][]InlineKeyboardButton{{}}
		for i := 1; i <= pages; i++ {
			var buttonText string
			numberString := strconv.Itoa(i)

			if i == index {
				buttonText = "- " + numberString + " -"
			} else {
				buttonText = numberString
			}

			keyboard[0] = append(keyboard[0], InlineKeyboardButton{buttonText, messageIdString + "-" + strconv.Itoa(i-1)})
		}
		return InlineKeyboardMarkup{keyboard}
	} else {
		keyboard := [][]InlineKeyboardButton{{}}

		if 1 <= index-2 && index+2 <= pages {
			for i, j := index-2, 0; i <= index+2; i, j = i+1, j+1 {
				var buttonText string
				numberString := strconv.Itoa(i)

				if j == 2 {
					buttonText = "- " + numberString + " -"
				} else if j == 0 {
					if index-2 > 1 {
						buttonText = "< " + numberString
					} else {
						buttonText = numberString
					}
				} else if j == 4 {
					if index+2 < pages {
						buttonText = numberString + " >"
					} else {
						buttonText = numberString
					}
				} else {
					buttonText = numberString
				}

				keyboard[0] = append(keyboard[0], InlineKeyboardButton{buttonText, messageIdString + "-" + strconv.Itoa(i-1)})
			}
		} else if index <= 2 {
			for i, j := 1, 0; i <= 5; i, j = i+1, j+1 {
				var buttonText string
				numberString := strconv.Itoa(i)

				if index == i {
					buttonText = "- " + numberString + " -"
				} else if j == 4 {
					buttonText = numberString + " >"
				} else {
					buttonText = numberString
				}

				keyboard[0] = append(keyboard[0], InlineKeyboardButton{buttonText, messageIdString + "-" + strconv.Itoa(i-1)})
			}
		} else {
			for i, j := pages-4, 0; i <= pages; i, j = i+1, j+1 {
				var buttonText string
				numberString := strconv.Itoa(i)

				if index == i {
					buttonText = "- " + numberString + " -"
				} else if j == 0 {
					buttonText = "< " + numberString
				} else {
					buttonText = numberString
				}

				keyboard[0] = append(keyboard[0], InlineKeyboardButton{buttonText, messageIdString + "-" + strconv.Itoa(i-1)})
			}
		}
		return InlineKeyboardMarkup{keyboard}
	}
}

func postMessage(chatId int, messageQuery string, message MessageData) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(message)
	if err != nil {
		log.Fatalf("client: error encoding body. %v\n", err)
	}
	query := telegramURL + "/bot" + os.Getenv("BOT_TOKEN") + "/" + messageQuery
	req, err := http.NewRequest(http.MethodPost, query, &buf)

	if err != nil {
		log.Fatalf("client: error assigning http request: %s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}
	body, _ := io.ReadAll(res.Body)
	myLog.LogTelegramResponses(body)
}

func SendMessage(chatId int, index, items, messageId int, text string) {
	message := MessageData{
		ChatId:      chatId,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: returnInlineKeyboardMarkup(index, items, messageId),
	}
	postMessage(chatId, "sendMessage", message)
}

func EditMessage(chatId int, index, items, messageId int, text string) {
	message := MessageData{
		ChatId:      chatId,
		MessageId:   messageId,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: returnInlineKeyboardMarkup(index, items, messageId),
	}
	postMessage(chatId, "editMessageText", message)
}

func SendPlainText(chatId int, text string) {
	message := PlainMessage{
		ChatId:    chatId,
		ParseMode: "HTML",
		Text:      text,
	}
	fmt.Printf("Kyky %+v\n", message)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(message)
	if err != nil {
		log.Fatalf("client: error encoding body. %v\n", err)
	}
	query := telegramURL + "/bot" + os.Getenv("BOT_TOKEN") + "/sendMessage"
	req, err := http.NewRequest(http.MethodPost, query, &buf)
	// bodyyy, _ := io.ReadAll(req.Body)
	// fmt.Println(string(bodyyy))

	if err != nil {
		log.Fatalf("client: error assigning http request: %s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}
	body, _ := io.ReadAll(res.Body)
	myLog.LogTelegramResponses(body)
}
