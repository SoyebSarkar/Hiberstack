package typesense

import "net/http"

type Client struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func New(url, key string) *Client {
	return &Client{
		BaseURL: url,
		APIKey:  key,
		Client:  &http.Client{},
	}
}
