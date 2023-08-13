package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("Leo OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget

//var btn6 fyne.Widget

var img fyne.CanvasObject
var Deskbtn fyne.Widget

var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("D:\\Leo OS\\peakpx.jpg")

	btn1 = widget.NewButtonWithIcon("Weather", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalculatorApp(myWindow)
	})

	btn3 = widget.NewButtonWithIcon("TextEditor", theme.FileTextIcon(), func() {
		showTextApp(myWindow)
	})

	btn4 = widget.NewButtonWithIcon("Gallery", theme.MediaPhotoIcon(), func() {
		showGalleryApp(myWindow)
	})

	btn5 = widget.NewButtonWithIcon("Music", theme.MediaMusicIcon(), func() {
		showMusicApp(myWindow)
	})

	//btn6 = widget.NewButtonWithIcon("NewsToday", theme.DocumentIcon(), func() {
	//	showNewsApp(myWindow)

	//})

	Deskbtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(6, Deskbtn, btn1, btn2, btn3, btn4, btn5))

	// Calculate minimum size of panel content
	panelContent.Resize(fyne.NewSize(800, 600))
	myWindow.Resize(panelContent.MinSize())

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()

}
