package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Prefix string
	Repos  []string
}

func LoadConfig() (*Config, error) {
	// if we don't have config file, return empty config
	if _, err := os.Stat(".gum"); errors.Is(err, os.ErrNotExist) {
		return &Config{}, nil
	}

	// if we do have config file, load it
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
