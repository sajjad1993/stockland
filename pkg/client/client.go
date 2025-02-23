package client

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Get(ctx context.Context, url string, customHeader map[string]string) (*http.Response, error)
	Post(ctx context.Context, url string, body []byte, customHeader map[string]string) (*http.Response, error)
	PostForm(ctx context.Context, url string, reader io.Reader, customHeader map[string]string) (*http.Response, error)
}
