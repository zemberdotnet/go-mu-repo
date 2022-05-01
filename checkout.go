package main

import (
	"os"
	"os/exec"
)

func Checkout(repo string, args ...string) error {
	args = append([]string{"checkout"}, args...)
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()
}
