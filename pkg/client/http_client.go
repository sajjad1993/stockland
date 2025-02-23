package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

func NewClient() Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &HTTPClient{
		client: &http.Client{Transport: tr},
	}
}
func (c *HTTPClient) Post(ctx context.Context, url string, body []byte, customHeader map[string]string) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range customHeader {
		req.Header.Add(k, v)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	//todo return status and io.reader
	//defer response.Body.Close()
	return response, nil
}
func (c *HTTPClient) Get(ctx context.Context, url string, customHeader map[string]string) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range customHeader {
		req.Header.Add(k, v)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *HTTPClient) PostForm(ctx context.Context, url string, reader io.Reader, customHeader map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range customHeader {
		req.Header.Add(k, v)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	//todo return status and io.reader
	//defer response.Body.Close()
	return response, nil
}
