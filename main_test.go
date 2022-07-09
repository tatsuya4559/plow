package main_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/tatsuya4559/plugin-template.vim"
)

func prepareTempDir(t *testing.T, prefix string) string {
	t.Helper()
	tempdirname, err := os.MkdirTemp("", prefix)
	if err != nil {
		t.Fatalf("cannot make tempdir: %v", err)
	}
	return tempdirname
}

func cleanupTempDir(t *testing.T, dirname string) {
	t.Helper()
	if err := os.RemoveAll(dirname); err != nil {
		if err, ok := err.(*os.PathError); ok {
			t.Fatalf("cannot cleanup temp dir %q: %v", dirname, err)
		}
	}
}

func TestMakePluginDirs(t *testing.T) {
	tempdirname := prepareTempDir(t, "MakePluginDirs")
	if err := os.Remove(tempdirname); err != nil {
		t.Fatalf("cannot remove tempdir: %v", err)
	}

	if err := MakePluginDirs(tempdirname); err != nil {
		t.Fatalf("Mkdir(%q) got an error: %v", tempdirname, err)
	}

	defer cleanupTempDir(t, tempdirname)

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

func TestPutLicenseFile(t *testing.T) {
	tempdirname := prepareTempDir(t, "PutLicenseFile")
	defer cleanupTempDir(t, tempdirname)

	if err := PutLicenseFile(tempdirname); err != nil {
		t.Fatalf("PutLicenseFile(%q) got an error: %v", tempdirname, err)
	}

	assertExists(t, filepath.Join(tempdirname, "LICENSE"))
}