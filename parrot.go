package main

import "github.com/nsf/termbox-go"
import "time"
import "syscall"
import "flag"
import "unsafe"
import "os"
import "fmt"

// Taken from https://github.com/golang/crypto/blob/master/ssh/terminal/util.go#L31
func IsTerminal(fd int) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), syscall.TIOCGETA, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func main() {
	loops := flag.Int("loops", 0, "number of times to loop (default: infinite)")
	flag.Parse()

	if IsTerminal(int(os.Stdout.Fd())) {
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
	for l := 0; l < loops; l++ {
		for i := 0; i <= 9; i++ {
			fmt.Print(frames[i], "\n")
			time.Sleep(75 * time.Millisecond)
		}
	}
}
