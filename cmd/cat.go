package cmd

import (
	"fmt"

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

	catCmd.Flags().StringVarP(&tcktNum, "n", "number", "", "get ticket details of Redmine")
}

func cat(cmd *cli.Command, args []string) error {
	return nil
}
