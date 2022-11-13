package webhook

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const telegramURL = "https://api.telegram.org"

func SetWebhook() {
	godotenv.Load(".env")
	query := telegramURL + "/bot" + os.Getenv("BOT_TOKEN") + "/setWebhook?url=" + os.Getenv("WEBHOOK_URL")

	res, err := http.Get(query)

	if err != nil {
		log.Fatalf("Error: set a webhook. %v\n", err)
	} else if !(200 <= res.StatusCode && res.StatusCode < 300) {
		log.Fatalf("Error: invalid response. %v\n", res)
	} else {
		fmt.Printf("Success: Webhook is successfully set on domain {%s} with message.\n", os.Getenv("WEBHOOK_URL"))
	}
	res.Body.Close()
}
