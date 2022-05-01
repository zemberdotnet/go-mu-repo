package main

import (
	"log"
	"os"
)

func ResolveRepoPath(repo string) string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd + "/" + repo

}
