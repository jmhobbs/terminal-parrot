package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed parrot
var parrot []byte

type Animation struct {
	Metadata map[string]string
	Frames   [][]byte
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

	if len(frames) <= 2 {
		return nil, fmt.Errorf("no frames found")
	}

	// The first "frame" is actually the metadata.
	metadata := make(map[string]string)
	for _, line := range bytes.Split(frames[0], []byte{'\n'}) {
		parts := bytes.SplitN(line, []byte{':'}, 2)
		if len(parts) != 2 {
			continue
		}
		metadata[strings.TrimSpace(string(parts[0]))] = strings.TrimSpace(string(parts[1]))
	}

	for i, frame := range frames[1:] {
		if len(frame) == 0 {
			return nil, fmt.Errorf("invalid animation: frame %d is empty", i)
		}
	}

	return &Animation{metadata, frames[1:]}, nil
}
