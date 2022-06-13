package main

import (
	"os"
	"sync"
)

type RunOptions struct {
	fn       Command
	targets  []string
	args     []string
	parallel bool
}

func Run(r *RunOptions) error {
	if r.parallel {
		return RunParallel(r.fn, r.targets, r.args...)
	} else {
		return RunSingleton(r.fn, r.targets, r.args...)
	}
}

func RunSingleton(fn Command, repos []string, args ...string) error {

	for _, repo := range repos {
		// We should copy the args
		// so we don't modify the original
		argsCpy := make([]string, len(args))
		copy(argsCpy, args)

		c := CommandOptions{
			target: repo,
			args:   argsCpy,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		fn(c)
	}
	return nil
}

func RunParallel(fn Command, repos []string, args ...string) error {

	wg := &sync.WaitGroup{}

	for _, repo := range repos {
		// Variables change inside the for range so we need to copy them
		repocpy := repo
		argsCpy := make([]string, len(args))
		copy(argsCpy, args)
		wg.Add(1)
		go func() {
			// TOOD: error aggregation
			c := CommandOptions{
				target: repocpy,
				args:   argsCpy,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}

			fn(c)
			// TODO: Is defer better?
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}
