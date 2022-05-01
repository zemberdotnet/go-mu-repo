package main

func Status(repo string, args ...string) error {
	args = append([]string{"status"}, args...)
	cmd := CreateCommandWithOuts("git", args...)
	cmd.Dir = ResolveRepoPath(repo)
	return cmd.Run()
}
