package main

import (
//	"fmt"
//	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
type Contact struct {
	firstName string
	lastName string
	email string
	phoneNumber string
	country string
	business bool
}
var contacts []Contact
var app = tview.NewApplication().EnableMouse(true)
var form = tview.NewForm()
var pages = tview.NewPages()
var country = []string{"PH", "US", "CN", "RS"}

var insideForm bool = false

func main() {
	text := tview.NewTextView().	
		SetTextColor(tcell.ColorBlue).
		SetText("(a) to add new contact (q) to quit").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)



	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	mainFlex.AddItem(tview.NewBox(), 0, 1, false)
	mainFlex.AddItem(text, 1, 0, false)
	mainFlex.SetBorder(true)
	mainFlex.SetTitle("TESTING")

	pages.AddPage("TEST", mainFlex, true, true)
	pages.AddPage("ADD CONTACTS", form, true, false)



	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		currentPage, _ := pages.GetFrontPage()

		if currentPage == "ADD CONTACTS" {
			return event
		}


		switch event.Rune() {
		case 'q':
			app.Stop()
		case 'a':
			if !insideForm {
				addContactForm()
				insideForm = true
			}
			pages.SwitchToPage("ADD CONTACTS")
			app.SetFocus(form)
			}
			return event
		})
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
func addContactForm() {
	form.Clear(true)
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})
	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})
	form.AddInputField("Email", "", 20, nil, func(email string) {
		contact.email = email 
	})
	form.AddInputField("Phone", "", 20, nil, func(phoneNumber string) {
		contact.phoneNumber = phoneNumber 
	})

	form.AddDropDown("Country", country, 0, func(country string, index int) {
		contact.country = country
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		pages.SwitchToPage("TEST")
		app.SetFocus(pages)
	})

}
