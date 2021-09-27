package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type SlackRequestBody struct {
	Text     string `json:"text"`
	IconUrl  string `json:"icon_url"`
	UserName string `json:"username"`
}

type Attachment struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
}

func SlackLog(msg string) error {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}

	webhookUrl := os.Getenv("WEBHOOK_URL")

	slackBody, _ := json.Marshal(SlackRequestBody{
		Text:     msg,
		IconUrl:  "https://img.icons8.com/emoji/96/000000/penguin--v2.png",
		UserName: "Gấp Gấp",
	})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		fmt.Println(buf.String())
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
