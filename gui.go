package main

import (
	"log"
	"github.com/andlabs/ui"
)

var window *ui.Window

func main() {
	err := ui.Main(func() {
		window = ui.NewWindow("Basic GUI!", 250, 80, false)
		box := ui.NewVerticalBox()

		saveButton := ui.NewButton("Exit")

		saveButton.OnClicked(func(b *ui.Button) {
			ui.Quit()
		})
		box.Append(saveButton, false)

		window.SetChild(box)
		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()

	})
	if err != nil {
		log.Fatalln(err)
	}
}
