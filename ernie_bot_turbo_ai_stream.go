package go_ernie

import (
	"context"
	"net/http"
)

type ErnireBotTurboAIChatCompletionStream struct {
	*streamReader[ErnieBotTurboAIResponse]
}

func (c *Client) CreateErnieBotTurboAIChatCompletionStream(
	ctx context.Context,
	request ErnieBotTurboAIRequest,
) (stream *ErnireBotTurboAIChatCompletionStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernireBotTurboAIURL), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[ErnieBotTurboAIResponse](c, req)
	if err != nil {
		return
	}
	stream = &ErnireBotTurboAIChatCompletionStream{
		streamReader: resp,
	}
	return
}
