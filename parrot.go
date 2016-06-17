package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"time"
)

func main() {
	loops := flag.Int("loops", 0, "number of times to loop (default: infinite)")
	flag.Parse()

	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		run_termbox(*loops)
	} else {
		run_no_tty(*loops)
	}

}

func run_termbox(loops int) {
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
	draw()

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC) {
				break loop
			}
		default:
			loop_index++
			if loops > 0 && (loop_index/9) >= loops {
				break loop
			}
			draw()
			time.Sleep(75 * time.Millisecond)
		}
	}
}

func run_no_tty(loops int) {
    loop_index := 0
	for  {
        loop_index++
        if loops != 0 && loop_index == loops {
            break
        }
		for i := 0; i <= 9; i++ {
			fmt.Print(frames[i], "\n")
			time.Sleep(75 * time.Millisecond)
		}
	}
}
