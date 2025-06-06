package main

import (
	"os"
	"testing"
)

func Test_LoadFromFile(t *testing.T) {
	a, err := LoadFromFile(os.DirFS("animations"), "parrot.animation")
	if err != nil {
		t.Fatal(err)
	}

	if len(a.Frames) != 10 {
		for _, frame := range a.Frames {
			t.Log(string(frame))
		}
		t.Fatalf("expected 10 frames, got %d", len(a.Frames))
	}
}

func Test_LoadFromBytes(t *testing.T) {
	t.Run("invalid: no frames", func(t *testing.T) {
		_, err := LoadFromBytes([]byte{})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("invalid: one frame", func(t *testing.T) {
		_, err := LoadFromBytes([]byte("frame"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("invalid: empty frame", func(t *testing.T) {
		_, err := LoadFromBytes([]byte("!--FRAME--!\n"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("valid", func(t *testing.T) {
		a, err := LoadFromBytes([]byte("description: test\n!--FRAME--!\nA\n!--FRAME--!\nB\n"))
		if err != nil {
			t.Fatal(err)
		}

		if len(a.Frames) != 2 {
			t.Fatalf("expected 2 frames, got %d", len(a.Frames))
		}
	})
}
