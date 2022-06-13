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
		SaveConfig(config)
	}

	if len(os.Args) < 2 {
		// should provide command
		return
	}

	cmdName := os.Args[1]

	cmd, err := ResolveCommand(cmdName, config)
	if err != nil {
		log.Fatal(err)
	}

	var args = []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	runOpts := &RunOptions{
		parallel: true,
		fn:       cmd,
		args:     args,
	}

	// The command either targets the active group or a cli specified target
	if CommandHasCLIBasedTarget(cmdName) {
		if cmdName == "clone" {
			runOpts.targets = []string{config.Prefix + os.Args[2]}
		} else {
			runOpts.targets = []string{os.Args[2]}
		}

	} else {
		runOpts.targets = config.ActiveGroup()
	}

	Run(runOpts)
	SaveConfig(config)

}
