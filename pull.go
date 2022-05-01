package main

import (
	"os"
	"os/exec"
)

// Pull runs the git pull command for the given repo
func Pull(repo string, args ...string) error {

	cmd := exec.Command("git", "pull")
	cmd.Dir = ResolveRepoPath(repo)
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
