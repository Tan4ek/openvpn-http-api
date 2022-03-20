package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	CAPrivateKeyPass string `yaml:"ca_private_key_pass"`
}

func LoadConfig() Config {
	var AppConfig Config

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Fatal(err)
	}

	return AppConfig
}
