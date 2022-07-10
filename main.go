package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	_ "embed"
)

func main() {
	name := os.Args[1]
	if err := MakePluginDirs(name); err != nil {
		panic("cannot make dirs")
	}
	if err := PutLicenseFile(name); err != nil {
		panic("cannot put license file")
	}
	if err := InitializeGit(name); err != nil {
		panic("cannot initialize git")
	}
}

// MakePluginDirs creates directories indispensable for vim plugin.
func MakePluginDirs(name string) error {
	var err error
	for _, dirname := range listPluginDirs(name) {
		err = os.Mkdir(dirname, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

var pluginDirs = [...]string{
	"plugin",
	"autoload",
	"doc",
}

func listPluginDirs(name string) []string {
	ds := []string{name}
	for _, d := range pluginDirs {
		ds = append(ds, filepath.Join(name, d))
	}
	return ds
}

//go:embed LICENSE
var MITLicense []byte

// PutLicenseFile creates a MIT license file.
func PutLicenseFile(dirname string) (err error) {
	f, err := os.Create(filepath.Join(dirname, "LICENSE"))
	if err != nil {
		return err
	}
	defer func() {
		if closingErr := f.Close(); closingErr != nil {
			err = closingErr
		}
	}()
	n, err := f.Write(MITLicense)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("cannot write to LICENSE")
	}
	return nil
}

// InitializeGit initialize git repository.
func InitializeGit(dirname string) error {
	cmd := exec.Command("git", "init", "-q")
	cmd.Dir = dirname
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
