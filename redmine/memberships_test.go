package redmine

import (
	"context"
	"testing"
)

func TestMemberShipsByID(t *testing.T) {
	c, err := NewClient(URL, API_KEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	err = c.MembershipsByID(context.Background(), 37)
	if err != nil {
		t.Errorf("MembershipsByID error %v", err)
		return
	}
}

func TestMemberShipsByName(t *testing.T) {
	c, err := NewClient(URL, API_KEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
	}

	err = c.MembershipsByName(context.Background(), "irai")
	if err != nil {
		t.Errorf("MembershipsByName error %v", err)
		return
	}
}
