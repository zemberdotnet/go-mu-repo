package main

func Switch(repo string, args ...string) error {
    args = append([]string{"switch"}, args...)
    cmd := CreateCommandWithOuts("git", args...)
    cmd.Dir = ResolveRepoPath(repo)
    return cmd.Run()
}

