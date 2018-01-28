package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/hlts2/gomine/redmine"
	cli "github.com/spf13/cobra"
)

var lsCmd = &cli.Command{
	Use:   "ls",
	Short: "List \"issues\", \"projects\", \"memberships\" of Redmine",
	Run: func(cmd *cli.Command, args []string) {
		if err := ls(cmd, args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	filters []string
)

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.PersistentFlags().StringArrayVarP(&filters, "filter", "f", []string{}, "get")
}

func ls(cmd *cli.Command, args []string) error {
	if len(args) == 0 {
		return lsUsage()
	}

	switch args[0] {

	//Issues
	case "i":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
			return err
		}

		obj, err := c.Issues(context.Background())
		if err != nil {
			return err
		}

		if len(filters) != 0 {
			obj.Issues = redmine.FilterIssues(obj.Issues, filters)
		}

		redmine.ShowIssues(obj.Issues)

	//Projects
	case "p":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
			return err
		}

		obj, err := c.Projects(context.Background())
		if err != nil {
			return err
		}

		if len(filters) != 0 {
			obj.Projects = redmine.FilterProjects(obj.Projects, filters)
		}

		redmine.ShowProjects(obj.Projects)

	default:
		return lsUsage()
	}

	return nil
}

func lsUsage() error {
	return errors.New(`gomine ls <option> <arguments>

Issues Command:
  ls    listing issues
        $ gomine ls i

        filtered listing issues
        $ gomine ls -f <filter word> i

Projects Command:
  ls    listing projects
        $ gomine ls p

        filtered listing projects
        $ gomine ls -f <filter word> p

Flags:
  -f filter the list
	`)
}
