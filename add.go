package main

import (
	"os"
	"os/exec"
)

func Add(repo string, args ...string) error {
	args = append([]string{"add"}, args...)

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()

}
