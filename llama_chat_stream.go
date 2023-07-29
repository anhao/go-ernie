package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

type LlamaChatCompletionStream struct {
	*streamReader[LlamaChatResponse]
}

func (c *Client) CreateLlamaChatCompletionStream(
	ctx context.Context,
	request LlamaChatRequest,
) (stream *LlamaChatCompletionStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", llamaChatURL, request.Model)), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[LlamaChatResponse](c, req)
	if err != nil {
		return
	}
	stream = &LlamaChatCompletionStream{
		streamReader: resp,
	}
	return
}
