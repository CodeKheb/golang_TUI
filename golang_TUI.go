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

type view struct {
	addedNames string
	hasBusiness bool
}

var contacts []Contact

var app = tview.NewApplication().EnableMouse(true)
var form = tview.NewForm()
var viewContacts = tview.NewList()

var viewedContact []view
var pages = tview.NewPages()
var country = []string{"PH", "US", "CN", "RS"}

var insideForm bool = false

func main() {
	text := tview.NewTextView().	
	SetTextColor(tcell.ColorBlue).
	SetText("(a) to add new contact (q) to quit, (v) to view contacts").
	SetDynamicColors(true).
	SetTextAlign(tview.AlignCenter)

	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	mainFlex.AddItem(tview.NewBox(), 0, 1, false)
	mainFlex.AddItem(text, 1, 0, false)
	mainFlex.SetBorder(true)
	mainFlex.SetTitle("TESTING")

	pages.AddPage("TEST", mainFlex, true, true)
	pages.AddPage("ADD CONTACTS", form, true, false)
	pages.AddPage("VIEW CONTACTS", viewContacts, true, false)




	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		currentPage, _ := pages.GetFrontPage()

		if currentPage == "ADD CONTACTS" {
			return event
		}


		switch event.Rune() {
		case 'q':
			app.Stop()
		case 'a':

			form.Clear(true)
			if !insideForm {
				addContactForm()
				insideForm = true
			}
			pages.SwitchToPage("ADD CONTACTS")
			app.SetFocus(form)
		case 'v':
			form.Clear(true)
			viewContactForm()

			pages.SwitchToPage("VIEW CONTACTS")
			app.SetFocus(viewContacts)

		}

		return event
	})
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func viewContactForm() {
	viewContacts.Clear()

	for _, con := range contacts {
		name := con.firstName + " " + con.lastName
		details := "Email: " + con.email + " Phone Number: " + con.phoneNumber + " Country: " + con.country

		viewContacts.AddItem(name, details, 0, nil)
	}

	viewContacts.AddItem("(a) to add new contact (q) to quit", "", 0, nil) 



}	

func addContactForm() {
	form.Clear(true)
	contact := Contact{}

	form.AddInputField("First Name", "", 40, nil, func(firstName string) {
		contact.firstName = firstName
	})
	form.AddInputField("Last Name", "", 40, nil, func(lastName string) {
		contact.lastName = lastName
	})
	form.AddInputField("Email", "", 40, nil, func(email string) {
		contact.email = email 
	})
	form.AddInputField("Phone", "", 40, nil, func(phoneNumber string) {
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
