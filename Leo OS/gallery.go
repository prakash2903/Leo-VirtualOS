package main

// import fyne

import (
	"image/color"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showGalleryApp(window fyne.Window) {
	// New app
	a := app.New()
	// New title
	w := a.NewWindow("GALLERY")

	//resize
	w.Resize(fyne.NewSize(400, 400))
	btn := widget.NewButton("CLICK TO VIEW THE FILES", func() {
		// dialog for opening files
		// 2 arguments
		rectangle := canvas.NewRectangle(color.Black)
		dialog.ShowCustom("Color Picked", "Ok", rectangle, w)
		fileDialog := dialog.NewFileOpen(
			// _ to ignore error
			func(uc fyne.URIReadCloser, _ error) {
				// reader to read data
				data, _ := ioutil.ReadAll(uc)
				// static resource
				// 2 arguments
				// first is file name (string)
				// second is data from reader
				res := fyne.NewStaticResource(uc.URI().Name(), data)
				// Now image widget to display our image
				img := canvas.NewImageFromResource(res)

				// setup new window for image and set content
				w := fyne.CurrentApp().NewWindow(uc.URI().Name())
				w.SetContent(img)
				// resize window
				w.Resize(fyne.NewSize(400, 400))
				w.Show() // display our image
			}, w)

		// filtering files
		fileDialog.SetFilter(
			// filter jpg and png
			// ignore rest of the files
			storage.NewExtensionFileFilter([]string{".png", ".jpg"}))
		fileDialog.Show()
		// we are done :)

	})

	// display button in parent window
	w.SetContent(btn)
	w.Show()

}
