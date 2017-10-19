package redmine

import (
	"context"
	"testing"
)

func TestMemberShipsByID(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
		return
	}

	obj, err := c.MembershipsByID(context.Background(), "37")
	if err != nil {
		t.Errorf("MembershipsByID error %v", err)
		return
	}

	showMemberships(obj.Memberships)
}

func TestMemberShipsByName(t *testing.T) {
	c, err := NewClient(conf.URL, conf.APIKEY)
	if err != nil {
		t.Errorf("NewClient error %v", err)
	}

	obj, err := c.MembershipsByName(context.Background(), "irai")
	if err != nil {
		t.Errorf("MembershipsByName error %v", err)
		return
	}

	showMemberships(obj.Memberships)
}
