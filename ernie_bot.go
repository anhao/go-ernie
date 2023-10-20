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

type ErnieFunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
	Thoughts  string `json:"thoughts"`
}

type ErnieFunctionExample struct {
	Role         string            `json:"role"`
	Content      string            `json:"content"`
	Name         string            `json:"name,omitempty"`
	FunctionCall ErnieFunctionCall `json:"function_call,omitempty"`
}

type ErnieFunction struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  any                    `json:"parameters"`
	Responses   any                    `json:"responses"`
	Examples    []ErnieFunctionExample `json:"examples,omitempty"`
}

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
	PenaltyScore    float32                 `json:"penalty_score,omitempty"`
	Functions       []ErnieFunction         `json:"functions,omitempty"`
}

type ErniePluginUsage struct {
	Name           string `json:"name"`
	ParseTokens    int    `json:"parse_tokens"`
	AbstractTokens int    `json:"abstract_tokens"`
	SearchTokens   int    `json:"search_tokens"`
	TotalTokens    int    `json:"total_tokens"`
}

type ErnieUsage struct {
	PromptTokens     int                `json:"prompt_tokens"`
	CompletionTokens int                `json:"completion_tokens"`
	TotalTokens      int                `json:"total_tokens"`
	Plugins          []ErniePluginUsage `json:"plugins"`
}

type ErnieBotResponse struct {
	Id               string            `json:"id"`
	Object           string            `json:"object"`
	Created          int               `json:"created"`
	SentenceId       int               `json:"sentence_id"`
	IsEnd            bool              `json:"is_end"`
	IsTruncated      bool              `json:"is_truncated"`
	Result           string            `json:"result"`
	NeedClearHistory bool              `json:"need_clear_history"`
	Usage            ErnieUsage        `json:"usage"`
	FunctionCall     ErnieFunctionCall `json:"function_call"`
	BanRound         int               `json:"ban_round"`
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
