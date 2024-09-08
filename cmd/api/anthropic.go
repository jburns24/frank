package api

import (
	"context"
	"errors"
	"fmt"

	anthropic "github.com/liushuangls/go-anthropic/v2"
)

var (
	client anthropic.Client
)

func SendChat(apiKey string, message string) (string, error) {
	client := anthropic.NewClient(apiKey)
	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaudeInstant1Dot2,
		Messages: []anthropic.Message{
			anthropic.NewUserTextMessage(message),
		},
		MaxTokens: 1000,
	})
	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			fmt.Printf("Messages error, type: %s, message: %s", e.Type, e.Message)
		} else {
			fmt.Printf("Messages error: %v\n", err)
		}
		return "", nil
	}
	return resp.Content[0].GetText(), nil
}
