package main

import (
	termbox "github.com/nsf/termbox-go"
)

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func termboxPoll() chan termbox.Event {
	evCh := make(chan termbox.Event)

	go func() {
		for {
			evCh <- termbox.PollEvent()
		}
	}()

	return evCh
}
