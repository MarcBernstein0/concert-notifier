package spotify

import (
	"context"
	"net/http"
)

type (
	FetchData interface {
		Fetch(ctx context.Context) ([]string, error)
	}

	customClient struct {
		baseURL string
		client  *http.Client
		config  struct {
			clientId     string
			clientSecret string
		}
		timeOut int
	}
)

func New(baseURL, clientId, clientSecret string, client *http.Client, timeOut int) customClient {
	return customClient{
		baseURL: baseURL,
		client:  client,
		config: struct {
			clientId     string
			clientSecret string
		}{
			clientId:     clientId,
			clientSecret: clientSecret,
		},
		timeOut: timeOut,
	}
}
