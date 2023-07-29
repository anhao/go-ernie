package go_ernie

import (
	"context"
	"net/http"
)

type ErnireBotTurboChatCompletionStream struct {
	*streamReader[ErnieBotResponse]
}

func (c *Client) CreateErnieBotTurboChatCompletionStream(
	ctx context.Context,
	request ErnieBotTurboRequest,
) (stream *ErnireBotChatCompletionStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernireBotTurboURL), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[ErnieBotTurboResponse](c, req)
	if err != nil {
		return
	}
	stream = &ErnireBotChatCompletionStream{
		streamReader: resp,
	}
	return
}
