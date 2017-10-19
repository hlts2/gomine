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

	catCmd.PersistentFlags().StringVarP(&tcktNum, "number", "n", "", "get ticket details of Redmine")
}

func cat(cmd *cli.Command, args []string) error {
	switch args[0] {

	//Issues
	case "i":

	//Projects
	case "p":

	//Memberships
	case "m":

	default:
		return fmt.Errorf("gomine %s cat: It is a noexistent command option", args[0])
	}

	return nil
}
