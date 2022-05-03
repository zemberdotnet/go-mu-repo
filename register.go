package main

// Register registers the given repo in the config file
func Register(repos *[]string, repo string) {
	for _, r := range *repos {
		if r == repo {
			return
		}
	}

	*repos = append(*repos, repo)
}

func Unregister(repos *[]string, repo string) {
	unregistered := []string{}
	for _, r := range *repos {
		if r != repo {
			unregistered = append(unregistered, r)
		}
	}
	*repos = unregistered
}

func UnregisterAll(repos *[]string) {
	*repos = []string{}
}
