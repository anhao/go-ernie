package go_ernie

import (
	"context"
	"net/http"
)

const bgeLargeZhURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/embeddings/bge_large_zh"

func (c *Client) CreateBgeLargeZhEmbedding(ctx context.Context, request EmbeddingRequest) (response EmbeddingResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(bgeLargeZhURL), withBody(request))
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)
	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
