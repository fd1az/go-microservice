package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigAccesToken(t *testing.T) {
	accessToken := GetGithubAccesToken()

	assert.EqualValues(t, "ccceb1efbf637bb6d80536154efcda4cc8bd52a5", accessToken)
}
