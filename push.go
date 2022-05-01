package main

func Push(repo string, args ...string) error {
	args = append([]string{"push"}, args...)
	cmd := CreateCommandWithOuts("git", args...)
	cmd.Dir = ResolveRepoPath(repo)
	return cmd.Run()
}
