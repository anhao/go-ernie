package go_ernie

import (
	"context"
	"fmt"
	"net/http"
)

const visualGLMURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/txt2img/"

const sdXLURl = "/rpc/2.0/ai_custom/v1/wenxinworkshop/text2image/"

type ImageRequest struct {
	Prompt string `json:"prompt"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Model  string `json:"-"`
}

type ImageResponse struct {
	Images []string `json:"images"`
	APIError
}

func (c *Client) CreateImageVisualGLM(ctx context.Context, request ImageRequest) (response ImageResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", visualGLMURL, request.Model)), withBody(request))
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}

func (c *Client) CreateImageStableDiffusionXL(ctx context.Context, request ImageRequest) (response ImageResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(fmt.Sprintf("%s%s", sdXLURl, request.Model)), withBody(request))
	if err != nil {
		return
	}
	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
