package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

type ErnieCustomPluginStream struct {
	*streamReader[ErnieCustomPluginResponse]
}

func (c *Client) CreateErnieCustomPluginStreamChatCompletion(
	ctx context.Context,
	request ErnieCustomPluginRequest,
) (stream *ErnieCustomPluginStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s/", ernireCustomPluginUrl, request.PluginName)), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[ErnieCustomPluginResponse](c, req)
	if err != nil {
		return
	}
	stream = &ErnieCustomPluginStream{
		streamReader: resp,
	}
	return
}
