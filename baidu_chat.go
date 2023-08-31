package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

const baiduChatURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/"

type BaiduChatRequest struct {
	Messages []ChatCompletionMessage `json:"messages"`
	Stream   bool                    `json:"stream"`
	UserId   string                  `json:"user_id"`
	Model    string                  `json:"-"`
}

type BaiduChatResponse struct {
	ErnieBotResponse
}

func (c *Client) CreateBaiduChatCompletion(
	ctx context.Context,
	request BaiduChatRequest,
) (response BaiduChatResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", baiduChatURL, request.Model)), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
