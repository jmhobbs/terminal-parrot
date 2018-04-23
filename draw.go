package main

import "github.com/nsf/termbox-go"
import "strings"

var frame_index = 0

func reverse(lines []string) []string {
	newLines := make([]string, len(lines))
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		newLines[i], newLines[j] = lines[j], lines[i]
	}
	return newLines
}

func draw(orientation string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := strings.Split(frames[frame_index], "\n")

	if orientation == "aussie" {
		lines = reverse(lines)
	}

	for x, line := range lines {
		for y, cell := range line {
			termbox.SetCell(y, x, cell, colors[frame_index], termbox.ColorDefault)
		}
	}

	termbox.Flush()
	frame_index++
	if frame_index == num_frames {
		frame_index = 0
	}
}
