package redmine

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := NewClient("http:/localhost/", "12345")

	if err != nil {
		t.Errorf("NewClient() error = %v", err)
		return
	}

	if c == nil {
		t.Errorf("NewClient() is nil")
	}

	if c.URL == nil {
		t.Errorf("URL is nil")
	}

	if c.APIKey == "" {
		t.Errorf("URL is nil")
	}

	if c.HTTPClient == nil {
		t.Errorf("HTTPClient is nil")
	}
}

func TestNewRequest(t *testing.T) {
	c, err := NewClient("http:/localhost/", "12345")

	if err != nil {
		t.Errorf("NewClient error = %v", err)
		return
	}

	req, err := c.newRequest(context.Background(), "GET", "hoge.html", nil, nil)
	if err != nil {
		t.Errorf("newRequest() error = %v", err)
		return
	}

	if req == nil {
		t.Errorf("newRequest() is nil")
	}
}

func TestDecodeBody(t *testing.T) {
}
