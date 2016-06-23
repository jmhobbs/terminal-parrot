package main

import "github.com/nsf/termbox-go"
import "time"
import "flag"

func main() {
	loops := flag.Int("loops", 0, "number of times to loop (default: infinite)")
	delay := flag.Int("delay", 75, "frame delay in ms")
	orientation := flag.String("orientation", "regular", "regular or aussie")
	flag.Parse()

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
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC) {
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
