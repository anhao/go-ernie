package go_ernie

import (
	"context"
	"net/http"
)

type ErnireBotChatCompletionStream struct {
	*streamReader[ErnieBotTurboResponse]
}

func (c *Client) CreateErnieBotChatCompletionStream(
	ctx context.Context,
	request ErnieBotRequest,
) (stream *ErnireBotChatCompletionStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernieBotURL), withBody(request))
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
