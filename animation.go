package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
)

//go:embed parrot
var parrot []byte

type Animation struct {
	Frames [][]byte
}

func LoadFromFile(file string) (*Animation, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadFromBytes(b)
}

func LoadFromBytes(b []byte) (*Animation, error) {
	frames := bytes.Split(b, []byte("!--FRAME--!\n"))

	if len(frames) <= 1 {
		return nil, fmt.Errorf("no frames found")
	}

	for i, frame := range frames {
		if len(frame) == 0 {
			return nil, fmt.Errorf("invalid animation: frame %d is empty", i)
		}
	}

	return &Animation{frames}, nil
}
