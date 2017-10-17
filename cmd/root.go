package cmd

import (
	"fmt"
	"os"

	c "github.com/hlts2/gomine/config"
	cli "github.com/spf13/cobra"
)

var (
	conf c.Config

	RootCmd = &cli.Command{
		Use:   "gomine",
		Short: "A CLI tool for Redmine tickets viewer",
	}
)

func Execute() {
	conf = c.GetConfig()
	if conf.URL == "" || conf.APIKEY == "" {
		fmt.Println("URL or APIKEY is not an exact value")
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
