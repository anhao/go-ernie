package go_ernie

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestClient_CreateErnieBotTurboChatCompletion(t *testing.T) {
	client := NewClient("xxx")
	request := ErnieBotTurboRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream:      false,
		Temperature: 0.1,
	}

	response, err := client.CreateErnieBotTurboChatCompletion(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}
func TestClient_CreateErnieBotTurboChatCompletionStream(t *testing.T) {
	client := NewClient("xxx")
	request := ErnieBotTurboRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: false,
	}

	response, err := client.CreateErnieBotTurboChatCompletionStream(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	defer response.Close()
	for {
		recv, err := response.Recv()
		if errors.Is(err, io.EOF) {
			t.Log("EOF")
			return
		}
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(recv)
	}
}
