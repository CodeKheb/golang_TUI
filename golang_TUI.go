package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
type Contact struct {
	firstName string
	lastName string
	email string
	phoneNumber string
	state string
	business bool
}
var contacts []Contact
var app = tview.NewApplication().EnableMouse(true)

func main() {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Hello World")
	screen, err := tcell.NewScreen()

	text := tview.NewTextView().	
		SetTextColor(tcell.ColorBlue).
		SetText("(q) to quit").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(box, 3, 1, false)
	flex.AddItem(text, 1, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 { // 113 for letter q in ascii hehe
			app.Stop()
			}
			return event
		})
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}



	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Println("ERROR SCREEN:", err)
		os.Exit(1)
}
//	defer screen.Fini()

//	options := []string{"(1)", "(2)", "(3)", "(4)"}
//	selected := 0

//	for {
//		screen.Clear()

//		for i, opt := range options {
//			style := tcell.StyleDefault
//			if i == selected {
//				style = style.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
//			}
//			for x, r := range opt {
//				screen.SetContent(x, i, r, nil, style)
//			}
//		}

//		screen.Show()

//	}


	screen.Clear()
	screen.Show()
}


