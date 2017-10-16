package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	c "github.com/hlts2/gomine/config"
	"gopkg.in/yaml.v2"

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
	buf, err := ioutil.ReadFile(c.SettingsPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
