package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showCalculatorApp(window fyne.Window) {

	a := app.New()
	w := a.NewWindow("calc")
	//w.Resize(fyne.NewSize(1920, 1280))

	var val1 int = 0
	var val2 int = 0
	var symbol string = ""
	var total string = ""
	var entryText string = ""

	entry1 := widget.NewEntry()
	entry1.TextStyle = fyne.TextStyle{Bold: true}
	entry1.SetPlaceHolder("Calculator..")

	btn1 := widget.NewButton("1", func() {
		total = total + "1"
		entryText = entryText + "1"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn2 := widget.NewButton("2", func() {
		total = total + "2"
		entryText = entryText + "2"
		entry1.SetText(fmt.Sprint(entryText))

	})
	btn3 := widget.NewButton("3", func() {
		total = total + "3"
		entryText = entryText + "3"
		entry1.SetText(fmt.Sprint(entryText))

	})
	btn4 := widget.NewButton("4", func() {
		total = total + "4"
		entryText = entryText + "4"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn5 := widget.NewButton("5", func() {
		total = total + "5"
		entryText = entryText + "5"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn6 := widget.NewButton("6", func() {
		total = total + "6"
		entryText = entryText + "6"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn7 := widget.NewButton("7", func() {
		total = total + "7"
		entryText = entryText + "7"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn8 := widget.NewButton("8", func() {
		total = total + "8"
		entryText = entryText + "8"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn9 := widget.NewButton("9", func() {
		total = total + "9"
		entryText = entryText + "9"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn0 := widget.NewButton("0", func() {
		total = total + "0"
		entryText = entryText + "0"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnDot := widget.NewButton(".", func() {

	})

	result := canvas.NewText("Result", color.Black)
	result.TextSize = 30
	result.Alignment = fyne.TextAlignCenter
	btnClear := widget.NewButton("Clear", func() {
		entry1.SetText("")
		total = ""
		result.Text = ""
		entryText = ""
	})
	// btnClear.Text = "NoClear"
	btnPlus := widget.NewButton("+", func() {
		val1, _ = strconv.Atoi(total)
		fmt.Print(fmt.Sprint(val1))
		fmt.Print(fmt.Sprint(val2))
		symbol = "+"
		fmt.Print(symbol)

		entryText = entryText + "+"
		entry1.SetText(fmt.Sprint(entryText))
		total = ""
	})
	btnMinus := widget.NewButton("-", func() {
		val1, _ = strconv.Atoi(total)
		fmt.Print(fmt.Sprint(val1))
		fmt.Print(fmt.Sprint(val2))
		symbol = "-"
		fmt.Print(symbol)

		total = ""
		entryText = entryText + "-"
		entry1.SetText(fmt.Sprint(entryText))

	})
	btnMultiply := widget.NewButton("x", func() {
		val1, _ = strconv.Atoi(total)
		fmt.Print(fmt.Sprint(val1))
		fmt.Print(fmt.Sprint(val2))
		symbol = "*"
		fmt.Print(symbol)

		total = ""

		entryText = entryText + "*"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnDivide := widget.NewButton("/", func() {
		val1, _ = strconv.Atoi(total)
		fmt.Print(fmt.Sprint(val1))
		fmt.Print(fmt.Sprint(val2))
		symbol = "/"
		fmt.Print(symbol)

		total = ""

		entryText = entryText + "/"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnEqual := widget.NewButton("=", func() {
		val2, _ := strconv.Atoi(total)
		if symbol == "+" {
			myanswer := (val1) + (val2)
			fmt.Println("Sum : ", myanswer)
			result.Text = fmt.Sprint(myanswer)
			result.Refresh()
		} else if symbol == "-" {
			myanswer := (val1) - (val2)
			fmt.Println("Minus : ", myanswer)
			result.Text = fmt.Sprint(myanswer)
			result.Refresh()
		} else if symbol == "*" {
			myanswer := (val1) * (val2)
			fmt.Println("Multiply : ", myanswer)
			result.Text = fmt.Sprint(myanswer)
			result.Refresh()
		} else {
			myanswer := (val1) / (val2)
			fmt.Println("divide : ", myanswer)
			result.Text = fmt.Sprint(myanswer)
			result.Refresh()
		}
		total = ""
	})

	w.SetContent(
		container.NewVBox(

			entry1,
			result,
			container.NewGridWithColumns(
				4,
				btn1,
				btn2,
				btn3,
				btn4,
				btn5,
				btn6,
				btn7,
				btn8,
				btn9,
				btn0,
				btnDot,
				btnEqual,
				btnPlus,
				btnMinus,
				btnMultiply,
				btnDivide,
				btnClear,
			),
		),
	)
	w.Show()

}
