package main

import (
	"fmt"
	"io"
)

type Command func(c CommandOptions) error

type CommandOptions struct {
	target string // Often the repo name, but sometimes the group or prefix
	args   []string
	Stdout io.Writer
	Stderr io.Writer
}

var CommandNames = []string{
	"clone",
	"commit",
	"pull",
	"push",
	"checkout",
	"add",
	"status",
	"switch",
	"sh",
	"stash",
	"register",
	"unregister",
	"make",
	"group",
	"prefix",
	"list",
	"reset",
}

var NameCommandMap = map[string]Command{
	"clone":    Clone,
	"commit":   Commit,
	"pull":     Pull,
	"push":     Push,
	"checkout": Checkout,
	"add":      Add,
	"status":   Status,
	"switch":   Switch,
	"sh":       Sh,
	"stash":    Stash,
	"reset":    Reset,
}

var ErrUnknownCommand = fmt.Errorf("unknown command")

// CommandHasClIBasedTarget checks if the command expects a target from the CLI
// as opposed to targets from the active group
func CommandHasCLIBasedTarget(cmd string) bool {
	return cmd == "clone" || cmd == "register" || cmd == "unregister" || cmd == "group" || cmd == "prefix"
}

// ResolveCommand resolves the command name to a command function
// or returns ErrUnknownCommand if the command is not found
func ResolveCommand(cmd string, config *Config) (Command, error) {
	if c, ok := NameCommandMap[cmd]; ok {
		return c, nil
	}

	switch cmd {
	case "register":
		return config.Register, nil
	case "unregister":
		return config.Unregister, nil
	case "list":
		return config.List, nil
	case "make":
		return config.Make, nil
	case "group":
		return config.SetGroup, nil
	case "prefix":
		return config.SetPrefix, nil
	}
	return nil, ErrUnknownCommand

}

// Clone runs the git clone command for the given repo
func Clone(c CommandOptions) error {
	args := append([]string{"clone", c.target}, c.args...)
	cmd := CreateCommand("git", args...)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Commit(c CommandOptions) error {
	args := append([]string{"commit"}, c.args...)
	cmd := CreateCommand("git", args...)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)
	cmd.Dir = ResolveRepoPath(c.target)

	return cmd.Run()
}

// Pull runs the git pull command for the given repo
func Pull(c CommandOptions) error {
	args := append([]string{"pull"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Push(c CommandOptions) error {
	args := append([]string{"push"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Checkout(c CommandOptions) error {
	args := append([]string{"checkout"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Add(c CommandOptions) error {
	args := append([]string{"add"}, c.args...)

	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()

}

func Status(c CommandOptions) error {
	args := append([]string{"status"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Switch(c CommandOptions) error {
	args := append([]string{"switch"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Stash(c CommandOptions) error {
	args := append([]string{"stash"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}

func Reset(c CommandOptions) error {
	args := append([]string{"reset"}, c.args...)
	cmd := CreateCommand("git", args...)
	cmd.Dir = ResolveRepoPath(c.target)

	return cmd.Run()
}

func Sh(c CommandOptions) error {
	cmd := CreateCommand(c.args[0], c.args[1:]...)
	cmd.Dir = ResolveRepoPath(c.target)
	AddOutsToCommand(cmd, c.Stdout, c.Stderr)

	return cmd.Run()
}
