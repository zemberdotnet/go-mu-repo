package main

import (
	"os"
)

func (c *Config) Make(repo string, args ...string) error {
	// test if the repo exists
	if _, err := os.Stat(repo); os.IsNotExist(err) {
		Clone(c.Prefix + repo)
	}
	return nil
}
