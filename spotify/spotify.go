package spotify

import (
	"context"
	"fmt"
	"net/http"
	"time"
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
		timeOut time.Duration
	}
)

func New(baseURL, clientId, clientSecret string, client *http.Client, timeOut time.Duration) customClient {
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

func (c customClient) Fetch(ctx context.Context) ([]string, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, c.timeOut)
	defer cancelCtx()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/me/top/artist", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)

	return nil, nil
}
