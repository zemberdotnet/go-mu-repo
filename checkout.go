package main

func Checkout(repo string, args ...string) error {
	args = append([]string{"checkout"}, args...)
	cmd := CreateCommandWithOuts("git", args...)
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()
}
