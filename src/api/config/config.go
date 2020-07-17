package config

import "os"

const (
	secretGithubAccessToken = "SCRET_GITHUB_ACCESS_TOKEN"
	LogLevel                = "info"
	goEnv                   = "GO_ENVIRONMEMT"
	production              = "production"
)

var (
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}

func IsProduction() bool {
	return os.Getenv(goEnv) == production
}
