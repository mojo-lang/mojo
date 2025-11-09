package mpm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGitLastCommit(t *testing.T) {
	commit := GetGitLatestCommit(".")
	assert.NotNil(t, commit)
}
