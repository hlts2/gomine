package redmine

import (
	"context"
	"testing"
)

const (
	URL     = ""
	API_KEY = ""
)

func TestIssues(t *testing.T) {
	c, err := NewClient(URL, API_KEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	err = c.Issues(context.Background())
	if err != nil {
		t.Errorf("Issues error %v", err)
		return
	}
}

func TestIssue(t *testing.T) {
	c, err := NewClient(URL, API_KEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	err = c.Issue(context.Background(), 10078)
	if err != nil {
		t.Errorf("Issue error %v", err)
		return
	}
}
