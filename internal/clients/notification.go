package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type telegramMsgRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func Notify(msgContent string, token string, chatId int64) error {
	if token == "" || chatId == 0 {
		return nil
	}

	msg := &telegramMsgRequest{
		ChatID: chatId,
		Text:   msgContent,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	client := GetHttpClient(true)
	res, rErr := client.Post(url, "application/json", bytes.NewBuffer(msgBytes))

	if rErr != nil {
		return rErr
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %d | %s\n", res.StatusCode, res.Status)
	}

	return nil
}