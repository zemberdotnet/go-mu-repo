package main

import (
	"fmt"
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

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("gum prefix <url>")
	fmt.Println("gum clone <repo>")
	fmt.Println("gum checkout <branch>")
	fmt.Println("gum switch <branch> [<branch>...]")
	fmt.Println("gum pull <branch> [<branch>...]")
	fmt.Println("gum commit [<file>...]")
	fmt.Println("gum status")
	fmt.Println("gum register <repo-path>")
	fmt.Println("gum unregister <repo-path>")
	fmt.Println("gum make")
}
