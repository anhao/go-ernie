package go_ernie

import (
	"context"
	"net/http"
)

const embeddingURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/embeddings/embedding-v1"

type EmbeddingRequest struct {
	Input  []string `json:"input"`
	UserId string   `json:"user_id"`
}
type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}
type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbeddingResponse struct {
	Id      string          `json:"id"`
	Object  string          `json:"object"`
	Created int             `json:"created"`
	Data    []EmbeddingData `json:"data"`
	Usage   EmbeddingUsage  `json:"usage"`
	APIError
}

func (c *Client) CreateEmbeddings(ctx context.Context, request EmbeddingRequest) (response EmbeddingResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(embeddingURL), withBody(request))
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)
	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
