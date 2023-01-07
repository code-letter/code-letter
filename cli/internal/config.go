package internal

type Config struct {
	localRepoPath string
}

func NewConfig(repoPath string) *Config {
	return &Config{
		localRepoPath: repoPath,
	}
}
