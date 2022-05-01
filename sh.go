package main

func Sh(repo string, args ...string) error {
	cmd := CreateCommandWithOuts(args[0], args[1:]...)
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()
}
