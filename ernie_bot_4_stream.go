package go_ernie

import (
	"context"
	"net/http"
)

type ErnireBot4ChatCompletionStream struct {
	*streamReader[ErnieBot4Response]
}

func (c *Client) CreateErnieBot4ChatCompletionStream(
	ctx context.Context,
	request ErnieBot4Request,
) (stream *ErnireBot4ChatCompletionStream, err error) {
	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernieBot4URL), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[ErnieBot4Response](c, req)
	if err != nil {
		return
	}
	stream = &ErnireBot4ChatCompletionStream{
		streamReader: resp,
	}
	return
}
