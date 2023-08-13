package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/widget"
)

const apiKey = "e9d63b0648d63d31affff098a151458f"

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func showWeatherApp(window fyne.Window) {
	a := app.New()
	w := a.NewWindow("Weather")

	cityEntry := widget.NewEntry()
	cityEntry.SetPlaceHolder("Enter city name")

	tempLabel := widget.NewLabel("")

	searchButton := widget.NewButton("Search", func() {
		city := cityEntry.Text
		weatherData, err := getWeatherData(city)
		if err != nil {
			tempLabel.SetText(fmt.Sprintf("Error: %v", err))
			return
		}
		tempLabel.SetText(fmt.Sprintf("%s: %.2f °C", weatherData.Name, weatherData.Main.Temp))
	})

	content := container.NewVBox(
		cityEntry,
		searchButton,
		tempLabel,
	)
	w.SetContent(content)
	w.Resize(fyne.Size{Width: 400, Height: 300})
	w.Show()
}

func getWeatherData(city string) (*WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
