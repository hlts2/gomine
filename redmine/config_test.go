package redmine

import (
	"testing"
)

var conf Config

func TestGetConfig(t *testing.T) {
	conf = GetConfig()

	if conf.URL == "" {
		t.Error("URL is empty")
	}

	if conf.APIKEY == "" {
		t.Error("APIKEY is empty")
	}
}
