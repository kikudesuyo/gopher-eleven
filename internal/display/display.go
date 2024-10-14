package display

import (
	"fmt"
	"time"
)

type Display struct {
	texts []string
}

func NewDisplay(texts ...string) Display {
	return Display{texts: texts}
}

func (d Display) Print() {
	for _, text := range d.texts {
		for _, char := range text {
			fmt.Print(string(char))
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println()
		// fmt.Println(text)
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(1200 * time.Millisecond)
}
