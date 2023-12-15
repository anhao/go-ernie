package go_ernie

import (
	"context"
	"net/http"
)

const ernieBot8KURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ernie_bot_8k"

type ErnieBot8KRequest struct {
	Messages       []ChatCompletionMessage `json:"messages"`
	Temperature    float32                 `json:"temperature,omitempty"`
	TopP           float32                 `json:"top_p,omitempty"`
	Stream         bool                    `json:"stream"`
	UserId         string                  `json:"user_id,omitempty"`
	Functions      []ErnieFunction         `json:"functions,omitempty"`
	PenaltyScore   float32                 `json:"penalty_score,omitempty"`
	System         string                  `json:"system,omitempty"`
	Stop           []string                `json:"stop,omitempty"`
	DisableSearch  bool                    `json:"disable_search,omitempty"`
	EnableCitation bool                    `json:"enable_citation,omitempty"`
}

type ErnieBot8KResponse struct {
	ErnieBotResponse
}

func (c *Client) CreateErnieBot8KChatCompletion(
	ctx context.Context,
	request ErnieBot8KRequest,
) (response ErnieBot8KResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernieBot8KURL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
