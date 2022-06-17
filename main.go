package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/frei-0xff/inmak_invoice_generator/internal/ui"
)

func main() {
	a := app.NewWithID("com.inmak.invoice_generator")
	w := a.NewWindow("Генератор счетов | ИнМАК")

	w.SetContent(ui.CreateMaster(w))
	w.Resize(fyne.NewSize(1280, 900))
	w.CenterOnScreen()
	w.ShowAndRun()
}
