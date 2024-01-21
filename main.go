package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	_ "embed"
)

//go:embed template/LICENSE
var MITLicense string

//go:embed template/plugin.vim
var PluginFile string

func main() {
	run(os.Args)
}

func run(args []string) {
	fset := parseFlag(args)
	dir := fset.Arg(0)
	files := newPluginFiles(dir)
	for _, f := range files {
		if err := f.create(); err != nil {
			log.Fatal(err)
		}
	}
	printDirTree(dir)
	if err := gitInit(dir); err != nil {
		log.Fatal(err)
	}
}

func parseFlag(args []string) *flag.FlagSet {
	fset := flag.NewFlagSet(args[0], flag.ExitOnError)
	fset.Usage = func() {
		fmt.Fprintf(fset.Output(), "plow: a scaffolding tool for vim plugin.\n\n")
		fmt.Fprintf(fset.Output(), "Usage: plow DIR\n")
	}
	fset.Parse(args[1:])
	return fset
}

type pluginFile struct {
	path string
	tpl  string
	data map[string]any
}

func newPluginFiles(dir string) []*pluginFile {
	name := getPluginName(dir)
	files := make([]*pluginFile, 0, 3)
	files = append(files, &pluginFile{
		path: filepath.Join(dir, "LICENSE"),
		tpl:  MITLicense,
		data: nil,
	})
	files = append(files, &pluginFile{
		path: filepath.Join(dir, "plugin", fmt.Sprintf("%s.vim", name)),
		tpl:  PluginFile,
		data: map[string]any{"Name": name},
	})
	files = append(files, &pluginFile{
		path: filepath.Join(dir, "autoload", fmt.Sprintf("%s.vim", name)),
		tpl:  "",
		data: nil,
	})
	return files
}

func getPluginName(dir string) string {
	name := filepath.Base(dir)
	name = strings.TrimSuffix(name, ".vim")
	name = strings.TrimPrefix(name, "vim-")
	return name
}

func (f *pluginFile) create() error {
	dir := filepath.Dir(f.path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot make directory for %s: %w", f.path, err)
		}
	}

	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer file.Close()

	tpl, err := template.New("").Parse(f.tpl)
	if err != nil {
		return err
	}
	if err := tpl.Execute(file, f.data); err != nil {
		return err
	}

	return nil
}

func printDirTree(rootDir string) error {
	return filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
}

func gitInit(dirname string) error {
	cmd := exec.Command("git", "init", "-q")
	cmd.Dir = dirname
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cannot init git: %w", err)
	}
	return nil
}
