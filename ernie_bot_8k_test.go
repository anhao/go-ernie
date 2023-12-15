package go_ernie

import (
	"context"
	"testing"
)

func TestClient_CreateErnieBot8KChatCompletion(t *testing.T) {
	client := NewClient("")
	request := ErnieBot8KRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: false,
	}

	response, err := client.CreateErnieBot8KChatCompletion(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}
