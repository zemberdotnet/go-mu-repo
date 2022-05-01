package main

import (
	"log"
	"os"
	"os/exec"
)

func ResolveRepoPath(repo string) string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd + "/" + repo
}

func CreateCommandWithOuts(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd
}
