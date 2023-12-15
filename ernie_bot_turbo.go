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
	PenaltyScore    float32                 `json:"penalty_score,omitempty"`
	System          string                  `json:"system,omitempty"`
}
type ErnieBotTurboUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ErnieBotTurboResponse struct {
	Id               string             `json:"id"`
	Object           string             `json:"object"`
	Created          int                `json:"created"`
	SentenceId       int                `json:"sentence_id"`
	IsEnd            bool               `json:"is_end"`
	IsTruncated      bool               `json:"is_truncated"`
	Result           string             `json:"result"`
	NeedClearHistory bool               `json:"need_clear_history"`
	Usage            ErnieBotTurboUsage `json:"usage"`
	BanRound         int                `json:"ban_round"`
	APIError
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
