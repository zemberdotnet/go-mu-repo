package main

import (
	"os"
	"os/exec"
)

func Sh(repo string, args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()
}
