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
		fmt.Println("No config file found, generating default config")
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

func (c *Config) SetGroup(cmdOpts CommandOptions) error {
	c.CurrentGroup = cmdOpts.target
	return nil
}

func (c *Config) SetPrefix(cmdOpts CommandOptions) error {
	c.Prefix = cmdOpts.target
	return nil
}

func (c *Config) ActiveGroup() []string {
	return c.Groups[c.CurrentGroup]
}

func (c *Config) Register(cmdOpts CommandOptions) error {
	for _, r := range c.ActiveGroup() {
		if r == cmdOpts.target {
			return nil
		}
	}
	c.Groups[c.CurrentGroup] = append(c.Groups[c.CurrentGroup], cmdOpts.target)
	return nil
}

func (c *Config) Unregister(cmdOpts CommandOptions) error {
	unregistered := []string{}
	for _, r := range c.ActiveGroup() {
		if r != cmdOpts.target {
			unregistered = append(unregistered, r)
		}
	}
	c.Groups[c.CurrentGroup] = unregistered
	return nil
}

func (c *Config) List(cmdOpts CommandOptions) error {
	for _, r := range c.ActiveGroup() {
		fmt.Println(r)
	}
	return nil
}

func (c *Config) Make(cmdOpts CommandOptions) error {
	if _, err := os.Stat(cmdOpts.target); os.IsNotExist(err) {
		Clone(CommandOptions{
			target: c.Prefix + cmdOpts.target,
			args:   cmdOpts.args,
			Stdout: cmdOpts.Stdout,
			Stderr: cmdOpts.Stderr,
		})
	}
	return nil
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
