package main

import (
	"bufio"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showTextApp(window fyne.Window) {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window and set its title
	myWindow := myApp.NewWindow("Text Editor")

	// Create a text editor widget
	myText := widget.NewMultiLineEntry()

	// Create a button to create a new file
	newFileBtn := widget.NewButton("New File", func() {
		// Clear the text editor
		myText.SetText("")
	})

	// Create a button to open an existing file
	openFileBtn := widget.NewButton("Open File", func() {
		// Show a file dialog to select a file to open
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				defer reader.Close()
				// Read the contents of the file and set it in the text editor
				scanner := bufio.NewScanner(reader)
				var text string
				for scanner.Scan() {
					text += scanner.Text() + "\n"
				}
				myText.SetText(text)
			}
		}, myWindow)
	})

	// Create a button to save the current file
	saveFileBtn := widget.NewButton("Save File", func() {
		// Show a file dialog to select a file to save
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err == nil && writer != nil {
				defer writer.Close()
				// Write the contents of the text editor to the file
				fmt.Fprint(writer, myText.Text)
			}
		}, myWindow)
	})

	// Create a horizontal box container for the buttons
	buttons := container.NewHBox(
		newFileBtn,
		openFileBtn,
		saveFileBtn,
		layout.NewSpacer(),
	)

	// Create a vertical box container for the text editor and buttons
	content := container.NewVBox(
		myText,
		buttons,
	)

	// Set the window's content to the vertical box container
	myWindow.SetContent(content)

	// Show the window and run the application
	myWindow.Show()
}
