package go_ernie

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestClient_CreateErnieBotChatCompletion(t *testing.T) {
	client := NewClient("")
	request := ErnieBotRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "南昌的天气怎么样",
			},
		},
		Stream: false,
		Functions: []ErnieFunction{
			{
				Name:        "get_weather",
				Description: "获取天气信息",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"location": map[string]interface{}{
							"type":        "string",
							"description": "要查询的城市地址",
						},
					},
				},
			},
		},
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
