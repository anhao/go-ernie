package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

const baiduTxt2ImgURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/txt2img/"

type BaiduTxt2ImgRequest struct {
	Prompt string `json:"prompt"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Model  string `json:"-"`
}

type BaiduTxt2ImgResponse struct {
	Images []string `json:"images"`
	APIError
}

func (c *Client) CreateBaiduTxt2Img(
	ctx context.Context,
	request BaiduTxt2ImgRequest,
) (response BaiduTxt2ImgResponse, err error) {

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", baiduTxt2ImgURL, request.Model)), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
