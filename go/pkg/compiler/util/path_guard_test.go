package util

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathGuardClearsOnlyOnce(t *testing.T) {
	dir := t.TempDir()
	old := writeTestFile(t, dir, "old.go", generatedHeader)
	manual := writeTestFile(t, dir, "manual.go", "package manual\n")
	guard := &PathGuard{OnlyClearGenerated: true, Suffixes: []string{".go"}}
	require.NoError(t, guard.Check(dir))
	require.NoFileExists(t, old)
	require.FileExists(t, manual)
	fresh := writeTestFile(t, dir, "fresh.go", generatedHeader)
	require.NoError(t, guard.Check(dir))
	require.NoError(t, guard.Clear(dir))
	require.FileExists(t, fresh)
}

func TestPathGuardCreatesAndPreservesNewDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nested", "output")
	guard := &PathGuard{Suffixes: []string{".go"}}
	require.NoError(t, guard.Check(dir))
	require.DirExists(t, dir)
	fresh := writeTestFile(t, dir, "fresh.go", generatedHeader)
	require.NoError(t, guard.Check(dir))
	require.FileExists(t, fresh)
}

func TestPathGuardDisableClear(t *testing.T) {
	dir := t.TempDir()
	file := writeTestFile(t, dir, "old.go", generatedHeader)
	guard := &PathGuard{DisableClear: true, Suffixes: []string{".go"}}
	require.NoError(t, guard.Check(dir))
	require.FileExists(t, file)
}

func TestPathGuardFailedClearCanBeRetried(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "output")
	guard := &PathGuard{Suffixes: []string{".go"}}
	require.Error(t, guard.Clear(dir))
	file := writeTestFile(t, dir, "old.go", generatedHeader)
	require.NoError(t, guard.Clear(dir))
	require.NoFileExists(t, file)
	blocked := writeTestFile(t, t.TempDir(), "file", "not a directory")
	require.Error(t, guard.Check(filepath.Join(blocked, "output")))
}
