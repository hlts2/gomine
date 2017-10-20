package cmd

import (
	"errors"
	"fmt"
	"os"

	c "github.com/hlts2/gomine/redmine"
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
	if conf.URL == "" {
		fmt.Println("URL is empty value")
		os.Exit(1)
	}

	if conf.APIKEY == "" {
		fmt.Println("API_KEY is empty value")
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func usage() error {
	return errors.New(`gomine <command> <option> <arguments>

Issues Command:
  ls    listing projects
        $ gomine ls i

  cat   show given issue
        $ gomine cat -n 10078 i

Projects Command:
  ls    listing projects
        $ gomine ls p

  cat   show given project
        $ gomine cat -n 1 p

Memberships Command:
  cat   show memberships of given project
        $ gomine cat -n 1 m
	`)
}
