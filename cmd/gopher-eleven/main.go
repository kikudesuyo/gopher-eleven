package main

import (
	"fmt"

	"github.com/kikudesuyo/gopher-eleven/internal/match"
)

func main() {
	match := match.InitMatch()
	for {
		cn, tn, isEnd, err := match.Proceed()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(cn + "「" + tn + "!!」")
		if isEnd {
			return
		}
	}
}
