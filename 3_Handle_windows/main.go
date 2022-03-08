package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello world"))
	w.SetMaster() //Master page (close w = close app)
	w.Show()      //1) Show first

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewButton("Open New", func() { //New button in w2 => func
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show() //To show the windows
	}))
	w2.Resize(fyne.NewSize(100, 100)) //Resize w2 100px 100px
	w2.Show()
	a.Run() //1) Run at the end
}
