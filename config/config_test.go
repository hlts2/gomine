package config

import (
	"io/ioutil"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

var conf Config

func TestLoadYAML(t *testing.T) {
	buf, err := ioutil.ReadFile(SettingsPath)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
