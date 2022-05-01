package main

import (
	"bufio"
	"log"
	"os"
)

type Config struct {
	Repos []string
}

func LoadConfig() (*Config, error) {
	f, err := os.OpenFile(".gum", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		config.Repos = append(config.Repos, scanner.Text())
	}
	return config, nil
}

func SaveConfig(config *Config) error {
	f, err := os.Create(".gum")
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range config.Repos {
		_, err = f.WriteString(repo + "\n")
		if err != nil {
			// TODO Return error
			log.Fatal(err)
		}
	}
	return nil
}
