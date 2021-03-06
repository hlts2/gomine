package redmine

import (
	"context"
	"fmt"
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Identifier  string    `json:"identifier"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type Projects []Project

type ReponseProjects struct {
	Projects   `json:"projects"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

type ResponseProject struct {
	Project `json:"project"`
}

func (c *Client) Projects(ctx context.Context) (*ReponseProjects, error) {
	spath := "/projects.json"

	params := map[string]string{
		"key": c.APIKey,
	}

	resp, err := c.doGet(ctx, spath, params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("Project Not Found")
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status Code Not Success: %d", resp.StatusCode)
	}

	obj := &ReponseProjects{}
	if err := decodeBody(resp, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *Client) Project(ctx context.Context, id string) (*ResponseProject, error) {
	spath := "/projects/" + id + ".json"

	params := map[string]string{
		"key": c.APIKey,
	}

	resp, err := c.doGet(ctx, spath, params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("Project Not Found")
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status Code Not Success: %d", resp.StatusCode)
	}

	obj := &ResponseProject{}
	if err := decodeBody(resp, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
