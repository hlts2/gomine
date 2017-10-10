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
	return nil
}
