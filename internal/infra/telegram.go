package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gabemanfroi/notificationWebhook/domain/model"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/core"
	"net/http"
)

func getTelegramUrl() string {
	fmt.Println(core.AppConfig.Telegram.Token)
	return fmt.Sprintf("https://api.telegram.org/bot%s", core.AppConfig.Telegram.Token)
}

func NotifyAlertTelegram(alert model.GetAnalysisByAlertIdResponse) {
	sendTelegramMessage(FormatMessage(alert))
}

func sendTelegramMessage(message string) {
	url := fmt.Sprintf("%s/sendMessage", getTelegramUrl())

	body, _ := json.Marshal(map[string]string{
		"chat_id": "6700295368",
		"text":    message,
	})

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	fmt.Println(response.Body)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

}
