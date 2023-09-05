package go_ernie

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestClient_CreateErnieCustomPluginChatCompletion(t *testing.T) {
	client := NewClient("xxx")
	request := ErnieCustomPluginRequest{
		PluginName: "xxx",
		Query:      "你好",
		Stream:     false,
	}

	response, err := client.CreateErnieCustomPluginChatCompletion(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}

func TestClient_CreateErnieCustomPluginStreamChatCompletion(t *testing.T) {
	client := NewClient("xxx")
	request := ErnieCustomPluginRequest{
		PluginName: "xxx",
		Query:      "你好?",
		Stream:     true,
	}

	response, err := client.CreateErnieCustomPluginStreamChatCompletion(context.Background(), request)
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
