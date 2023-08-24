package go_ernie

import (
	"context"
	"net/http"
)

const ernireBotTurboURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant"

type ErnieBotTurboRequest struct {
	Messages        []ChatCompletionMessage `json:"messages"`
	Stream          bool                    `json:"stream"`
	UserId          string                  `json:"user_id"`
	Temperature     float32                 `json:"temperature,omitempty"`
	TopP            float32                 `json:"top_p,omitempty"`
	PresencePenalty float32                 `json:"presence_penalty,omitempty"`
}

type ErnieBotTurboResponse struct {
	ErnieBotResponse
}

func (c *Client) CreateErnieBotTurboChatCompletion(
	ctx context.Context,
	request ErnieBotTurboRequest,
) (response ErnieBotTurboResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernireBotTurboURL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
