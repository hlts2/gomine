package redmine

import (
	"context"
	"time"
)

type Issue struct {
	ID      int `json:"id"`
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"project"`
	Tracker struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tracker"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	Priority struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"priority"`
	Author struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
	AssignedTo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"assigned_to"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"`
	DueDate     string    `json:"due_date,omitempty"`
	DoneRatio   int       `json:"done_ratio"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type Issues []Issue

type ResponseIssues struct {
	Issues     `json:"issues"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

func (c *Client) issues(ctx context.Context, params map[string]string) (*ResponseIssues, error) {
	spath := "/issues.json"

	req, err := c.newRequest(ctx, "GET", spath, params, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	resIssues := &ResponseIssues{}
	if err := decodeBody(resp, resIssues); err != nil {
		return nil, err
	}

	return resIssues, nil
}

func (c *Client) Issues(ctx context.Context) error {
	params := map[string]string{
		"key": c.APIKey,
	}

	res, err := c.issues(ctx, params)
	if err != nil {
		return err
	}

	viewer(res.Issues)

	return nil
}
