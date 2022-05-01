package main

// Pull runs the git pull command for the given repo
func Pull(repo string, args ...string) error {
	args = append([]string{"pull"}, args...)
	cmd := CreateCommandWithOuts("git", args...)
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()
}
