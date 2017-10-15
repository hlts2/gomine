package redmine

import (
	"context"
	"fmt"
	"strconv"
)

type Membership struct {
	ID      int `json:"id"`
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"project"`
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Roles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"roles"`
}

type Memberships []Membership

type ResponseMembership struct {
	Memberships `json:"memberships"`
	TotalCount  int `json:"total_count"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
}

func (c *Client) memberships(ctx context.Context, spath string, params map[string]string) error {
	resp, err := c.doGet(ctx, spath, params)
	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return fmt.Errorf("Memberships Not Found")
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Status Code Not Success: %d", resp.StatusCode)
	}

	obj := &ResponseMembership{}
	if err := decodeBody(resp, obj); err != nil {
		return err
	}

	//cmdパッケージ内で書く
	showMemberships(obj.Memberships)

	return nil
}

func (c *Client) MembershipsByID(ctx context.Context, id int) error {
	spath := "/projects/" + strconv.Itoa(id) + "/memberships.json"

	params := map[string]string{
		"key": c.APIKey,
	}

	return c.memberships(ctx, spath, params)
}

func (c *Client) MembershipsByName(ctx context.Context, name string) error {
	spath := "/projects/" + name + "/memberships.json"

	params := map[string]string{
		"key": c.APIKey,
	}

	return c.memberships(ctx, spath, params)
}
