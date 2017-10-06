package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"
)

var RootCmd = &cli.Command{
	Use:   "gomine",
	Short: "A CLI tool for Redmine",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
