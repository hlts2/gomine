package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	URL    string `yaml:"url"`
	APIKEY string `yaml:"apikey"`
}

func GetConfig() Config {
	file := "config.yaml"

	if runtime.GOOS == "windows" {
		file = filepath.Join(os.Getenv("APPDATA"), ".config", "gomine", file)
	} else {
		file = filepath.Join(os.Getenv("HOME"), ".config", "gomine", file)
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return c
}
