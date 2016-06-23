package main

import "github.com/nsf/termbox-go"
import "strings"

var frame = 0

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := strings.Split(frames[frame], "\n")
	for x, line := range lines {
		for y, cell := range line {
			termbox.SetCell(y, x, cell, colors[frame], termbox.ColorDefault)
		}
	}

	termbox.Flush()
	frame++
	if frame > 8 {
		frame = 0
	}
}
