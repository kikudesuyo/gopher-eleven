package display

import "fmt"

type Display struct {
	texts []string
}

func NewDisplay(texts ...string) Display {
	return Display{texts: texts}
}

func (d Display) Print() {
	for _, text := range d.texts {
		fmt.Println(text)
	}
}
