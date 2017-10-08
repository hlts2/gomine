package redmine

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	URL        *url.URL
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(urlStr, apiKey string) (*Client, error) {
	url, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	c := &Client{}
	c.URL = url
	c.APIKey = apiKey
	c.HTTPClient = &http.Client{}
	return c, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, params map[string]string, body io.Reader) (*http.Request, error) {
	url := c.URL
	url.Path = path.Join(url.Path, spath)

	p := url.Query()
	for key, v := range params {
		p.Add(key, v)
	}

	url.RawQuery = p.Encode()

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	return req.WithContext(ctx), nil
}

func (c *Client) doGet(ctx context.Context, spath string, params map[string]string) (*http.Response, error) {
	req, err := c.newRequest(ctx, "GET", spath, params, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(out)
}
