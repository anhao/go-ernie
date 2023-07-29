package go_ernie

import (
	"context"
	"net/http"
)

type Bloomz7b1ChatCompletionStream struct {
	*streamReader[Bloomz7b1Response]
}

func (c *Client) CreateBloomz7b1ChatCompletionStream(
	ctx context.Context,
	request Bloomz7b1Request,
) (stream *Bloomz7b1ChatCompletionStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(bloomz7b1URL), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[Bloomz7b1Response](c, req)
	if err != nil {
		return
	}
	stream = &Bloomz7b1ChatCompletionStream{
		streamReader: resp,
	}
	return
}
