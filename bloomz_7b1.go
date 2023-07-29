package go_ernie

import (
	"context"
	"net/http"
)

const bloomz7b1URL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/bloomz_7b1"

type Bloomz7b1Request struct {
	Messages []ChatCompletionMessage `json:"messages"`
	Stream   bool                    `json:"stream"`
	UserId   string                  `json:"user_id"`
}

type Bloomz7b1Response struct {
	ErnieBotResponse
}

func (c *Client) CreateBloomz7b1ChatCompletion(
	ctx context.Context,
	request Bloomz7b1Request,
) (response Bloomz7b1Response, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(bloomz7b1URL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
