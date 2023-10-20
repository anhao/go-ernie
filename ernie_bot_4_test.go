package go_ernie

import (
	"context"
	"testing"
)

func TestClient_CreateErnieBot4ChatCompletion(t *testing.T) {
	client := NewDefaultClient("", "")
	request := ErnieBot4Request{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "你是4.0吗？",
			},
		},
		Stream: false,
	}

	response, err := client.CreateErnieBot4ChatCompletion(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}
