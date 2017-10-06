package redmine

import (
	"context"
	"testing"
)

func TestIssues(t *testing.T) {

	t.Run("Issues", func(t *testing.T) {

		c, err := NewClient("", "")
		if err != nil {
			t.Errorf("NewClient error %v", err)
			return
		}

		err = c.Issues(context.Background())
		if err != nil {
			t.Errorf("Issues error %v", err)
			return
		}
	})

}
