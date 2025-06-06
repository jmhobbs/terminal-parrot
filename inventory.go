package main

import (
	"embed"
	"io/fs"
	"os"
	"strings"
)

//go:embed animations/*
var animations embed.FS

type Inventory map[string]Animation

func (i Inventory) LoadFromPaths(paths []string) error {
	for _, path := range paths {
		err := i.LoadFromFS(os.DirFS(path))
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func (i Inventory) LoadFromFS(filesystem fs.FS) error {
	files, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".animation") {
			animation, err := LoadFromFile(filesystem, file.Name())
			if err != nil {
				return err
			}
			i[strings.TrimSuffix(file.Name(), ".animation")] = *animation
		}
	}

	return nil
}

func NewInventory() Inventory {
	i := make(Inventory)
	s, err := fs.Sub(animations, "animations")
	if err != nil {
		panic(err)
	}
	err = i.LoadFromFS(s)
	if err != nil {
		panic(err)
	}
	return i
}
