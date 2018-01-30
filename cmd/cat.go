package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"

	redmine "github.com/hlts2/gomine/redmine"
	cli "github.com/spf13/cobra"
)

var catCmd = &cli.Command{
	Use:   "cat",
	Short: "Show \"issues\", \"projects\", \"memberships\" details of Redmine",
	Run: func(cmd *cli.Command, args []string) {
		if err := cat(cmd, args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	id string
)

func init() {
	rootCmd.AddCommand(catCmd)

	catCmd.PersistentFlags().StringVarP(&id, "number", "n", "", "get ticket details of Redmine")
}

func cat(cmd *cli.Command, args []string) error {
	if len(args) == 0 {
		return catUsage()
	}

	progress := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	progress.Start()
	defer progress.Stop()

	switch args[0] {

	//Issues
	case "i":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
			return err
		}

		obj, err := c.Issue(context.Background(), id)
		if err != nil {
			return err
		}

		redmine.ShowDetIssue(obj.Issue)

	//Projects
	case "p":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
			return err
		}

		obj, err := c.Project(context.Background(), id)
		if err != nil {
			return err
		}

		redmine.ShowDetProject(obj.Project)

	//Memberships
	case "m":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
			return err
		}

		obj, err := c.MembershipsByID(context.Background(), id)
		if err != nil {
			return err
		}

		redmine.ShowMemberships(obj.Memberships)

	default:
		return catUsage()
	}

	return nil
}

func catUsage() error {
	return errors.New(`gomine cat <option> <arguments>

Issues Command:
  cat   show given issue
        $ gomine cat -n 10078 i

Projects Command:
  cat   show given project
        $ gomine cat -n 1 p

Memberships Command:
  cat   show memberships of given project
        $ gomine cat -n 1 m
	`)
}
