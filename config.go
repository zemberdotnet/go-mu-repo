package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Prefix string
	Repos  []string
}

func LoadConfig() (*Config, error) {
	f, err := ioutil.ReadFile(".gum")
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(f, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	f, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	return ioutil.WriteFile(".gum", f, 0644)
}
