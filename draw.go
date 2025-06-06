package main

import (
	"bytes"
	"slices"

	"github.com/nsf/termbox-go"
)

var frame_index = 0
var color_index = 0

func reverse(lines [][]byte) [][]byte {
	slices.Reverse(lines)
	return lines
}

func draw(animation Animation, orientation string) {
	_ = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := bytes.Split(animation.Frames[frame_index], []byte{'\n'})

	if orientation == "aussie" {
		lines = reverse(lines)
	}

	for x, line := range lines {
		for y, cell := range line {
			termbox.SetCell(y, x, rune(cell), colors[color_index], termbox.ColorDefault)
		}
	}

	_ = termbox.Flush()
	frame_index++
	color_index++
	if frame_index >= len(animation.Frames) {
		frame_index = 0
	}
	if color_index >= len(colors) {
		color_index = 0
	}
}
