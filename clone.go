package main

import (
	"os"
	"os/exec"
)

// Clone runs the git clone command for the given repo
func Clone(repo string, args ...string) error {
	args = append([]string{"clone", repo}, args...)
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
