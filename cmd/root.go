package cmd

import (
	"fmt"
	"os"

	c "github.com/hlts2/gomine/redmine"
	cli "github.com/spf13/cobra"
)

var (
	conf c.Config

	rootCmd = &cli.Command{
		Use:   "gomine",
		Short: "A CLI tool for Redmine tickets viewer",
	}
)

func Execute() {
	conf = c.GetConfig()
	if conf.URL == "" {
		fmt.Println("URL is empty value")
		os.Exit(1)
	}

	if conf.APIKEY == "" {
		fmt.Println("API_KEY is empty value")
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
