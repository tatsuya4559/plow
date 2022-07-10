package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	_ "embed"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	if err := MakePluginDirs(name); err != nil {
		panic("cannot make dirs")
	}
	if err := PutPluginFile(name); err != nil {
		panic(fmt.Sprintf("cannot put plugin/%s.vim", name))
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

//go:embed template/LICENSE
var MITLicense string

// PutLicenseFile creates a MIT license file.
func PutLicenseFile(dirname string) error {
	licensePath := filepath.Join(dirname, "LICENSE")
	return createFileFromTemplate(licensePath, MITLicense, nil)
}

//go:embed template/plugin.vim
var PluginFile string

// PutPluginFile creates plugin/<name>.vim.
func PutPluginFile(dirname string) error {
	pluginName := filepath.Base(dirname)
	pluginFilePath := filepath.Join(dirname, "plugin", fmt.Sprintf("%s.vim", pluginName))
	data := map[string]any{
		"Name": pluginName,
	}
	return createFileFromTemplate(pluginFilePath, PluginFile, data)
}

func createFileFromTemplate(path, tmpl string, data map[string]any) (err error) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot make directory for %s: %w", path, err)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if closingErr := f.Close(); closingErr != nil {
			err = closingErr
		}
	}()

	tpl, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}
	if err := tpl.Execute(f, data); err != nil {
		return err
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
