package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) { //return entry & label widget
	out := widget.NewLabel("Hello World !") // The label
	in := widget.NewEntry()                 //Entry = Area text
	in.OnChanged = func(content string) {   //change in real time "out"
		out.SetText("Hello " + content + " !")
	}
	return out, in
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello person")

	w.SetContent(container.NewVBox(makeUI())) //set the container
	w.ShowAndRun()
}

/*
func main() {
	a := app.New()
	w := a.NewWindow("Hello person")

	w.SetContent(container.NewVBox( //new container (similar to div in html)
		widget.NewLabel("Hello World"), //Name of "text"
		widget.NewEntry(),              //Text container
	))
	w.ShowAndRun()
}
*/
