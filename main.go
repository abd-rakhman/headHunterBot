package main

import (
	"headHunterBot/telegram/webhook"
)

func main() {
	webhook.SetWebhook()
	webhook.ListenToWebhook()
}
