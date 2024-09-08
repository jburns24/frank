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

func SendChat(apiKey string, question string) (string, error) {
	c := anthropic.NewClient(apiKey)
	resp, err := c.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaude3Dot5Sonnet20240620,
		Messages: []anthropic.Message{
			anthropic.NewUserTextMessage(question),
		},
		MaxTokens: 1024,
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
