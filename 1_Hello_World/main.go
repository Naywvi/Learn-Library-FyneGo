package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()                  //Create app
	w := a.NewWindow("Hello World") //name of window of the app

	w.SetContent(widget.NewLabel("Hello world")) //Content > string
	w.ShowAndRun()                               //Runnning
}

//To create it "go build ."
