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
		fmt.Println(text)
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(1200 * time.Millisecond)
}
