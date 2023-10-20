package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

const completionURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/completions/"

// CompletionRequest 文本续写模型
type CompletionRequest struct {
	Model  string `json:"-"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

type CompletionResponse struct {
	Id         string     `json:"id"`
	Object     string     `json:"object"`
	Created    int        `json:"created"`
	SentenceId int        `json:"sentence_id"`
	IsEnd      bool       `json:"is_end"`
	Result     string     `json:"result"`
	IsSafe     bool       `json:"is_safe"`
	Usage      ErnieUsage `json:"usage"`
	APIError
}

func (c *Client) CreateCompletion(
	ctx context.Context,
	request CompletionRequest,
) (response CompletionResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", completionURL, request.Model)), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
