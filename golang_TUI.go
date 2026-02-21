package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Println("ERROR SCREEN:", err)
		os.Exit(1)
}
	defer screen.Fini()

	options := []string{"1", "2"}
	selected := 0

	for {
		screen.Clear()

		for i, opt := range options {
			style := tcell.StyleDefault
			if i == selected {
				style = style.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
			}
			for x, r := range opt {
				screen.SetContent(x, i, r, nil, style)
			}
		}

		screen.Show()

	}


	screen.Clear()
	screen.Show()
}


