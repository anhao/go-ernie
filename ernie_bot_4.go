package go_ernie

import (
	"context"
	"net/http"
)

const ernieBot4URL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro"

type ErnieBot4Request struct {
	Messages     []ChatCompletionMessage `json:"messages"`
	Temperature  float32                 `json:"temperature,omitempty"`
	TopP         float32                 `json:"top_p,omitempty"`
	Stream       bool                    `json:"stream"`
	UserId       string                  `json:"user_id,omitempty"`
	Functions    []ErnieFunction         `json:"functions,omitempty"`
	PenaltyScore float32                 `json:"penalty_score,omitempty"`
	System       string                  `json:"system,omitempty"`
}

type ErnieBot4Response struct {
	ErnieBotResponse
}

func (c *Client) CreateErnieBot4ChatCompletion(
	ctx context.Context,
	request ErnieBot4Request,
) (response ErnieBot4Response, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernieBot4URL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
