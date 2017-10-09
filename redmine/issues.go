package redmine

import (
	"context"
	"strconv"
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

func (c *Client) Issues(ctx context.Context) error {
	spath := "/issues.json"

	params := map[string]string{
		"key": c.APIKey,
	}

	resp, err := c.doGet(ctx, spath, params)
	if err != nil {
		return err
	}

	obj := &ResponseIssues{}
	if err := decodeBody(resp, obj); err != nil {
		return err
	}

	//cmdパッケーで書く
	showIssues(obj.Issues)

	return nil
}

func (c *Client) Issue(ctx context.Context, id int) error {
	spath := "/issues/" + strconv.Itoa(id) + ".json"

	params := map[string]string{
		"key": c.APIKey,
	}

	resp, err := c.doGet(ctx, spath, params)
	if err != nil {
		return err
	}

	obj := &Issue{}
	if err := decodeBody(resp, obj); err != nil {
		return nil
	}

	//cmdパッケーで書く
	showDetIssue(*obj)

	return nil
}
