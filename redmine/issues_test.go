package redmine

import (
	"context"
	"testing"
)

func TestIssues(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	obj, err := c.Issues(context.Background())
	if err != nil {
		t.Errorf("Issues error %v", err)
		return
	}

	showIssues(obj.Issues)
}

func TestIssue(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	obj, err := c.Issue(context.Background(), 10078)
	if err != nil {
		t.Errorf("Issue error %v", err)
		return
	}

	showDetIssue(obj.Issue)
}
