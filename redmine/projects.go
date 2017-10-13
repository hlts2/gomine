package redmine

import (
	"context"
	"strconv"
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
	Project `json:project`
}

func (c *Client) Projects(ctx context.Context) error {
	spath := "/projects.json"

	params := map[string]string{
		"key": c.APIKey,
	}

	resp, err := c.doGet(ctx, spath, params)
	if err != nil {
		return err
	}

	obj := &ReponseProjects{}
	if err := decodeBody(resp, obj); err != nil {
		return nil
	}

	//TODO cmd側で書く
	showProjects(obj.Projects)

	return nil
}

func (c *Client) Project(ctx context.Context, id int) error {
	spath := "/projects/" + strconv.Itoa(id) + ".json"

	params := map[string]string{
		"key": c.APIKey,
	}

	res, err := c.doGet(ctx, spath, params)
	if err != nil {
		return err
	}

	obj := &ResponseProject{}
	if err := decodeBody(res, obj); err != nil {
		return nil
	}

	//TODO cmd側で書く
	showDetProject(obj.Project)

	return nil
}
