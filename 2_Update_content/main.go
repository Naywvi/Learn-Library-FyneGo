package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Update Content")
	clock := widget.NewLabel("") //create a content

	updateTime(clock) //1) Send

	w.SetContent(clock)
	go func() { //Reset with "goroutine func" => go func() ...
		for range time.Tick(time.Second) { //All second =>infinit loop
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}
func updateTime(clock *widget.Label) { //2
	formatted := time.Now().Format("The time is : 03:04:05") //Set time "03:04:05"
	clock.SetText(formatted)                                 //Set clock with text
}
