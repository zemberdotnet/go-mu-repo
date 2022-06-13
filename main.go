package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var helpFlag = flag.Bool("help", false, "print help message")

func main() {
	flag.Parse()
	if *helpFlag {
		PrintUsage()
		os.Exit(0)
	}

	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	if len(os.Args) < 2 {
		// should provide command
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "clone":
		// TODO check the args len
		RunParallel(Clone, []string{config.Prefix + os.Args[2]}, os.Args[3:]...)
	case "checkout":
		// TODO make a function to handle running parallel/serial and args
		if len(os.Args) < 3 {
			RunParallel(Checkout, config.ActiveGroup())
		} else {
			RunParallel(Checkout, config.ActiveGroup(), os.Args[2:]...)
		}
	case "switch":
		if len(os.Args) < 3 {
			RunParallel(Switch, config.ActiveGroup())
		} else {
			RunParallel(Switch, config.ActiveGroup(), os.Args[2:]...)
		}
	case "pull":
		if len(os.Args) < 3 {
			RunParallel(Pull, config.ActiveGroup())
		} else {
			RunParallel(Pull, config.ActiveGroup(), os.Args[2:]...)
		}
	case "add":
		if len(os.Args) < 3 {
			RunParallel(Add, config.ActiveGroup())
		} else {
			RunParallel(Add, config.ActiveGroup(), os.Args[2:]...)
		}
	case "commit":
		if len(os.Args) < 3 {
			RunParallel(Commit, config.ActiveGroup())
		} else {
			RunParallel(Commit, config.ActiveGroup(), os.Args[2:]...)
		}
	case "push":
		if len(os.Args) < 3 {
			RunParallel(Push, config.ActiveGroup())
		} else {
			RunParallel(Push, config.ActiveGroup(), os.Args[2:]...)
		}
	case "status":
		RunParallel(Status, config.ActiveGroup())

	case "register":
		config.Register(os.Args[2])

	case "unregister":
		config.Unregister(os.Args[2])
	case "group":
		config.SetGroup(os.Args[2])
	case "prefix":
		fullPath := ""
		for _, arg := range os.Args[2:] {
			fullPath += arg
		}
		SetPrefix(config, fullPath)
	case "sh":
		RunParallel(Sh, config.ActiveGroup(), os.Args[2:]...)
	case "make":
		RunParallel(config.Make, config.ActiveGroup())

	default:
		args := ""
		for _, arg := range os.Args[1:] {
			args += arg
		}
		log.Println("Unknown command:", args)
	}

	SaveConfig(config)

}
