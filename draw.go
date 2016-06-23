package main

import "github.com/nsf/termbox-go"
import "strings"

var frame = 0

func reverse(lines []string) []string {
	newLines := make([]string, len(lines))
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		newLines[i], newLines[j] = lines[j], lines[i]
	}
	return newLines
}

func draw(orientation string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := strings.Split(frames[frame], "\n")

	if orientation == "aussie" {
		lines = reverse(lines)
	}

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
