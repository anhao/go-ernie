package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

const ernireCustomPluginUrl = "/rpc/2.0/ai_custom/v1/wenxinworkshop/plugin/"

type ErnieCustomPluginRequest struct {
	PluginName     string                  `json:"-"`
	Query          string                  `json:"query"`
	Stream         bool                    `json:"stream"`
	Plugins        []string                `json:"plugins"`
	LLM            any                     `json:"llm,omitempty"`
	InputVariables any                     `json:"input_variables,omitempty"`
	History        []ChatCompletionMessage `json:"history,omitempty"`
	Verbose        bool                    `json:"verbose,omitempty"`
	FileUrl        string                  `json:"fileurl,omitempty"`
}

type ErnieCustomPluginResponse struct {
	LogId            int64      `json:"log_id"`
	Id               string     `json:"id"`
	Object           string     `json:"object"`
	Created          int        `json:"created"`
	SentenceId       int        `json:"sentence_id"`
	IsEnd            bool       `json:"is_end"`
	Result           string     `json:"result"`
	NeedClearHistory bool       `json:"need_clear_history"`
	BanRound         int        `json:"ban_round"`
	Usage            ErnieUsage `json:"usage"`
	MetaInfo         any        `json:"meta_info"`
	APIError
}

func (c *Client) CreateErnieCustomPluginChatCompletion(
	ctx context.Context,
	request ErnieCustomPluginRequest,
) (response ErnieCustomPluginResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s/", ernireCustomPluginUrl, request.PluginName)), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
