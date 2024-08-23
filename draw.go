package main

import (
	"bytes"
	"slices"

	"github.com/nsf/termbox-go"
)

var frame_index = 0

func reverse(lines [][]byte) [][]byte {
	slices.Reverse(lines)
	return lines
}

func draw(animation Animation, orientation string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := bytes.Split(animation.Frames[frame_index], []byte{'\n'})

	if orientation == "aussie" {
		lines = reverse(lines)
	}

	for x, line := range lines {
		for y, cell := range line {
			termbox.SetCell(y, x, rune(cell), colors[frame_index%len(colors)], termbox.ColorDefault)
		}
	}

	termbox.Flush()
	frame_index++
	if frame_index >= len(animation.Frames) {
		frame_index = 0
	}
}
