package main

func Add(repo string, args ...string) error {
	args = append([]string{"add"}, args...)

	cmd := CreateCommandWithOuts("git", args...)
	cmd.Dir = ResolveRepoPath(repo)

	return cmd.Run()

}
