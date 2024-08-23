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
	//framePath := flag.String("path", "/etc/terminal-parrot;/opt/homebrew/etc/terminal-parrot", "path to additional frame files")
	loops := flag.Int("loops", 0, "number of times to loop (default: infinite)")
	delay := flag.Int("delay", 75, "frame delay in ms")
	orientation := flag.String("orientation", "regular", "regular or aussie")
	list := flag.Bool("list", false, "list available animations and exit")
	flag.Parse()

	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Fprintf(os.Stderr, "%s must be run in a terminal!\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if *list {
		fmt.Println("Available animations:\n")
		longestName := 0
		for name := range Animations {
			longestName = max(longestName, len(name))
		}
		fmtString := fmt.Sprintf("  %% %ds : %%s\n", longestName)

		for name, animation := range Animations {
			description := ""
			if mdDescription, ok := animation.Metadata["description"]; ok {
				description = mdDescription
			}
			fmt.Printf(fmtString, name, description)
		}
		os.Exit(0)
	}

	// Load an Animation

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
	draw(Animations["parrot"], *orientation)

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
			draw(Animations["parrot"], *orientation)
			time.Sleep(time.Duration(*delay) * time.Millisecond)
		}
	}
}
