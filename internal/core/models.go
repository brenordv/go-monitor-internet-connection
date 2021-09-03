package core

import (
	"bytes"
	"encoding/gob"
	"time"
)

type RuntimeConfig struct {
	Headers              map[string]interface{} `json:"headers"`
	DelayInSeconds       int                    `json:"delayInSeconds"`
	NoConnInfoTTLinHours int                    `json:"noConnInfoTTLinHours"`
	TelegramChatId       int64                  `json:"telegramChatId"`
	TelegramBotToken     string                 `json:"telegramBotToken"`
}

type NoConnectionInfo struct {
	Start    time.Time
	End      time.Time
	DownTime time.Duration
	Url      string
	Error    error
}

func (n *NoConnectionInfo) Serialize() ([]byte, []byte, error) {
	var res bytes.Buffer
	key := bytes.Join([][]byte{[]byte(n.Start.Format(time.RFC3339)), []byte(n.Url)}, []byte{})
	enc := gob.NewEncoder(&res)
	err := enc.Encode(n)

	return key, res.Bytes(), err
}
