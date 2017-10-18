package redmine

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	c := GetConfig()

	if c.URL == "" {
		t.Error("URL is empty")
	}

	if c.APIKEY == "" {
		t.Error("APIKEY is empty")
	}
}
