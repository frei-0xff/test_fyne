package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"github.com/frei-0xff/test_fyne/internal/assets"
)

func main() {
	a := app.New()
	//a.SetIcon(assets.Icon)
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
