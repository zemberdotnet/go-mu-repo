package main

import (
	"sync"
)

type Command func(string, ...string) error

func RunParallel(fn Command, repos []string, args ...string) error {

	wg := &sync.WaitGroup{}

	for _, repo := range repos {
		repocpy := repo
		wg.Add(1)
		go func() {
			// TOOD: error aggregation
			fn(repocpy, args...)
			// TODO: Is defer better?
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}
