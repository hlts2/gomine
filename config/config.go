package config

type Config struct {
	Settings `yaml:"settings"`
}

type Settings struct {
	URL    string `yaml:"url"`
	APIKEY string `yaml:"apikey"`
}

const SettingsPath = "~/.config/gomine/config.yaml"
