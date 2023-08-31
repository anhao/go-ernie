package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

type BaiduChatCompletionStream struct {
	*streamReader[BaiduChatResponse]
}

func (c *Client) CreateBaiduChatCompletionStream(
	ctx context.Context,
	request BaiduChatRequest,
) (stream *BaiduChatCompletionStream, err error) {

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", baiduChatURL, request.Model)), withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendRequestStream[BaiduChatResponse](c, req)
	if err != nil {
		return
	}
	stream = &BaiduChatCompletionStream{
		streamReader: resp,
	}
	return
}
