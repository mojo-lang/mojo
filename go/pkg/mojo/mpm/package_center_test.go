package mpm

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetGitLastCommit(t *testing.T) {
    commit := GetGitLatestCommit(".")
    assert.NotNil(t, commit)
}
