package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

type CompletionResponseStream struct {
	*streamReader[CompletionResponse]
}

func (c *Client) CreateCompletionStream(
	ctx context.Context,
	request CompletionRequest,
) (stream *CompletionResponseStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", completionURL, request.Model)), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[CompletionResponse](c, req)
	if err != nil {
		return
	}
	stream = &CompletionResponseStream{
		streamReader: resp,
	}
	return
}
