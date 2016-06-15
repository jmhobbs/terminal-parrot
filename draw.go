package main

import "github.com/nsf/termbox-go"

var frame = 0

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	chars := len(frames[frame])
	x := 8
	y := 1
	for index := 0; index < chars; index++ {
		if '\n' == frames[frame][index] {
			y++
			x = 8
			continue
		}
		if ' ' == frames[frame][index] {
			termbox.SetCell(x, y, rune(frames[frame][index]), termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(x, y, rune(frames[frame][index]), colors[frame], termbox.ColorDefault)
		}
		x++
	}

	termbox.Flush()
	frame++
	if frame > 8 {
		frame = 0
	}
}
