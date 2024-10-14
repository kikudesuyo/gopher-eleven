package main

import (
	"github.com/kikudesuyo/gopher-eleven/internal/match"
)

func main() {
	match := match.InitMatch()
	for {
		display, isMatchEnd := match.Proceed()
		display.Print()
		if isMatchEnd {
			return
		}
	}
}
