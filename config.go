package go_ernie

import "net/http"

const (
	baiduBceURL                    = "https://aip.baidubce.com"
	defaultEmptyMessagesLimit uint = 300
)

type ClientConfig struct {
	accessToken        string
	ClientId           string
	ClientSecret       string
	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
	Cache              *Cache
}

func DefaultConfig(accessToken string) ClientConfig {
	return ClientConfig{
		accessToken:        accessToken,
		HTTPClient:         &http.Client{},
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
		BaseURL:            baiduBceURL,
		Cache:              NewCache(),
	}
}
