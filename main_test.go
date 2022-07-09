package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMkdir(t *testing.T) {
	tempdirname, err := os.MkdirTemp("", "TestMkdir")
	if err != nil {
		t.Fatalf("cannot make tempdir: %v", err)
	}
	if err := os.Remove(tempdirname); err != nil {
		t.Fatalf("cannot remove tempdir: %v", err)
	}

	if err := MakePluginDirs(tempdirname); err != nil {
		t.Fatalf("Mkdir(%q) got an error: %v", tempdirname, err)
	}

	defer func() {
		os.RemoveAll(tempdirname)
	}()

	assertExists(t, tempdirname)
	assertExists(t, filepath.Join(tempdirname, "plugin"))
	assertExists(t, filepath.Join(tempdirname, "autoload"))
	assertExists(t, filepath.Join(tempdirname, "doc"))
}

func assertExists(t *testing.T, dirname string) {
	t.Helper()
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		t.Errorf("%s does not exist", dirname)
	}
}
