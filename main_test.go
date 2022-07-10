package main_test

import (
	"fmt"
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

func assertExists(t *testing.T, dirname string) {
	t.Helper()
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		t.Errorf("%s does not exist", dirname)
	}
}

func assertIsNotEmpty(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("cannot get stat of %s: %v", path, err)
	}
	if info.Size() <= 0 {
		t.Errorf("%s is empty", path)
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

func TestPutLicenseFile(t *testing.T) {
	tempdirname := prepareTempDir(t, "PutLicenseFile")
	defer cleanupTempDir(t, tempdirname)

	if err := PutLicenseFile(tempdirname); err != nil {
		t.Fatalf("PutLicenseFile(%q) got an error: %v", tempdirname, err)
	}

	licenseFilepath := filepath.Join(tempdirname, "LICENSE")
	assertExists(t, licenseFilepath)
	assertIsNotEmpty(t, licenseFilepath)
}

func TestPutPluginFile(t *testing.T) {
	tempdirname := prepareTempDir(t, "PutPluginFile")
	defer cleanupTempDir(t, tempdirname)

	if err := PutPluginFile(tempdirname); err != nil {
		t.Fatalf("PutPluginFile(%q) got an error: %v", tempdirname, err)
	}

	pluginFilepath := filepath.Join(
		tempdirname,
		"plugin",
		fmt.Sprintf("%s.vim", filepath.Base(tempdirname)),
	)
	assertExists(t, pluginFilepath)
	assertIsNotEmpty(t, pluginFilepath)
}

func TestInitializeGit(t *testing.T) {
	tempdirname := prepareTempDir(t, "PutLicenseFile")
	defer cleanupTempDir(t, tempdirname)

	if err := InitializeGit(tempdirname); err != nil {
		t.Fatalf("InitializeGit(%q) got an error: %v", tempdirname, err)
	}

	assertExists(t, filepath.Join(tempdirname, ".git"))
}
