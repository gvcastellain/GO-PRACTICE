package main

import (
	"os"
)

type FileMethods interface {
	CreateFile(name string) error
}

type File struct {
	Name string
}

func (f *File) CreateFile(name string) error {
	file, err := os.Create(name)

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
