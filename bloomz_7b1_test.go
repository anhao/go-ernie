package go_ernie

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestClient_CreateBloomz7b1ChatCompletion(t *testing.T) {
	client := NewClient("xxx")
	request := Bloomz7b1Request{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: false,
	}

	_, err := client.CreateBloomz7b1ChatCompletion(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClient_CreateBloomz7b1ChatCompletionStream(t *testing.T) {
	client := NewClient("xxxx")
	request := Bloomz7b1Request{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: false,
	}

	response, err := client.CreateBloomz7b1ChatCompletionStream(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	defer response.Close()
	for {
		recv, err := response.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("eof")
			return
		}
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(recv)
	}
}
