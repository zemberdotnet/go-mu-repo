package main

import (
	"fmt"
	"os"
)

// ParseInput parses the input and returns the command, arguments, or an error
func ParseInput() (string, []string, error) {
	for i, arg := range os.Args[1:] {
		for _, cmd := range CommandNames {
			if arg == cmd {
				cmdName := arg
				args := os.Args[i+2:]
				return cmdName, args, nil
			}
		}
		if arg == "--help" {
			PrintUsage()
			os.Exit(1)
		}

		if arg == "--json" {
			JsonFlag = true
		}

		if arg == "--debug" {
			DebugFlag = true
		}

	}
	return "", nil, fmt.Errorf("no command")
}
