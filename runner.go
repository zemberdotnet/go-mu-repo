package main

import (
	"os"
	"os/exec"
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

	writers := make([]*OutputWriter, 0, len(repos))
	for _, repo := range repos {
		// We should copy the args
		// so we don't modify the original
		argsCpy := make([]string, len(args))
		copy(argsCpy, args)

		writer := NewOuputWriter(repo)
		writers = append(writers, writer)

		c := CommandOptions{
			target: repo,
			args:   argsCpy,
			Stdout: writer,
			Stderr: os.Stderr,
		}

		err := fn(c)
		if err != nil {
			switch err := err.(type) {
			case *exec.ExitError:
				exitCode := err.ExitCode()
				writer.exitCode = exitCode
			case *exec.Error:
				// TODO: Handle execErrors
			}
		}

	}

	for _, writer := range writers {
		writer.Flush()
	}

	return nil
}

func RunParallel(fn Command, repos []string, args ...string) error {

	wg := &sync.WaitGroup{}

	writers := make([]*OutputWriter, 0, len(repos))
	for _, repo := range repos {
		// Variables change inside the for range so we need to copy them
		repocpy := repo
		argsCpy := make([]string, len(args))
		copy(argsCpy, args)
		wg.Add(1)
		go func() {
			// TODO: This is a data race!!!
			writer := NewOuputWriter(repocpy)
			writers = append(writers, writer)

			c := CommandOptions{
				target: repocpy,
				args:   argsCpy,
				Stdout: writer,
				Stderr: os.Stderr,
			}

			err := fn(c)
			// TODO: maybe improve the way exit codes are assigned
			if err != nil {
				switch err := err.(type) {
				case *exec.ExitError:
					exitCode := err.ExitCode()
					writer.exitCode = exitCode
				case *exec.Error:
					// TODO: Handle execErrors
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	for _, writer := range writers {
		writer.Flush()
	}

	return nil
}
