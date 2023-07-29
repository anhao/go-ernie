package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

const llamaChatURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/"

type LlamaChatRequest struct {
	Messages []ChatCompletionMessage `json:"messages"`
	Stream   bool                    `json:"stream"`
	UserId   string                  `json:"user_id"`
	Model    string                  `json:"-"`
}
type LlamaChatResponse struct {
	ErnieBotResponse
}

func (c *Client) CreateLlamaChatCompletion(
	ctx context.Context,
	request LlamaChatRequest,
) (response LlamaChatResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", llamaChatURL, request.Model)), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
