package myLog

import (
	"encoding/json"
	"fmt"
	"os"
)

func log(fileName, text string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error: couldn't open hhResponses.txt")
		return
	}

	defer file.Close()

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Printf("Error: couldn't write %s", fileName)
	}
}

func LogHHResponse(requestText string, responseText string) {
	text := fmt.Sprintf("requestText: %s\nresponseText: %s\n", requestText, responseText)
	log("myLog/hhResponses.txt", text)
}

func LogWebhookJson(body []byte) {
	text, _ := json.MarshalIndent(string(body), "", "\t")
	log("myLog/webhookMessages.txt", string(text)+"\n")
}

func LogWebhookText(text string) {
	log("myLog/webhookMessages.txt", text)
}

func LogTelegramResponses(body []byte) {
	text, _ := json.MarshalIndent(string(body), "", "\t")
	log("myLog/telegramResponses.txt", string(text)+"\n")
}
