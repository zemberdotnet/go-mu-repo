package main

import (
	"fmt"
	"log"
	"os"
)

var JsonFlag = false
var DebugFlag = false

func main() {

	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmdName, args, err := ParseInput()
	if err != nil {
		PrintUsage()
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		err = SaveConfig(config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// should provide command
		return
	}

	cmd, err := ResolveCommand(cmdName, config)
	if err != nil {
		log.Fatal(err)
	}

	if DebugFlag {
		fmt.Printf("[Command: %v]\n[Args: %v]\n", cmdName, args)
	}

	runOpts := &RunOptions{
		parallel: true,
		fn:       cmd,
		args:     args,
	}

	// The command either targets the active group or a cli specified target
	if CommandHasCLIBasedTarget(cmdName) {
		if cmdName == "clone" {
			for _, arg := range args {
				runOpts.targets = append(runOpts.targets, config.Prefix+arg)
			}
		} else {
			runOpts.targets = args
		}
	} else {
		runOpts.targets = config.ActiveGroup()
	}

	if cmdName == "list" {
		runOpts.once = true
	}

	Run(runOpts)
	SaveConfig(config)

}
