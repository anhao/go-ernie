package go_ernie

import (
	"context"
	"testing"
)

func TestClient_CreateEmbeddings(t *testing.T) {
	client := NewClient("xxx")

	request := EmbeddingRequest{
		Input: []string{
			"Hello",
		},
		UserId: "",
	}

	embeddings, err := client.CreateEmbeddings(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(embeddings)
}
