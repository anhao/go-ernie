package go_ernie

import (
	"context"
	"net/http"
)

type ErnireBot8KChatCompletionStream struct {
	*streamReader[ErnieBot8KResponse]
}

func (c *Client) CreateErnieBot8KChatCompletionStream(
	ctx context.Context,
	request ErnieBot8KRequest,
) (stream *ErnireBot8KChatCompletionStream, err error) {
	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernieBot8KURL), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[ErnieBot8KResponse](c, req)
	if err != nil {
		return
	}
	stream = &ErnireBot8KChatCompletionStream{
		streamReader: resp,
	}
	return
}
