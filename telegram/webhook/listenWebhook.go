package webhook

import (
	"fmt"
	"headHunterBot/hh"
	"headHunterBot/myLog"
	"headHunterBot/telegram"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func ListenToWebhook() {
	e := echo.New()

	messageTexts := make(map[int]map[int]string)

	e.POST("/webhook", func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		myLog.LogWebhookJson(reqBody)
		fmt.Println(reqBody)

		if err != nil {
			errorText := fmt.Sprintf("%v\n", err)
			myLog.LogWebhookText(errorText)
			return c.String(http.StatusBadRequest, errorText)
		}

		if message := telegram.ParseMessage(reqBody); message != nil {
			if messageTexts[message.Message.From.Id] == nil {
				messageTexts[message.Message.From.Id] = make(map[int]string)
			}
			if message.Message.Text[0] == '/' {
				id := message.Message.Text[1:]
				returnMessage := hh.GetVacancyInformationById(id)
				telegram.SendPlainText(message.Message.From.Id, returnMessage)
			} else {
				message.Message.MessageId++
				messageTexts[message.Message.From.Id][message.Message.MessageId] = message.Message.Text
				returnMessage, items := hh.GetVacanciesTextByRange(message.Message.Text, 0)
				telegram.SendMessage(message.Message.From.Id, 0, items, message.Message.MessageId, returnMessage)
			}
		} else if callbackQuery := telegram.ParseCallbackQuery(reqBody); callbackQuery != nil {
			message := callbackQuery.CallbackQuery.Message
			if err != nil {
				telegram.SendMessage(message.Chat.Id, -1, -1, message.MessageId, "Неопознанная команда")
				return c.String(http.StatusOK, "OK")
			}
			messageId, index := telegram.EncodeQueryString(callbackQuery.CallbackQuery.Data)
			returnMessage, items := hh.GetVacanciesTextByRange(messageTexts[callbackQuery.CallbackQuery.Message.Chat.Id][*messageId], *index)
			telegram.EditMessage(message.Chat.Id, *index, items, *messageId, returnMessage)
		} else {
			return c.String(http.StatusOK, "OK")
		}

		return c.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
