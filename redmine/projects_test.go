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

	obj, err := c.Projects(context.Background())
	if err != nil {
		t.Errorf("Projects error %v", err)
		return
	}

	ShowProjects(obj.Projects)
}

func TestProject(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	obj, err := c.Project(context.Background(), "63")
	if err != nil {
		t.Errorf("Project error %v", err)
		return
	}

	ShowDetProject(obj.Project)
}
