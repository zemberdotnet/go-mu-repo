package main

// Register registers the given repo in the config file
func Register(config *Config, repo string) {
	for _, r := range config.Repos {
		if r == repo {
			return
		}
	}

	config.Repos = append(config.Repos, repo)
}

func Unregister(config *Config, repo string) {
	updated := []string{}
	for _, r := range config.Repos {
		if r != repo {
			updated = append(updated, r)
		}
	}
	config.Repos = updated
}
