//+build indows linux arwin
package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	hello := widget.NewLabel("Hello Gui")
	w.SetContent(
		widget.NewVBox(
			hello, widget.NewButton("Hi", func() {
				hello.SetText("Wellcome :)")
			})))
	w.ShowAndRun()
}
