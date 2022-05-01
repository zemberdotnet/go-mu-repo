package main

func Commit(repo string, args ...string) error {
	args = append([]string{"commit"}, args...)
	cmd := CreateCommandWithOuts("git", args...)
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()
}
