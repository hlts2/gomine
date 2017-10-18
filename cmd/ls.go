package cmd

import (
	"fmt"

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

	lsCmd.Flags().StringVarP(&toWhom, "to", "t", "me", "get tickets to me")
	lsCmd.Flags().StringVarP(&filter, "filter", "f", "", "get filtered tickets")
}

func ls(cmd *cli.Command, args []string) error {
	switch args[0] {

	//Issues
	case "i":

	//Projects
	case "p":

	//Memberships
	case "m":
	default:
		return fmt.Errorf("gomine ls: %v: It is a noexistent command option", args[0])
	}

	return nil
}
