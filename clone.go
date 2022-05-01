package main

// Clone runs the git clone command for the given repo
func Clone(repo string, args ...string) error {
	args = append([]string{"clone", repo}, args...)
	cmd := CreateCommandWithOuts("git", args...)

	return cmd.Run()
}
