package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	switch os.Args[1] {
	case "clone":
		// TODO check the args len
		RunParallel(Clone, []string{config.Prefix + os.Args[2]}, os.Args[3:]...)
	case "checkout":
		// TODO make a function to handle running parallel/serial and args
		if len(os.Args) < 3 {
			RunParallel(Checkout, config.Repos)
		} else {
			RunParallel(Checkout, config.Repos, os.Args[2:]...)
		}
	case "pull":
		if len(os.Args) < 3 {
			RunParallel(Pull, config.Repos)
		} else {
			RunParallel(Pull, config.Repos, os.Args[2:]...)
		}
	case "add":
		if len(os.Args) < 3 {
			RunParallel(Add, config.Repos)
		} else {
			RunParallel(Add, config.Repos, os.Args[2:]...)
		}
	case "commit":
		if len(os.Args) < 3 {
			RunParallel(Commit, config.Repos)
		} else {
			RunParallel(Commit, config.Repos, os.Args[2:]...)
		}
	case "push":
		if len(os.Args) < 3 {
			RunParallel(Push, config.Repos)
		} else {
			RunParallel(Push, config.Repos, os.Args[2:]...)
		}
	case "status":
		RunParallel(Status, config.Repos)

	case "register":
		Register(config, os.Args[2])

	case "unregister":
		Unregister(config, os.Args[2])
	case "prefix":
		fullPath := ""
		for _, arg := range os.Args[2:] {
			fullPath += arg
		}
		SetPrefix(config, fullPath)
	case "sh":
		RunParallel(Sh, config.Repos, os.Args[2:]...)

	}

	SaveConfig(config)

}
