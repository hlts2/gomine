package redmine

import (
	"context"
	"testing"
)

func TestProjects(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	err = c.Projects(context.Background())
	if err != nil {
		t.Errorf("Projects error %v", err)
		return
	}
}

func TestProject(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	err = c.Project(context.Background(), 63)
	if err != nil {
		t.Errorf("Project error %v", err)
		return
	}
}
