package main

import (
	"fmt"

	"github.com/kikudesuyo/gopher-eleven/internal/match"
)

func main() {
	match := match.InitMatch()
	for {
		display, isEnd, err := match.Proceed()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		display.Print()
		if isEnd {
			return
		}
	}
}
