package cmd

import (
	"context"
	"fmt"

	redmine "github.com/hlts2/gomine/redmine"
	cli "github.com/spf13/cobra"
)

var catCmd = &cli.Command{
	Use:   "cat",
	Short: "Show ticket details of Redmine",
	Run: func(cmd *cli.Command, args []string) {
		if err := cat(cmd, args); err != nil {
			fmt.Println(err)
		}
	},
}

var (
	tcktNum string
)

func init() {
	RootCmd.AddCommand(catCmd)

	catCmd.PersistentFlags().StringVarP(&tcktNum, "number", "n", "", "get ticket details of Redmine")
}

func cat(cmd *cli.Command, args []string) error {
	switch args[0] {

	//Issues
	case "i":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
			//TODO error
		}

		obj, err := c.Issue(context.Background(), tcktNum)
		if err != nil {
			//TODO error
		}

		redmine.ShowDetIssue(obj.Issue)

	//Projects
	case "p":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
		}

		obj, err := c.Project(context.Background(), tcktNum)
		if err != nil {
		}

		redmine.ShowDetProject(obj.Project)

	//Memberships
	case "m":
		c, err := redmine.NewClient(conf.URL, conf.APIKEY)
		if err != nil {
		}

		obj, err := c.MembershipsByID(context.Background(), tcktNum)
		if err != nil {
		}

		redmine.ShowMemberships(obj.Memberships)

	default:
		return fmt.Errorf("gomine %s cat: It is a noexistent command", args[0])
	}

	return nil
}
