package main

import (
	"os"
	"os/exec"
)

func Commit(repo string, args ...string) error {
	args = append([]string{"commit"}, args...)
	cmd := exec.Command("git", args...)
	cmd.Dir = ResolveRepoPath(repo)
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
