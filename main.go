package main

import (
	"os"
	"path/filepath"
)

func main() {
	name := os.Args[1]
	if err := MakePluginDirs(name); err != nil {
		panic("cannot make dirs")
	}
}

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
