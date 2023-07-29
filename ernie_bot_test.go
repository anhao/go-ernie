package go_ernie

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestClient_CreateErnieBotChatCompletion(t *testing.T) {
	client := NewClient("xxx")
	request := ErnieBotRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: false,
	}

	response, err := client.CreateErnieBotChatCompletion(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}
func TestName(t *testing.T) {
	client := NewClient("xxx")
	request := ErnieBotRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: true,
	}

	response, err := client.CreateErnieBotChatCompletionStream(context.Background(), request)
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
