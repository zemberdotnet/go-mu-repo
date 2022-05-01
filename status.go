package main

import (
	"os"
	"os/exec"
)

func Status(repo string, args ...string) error {
	args = append([]string{"status"}, args...)
	cmd := exec.Command("git", args...)
	cmd.Dir = ResolveRepoPath(repo)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
