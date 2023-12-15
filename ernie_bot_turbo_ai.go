package go_ernie

import (
	"context"
	"net/http"
)

const ernireBotTurboAIURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ai_apaas"

type ErnieBotTurboAIRequest struct {
	Messages        []ChatCompletionMessage `json:"messages"`
	Stream          bool                    `json:"stream"`
	UserId          string                  `json:"user_id"`
	Temperature     float32                 `json:"temperature,omitempty"`
	TopP            float32                 `json:"top_p,omitempty"`
	PresencePenalty float32                 `json:"presence_penalty,omitempty"`
	PenaltyScore    float32                 `json:"penalty_score,omitempty"`
	System          string                  `json:"system,omitempty"`
}

type ErnieBotTurboAIResponse struct {
	ErnieBotTurboResponse
}

func (c *Client) CreateErnieBotTurboAIChatCompletion(
	ctx context.Context,
	request ErnieBotTurboAIRequest,
) (response ErnieBotTurboAIResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernireBotTurboAIURL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
