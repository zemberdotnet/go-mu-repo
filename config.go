package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Prefix       string
	CurrentGroup string
	Groups       map[string][]string
}

var DefaultConfig = Config{
	Prefix:       "",
	CurrentGroup: "default",
	Groups: map[string][]string{
		"default": {},
	},
}

func LoadConfig() (*Config, error) {
	// if we don't have config file, return empty config
	if _, err := os.Stat(".gum"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No config file found, using default config")
		return &DefaultConfig, nil
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

func (c *Config) SetGroup(group string) {
	c.CurrentGroup = group
}

func (c *Config) SetPrefix(prefix string) {
	c.Prefix = prefix
}

func (c *Config) ActiveGroup() []string {
	return c.Groups[c.CurrentGroup]
}

func (c *Config) Register(repo string) {
	for _, r := range c.ActiveGroup() {
		if r == repo {
			return
		}
	}
	c.Groups[c.CurrentGroup] = append(c.Groups[c.CurrentGroup], repo)
}

func (c *Config) Unregister(repo string) {
	unregistered := []string{}
	for _, r := range c.ActiveGroup() {
		if r != repo {
			unregistered = append(unregistered, r)
		}
	}
	c.Groups[c.CurrentGroup] = unregistered
}

func (c *Config) UnregisterAll() {
	c.Groups[c.CurrentGroup] = []string{}
}

func SaveConfig(config *Config) error {
	f, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	return ioutil.WriteFile(".gum", f, 0644)
}
