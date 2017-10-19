package cmd

import (
	"context"
	"errors"
	"fmt"

	redmine "github.com/hlts2/gomine/redmine"
	cli "github.com/spf13/cobra"
)

var lsCmd = &cli.Command{
	Use:   "ls",
	Short: "List the Redmine tickets",
	Run: func(cmd *cli.Command, args []string) {
		if err := ls(cmd, args); err != nil {
			fmt.Println(err)
		}
	},
}

var (
	toWhom string
	filter string
)

func init() {
	RootCmd.AddCommand(lsCmd)

	lsCmd.PersistentFlags().StringVarP(&toWhom, "to", "t", "me", "get tickets to me")
	lsCmd.PersistentFlags().StringVarP(&filter, "filter", "f", "", "get filtered tickets")
}

func ls(cmd *cli.Command, args []string) error {
	if len(args) == 0 {
		//TODO error
		return errors.New("")
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

		redmine.ShowProjects(obj.Projects)
	default:
		return fmt.Errorf("gomine ls %s : It is a noexistent command", args[0])
	}

	return nil
}
