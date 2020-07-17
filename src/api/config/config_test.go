package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	assert.EqualValues(t, "SCRET_GITHUB_ACCESS_TOKEN", secretGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GetGithubAccessToken())
}
