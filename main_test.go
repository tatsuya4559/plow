package main

import (
	"os"
	"path/filepath"
	"testing"
)

func assertExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("%s does not exist", path)
	}
}

func TestRun(t *testing.T) {
	dir := t.TempDir()
	arg := filepath.Join(dir, "vim-foo")
	run([]string{"plow", arg})

	assertExists(t, filepath.Join(arg, "plugin"))
	assertExists(t, filepath.Join(arg, "autoload"))
	assertExists(t, filepath.Join(arg, "LICENSE"))
	assertExists(t, filepath.Join(arg, "plugin", "foo.vim"))
	assertExists(t, filepath.Join(arg, "autoload", "foo.vim"))
}
