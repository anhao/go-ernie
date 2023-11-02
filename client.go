package go_ernie

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	utils "github.com/anhao/go-ernie/internal"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	config ClientConfig

	requestBuilder    utils.RequestBuilder
	createFormBuilder func(io.Writer) utils.FormBuilder
}
type AccessTokenResponse struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewClient(accessToken string) *Client {
	config := DefaultConfig(accessToken)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: utils.NewRequestBuilder(),
		createFormBuilder: func(body io.Writer) utils.FormBuilder {
			return utils.NewFormBuilder(body)
		},
	}
}

func NewDefaultClient(clientId, clientSecret string) *Client {
	config := ClientConfig{
		ClientId:           clientId,
		ClientSecret:       clientSecret,
		HTTPClient:         &http.Client{},
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
		BaseURL:            baiduBceURL,
		Cache:              NewCache(),
	}

	return NewClientWithConfig(config)
}

type requestOptions struct {
	body   any
	header http.Header
}

type requestOption func(*requestOptions)

func withBody(body any) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func withContentType(contentType string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("Content-Type", contentType)
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, setters ...requestOption) (*http.Request, error) {
	// Default Options
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}
	for _, setter := range setters {
		setter(args)
	}
	req, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
	if err != nil {
		return nil, err
	}
	if len(c.config.accessToken) == 0 {
		accessToken, err := c.GetAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		c.config.accessToken = *accessToken
	}

	c.setCommonQuery(req)
	return req, nil
}

func (c *Client) sendRequest(req *http.Request, v any) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if isFailureStatusCode(res) {
		return c.handleErrorResp(res)
	}

	return decodeResponse(res.Body, v)
}

func sendRequestStream[T streamable](client *Client, req *http.Request) (*streamReader[T], error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.config.HTTPClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return new(streamReader[T]), err
	}
	if isFailureStatusCode(resp) {
		return new(streamReader[T]), client.handleErrorResp(resp)
	}
	return &streamReader[T]{
		emptyMessagesLimit: client.config.EmptyMessagesLimit,
		reader:             bufio.NewReader(resp.Body),
		response:           resp,
		errAccumulator:     utils.NewErrorAccumulator(),
		unmarshaler:        &utils.JSONUnmarshaler{},
	}, nil
}

func (c *Client) setCommonQuery(req *http.Request) {
	params := url.Values{}
	params.Set("access_token", c.config.accessToken)
	queryString := params.Encode()

	apiUrl := req.URL
	if apiUrl.RawQuery != "" {
		apiUrl.RawQuery += "&" + queryString
	} else {
		apiUrl.RawQuery = queryString
	}
}

func (c *Client) GetAccessToken(ctx context.Context) (*string, error) {
	//判断是否有缓存
	cacheAccessToken, ok := c.config.Cache.Get("cache_" + c.config.ClientId)
	if ok {
		token := cacheAccessToken.(string)
		return &token, nil
	}
	apiUrl := "https://aip.baidubce.com/oauth/2.0/token?client_id=" + c.config.ClientId + "&client_secret=" + c.config.ClientSecret + "&grant_type=client_credentials"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rep AccessTokenResponse
	unmarshaler := utils.JSONUnmarshaler{}
	err = unmarshaler.Unmarshal(body, &rep)
	if err != nil {
		return nil, err
	}
	if rep.Error != "" || rep.AccessToken == "" {
		return nil, errors.New(rep.ErrorDescription)
	}
	//提前100s过期
	c.config.Cache.Set("cache_"+c.config.ClientId, rep.AccessToken, time.Duration(rep.ExpiresIn-100)*time.Second)

	return &rep.AccessToken, nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return decodeString(body, result)
	}
	return json.NewDecoder(body).Decode(v)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes APIError

	err := decodeResponse(resp.Body, errRes)
	if err != nil {
		return err
	}

	return err
}
