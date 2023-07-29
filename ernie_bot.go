package go_ernie

import (
	"context"
	"errors"
	"net/http"
)

const (
	MessageRoleUser      = "user"
	MessageRoleAssistant = "assistant"
)

var (
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream") //nolint:lll
)

const ernieBotURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions"

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ErnieBotRequest struct {
	Messages        []ChatCompletionMessage `json:"messages"`
	Temperature     float32                 `json:"temperature,omitempty"`
	TopP            float32                 `json:"top_p,omitempty"`
	PresencePenalty float32                 `json:"presence_penalty,omitempty"`
	Stream          bool                    `json:"stream"`
	UserId          string                  `json:"user_id,omitempty"`
}

type ErnieUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ErnieBotResponse struct {
	Id               string     `json:"id"`
	Object           string     `json:"object"`
	Created          int        `json:"created"`
	SentenceId       int        `json:"sentence_id"`
	IsEnd            bool       `json:"is_end"`
	IsTruncated      bool       `json:"is_truncated"`
	Result           string     `json:"result"`
	NeedClearHistory bool       `json:"need_clear_history"`
	Usage            ErnieUsage `json:"usage"`
	APIError
}

func (c *Client) CreateErnieBotChatCompletion(
	ctx context.Context,
	request ErnieBotRequest,
) (response ErnieBotResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernieBotURL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
