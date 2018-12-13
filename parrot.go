package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	isatty "github.com/mattn/go-isatty"
	"github.com/nsf/termbox-go"
)

func main() {
	loops := flag.Int("loops", 0, "number of times to loop (default: infinite)")
	delay := flag.Int("delay", 75, "frame delay in ms")
	orientation := flag.String("orientation", "regular", "regular or aussie")
	flag.Parse()

	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Fprintf(os.Stderr, "%s must be run in a terminal!\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	termbox.SetOutputMode(termbox.Output256)

	loop_index := 0
	draw(*orientation)

loop:
	for {
		select {
		case ev := <-event_queue:
			if (ev.Type == termbox.EventKey && (ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC)) || ev.Type == termbox.EventInterrupt {
				break loop
			}
		default:
			loop_index++
			if *loops > 0 && (loop_index/9) >= *loops {
				break loop
			}
			draw(*orientation)
			time.Sleep(time.Duration(*delay) * time.Millisecond)
		}
	}
}
