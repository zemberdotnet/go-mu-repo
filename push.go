package main

import (
	"os"
	"os/exec"
)

func Push(repo string, args ...string) error {
	args = append([]string{"push"}, args...)
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Dir = ResolveRepoPath(repo)
	return cmd.Run()
}
