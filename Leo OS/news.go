package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"strconv"
	// "time"

	"fyne.io/fyne"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Article represents a news article
type Article struct {
	Title       string
	Description string
	URL         string
	Source      struct {
		Name string
	}
}

// NewsAPIResponse represents the response from the News API
type NewsAPIResponse struct {
	Articles []Article
}

func showNewsApp(window fyne.Window) {
	// Create a new Fyne app
	a := app.New()

	// Create a window with the specified size and title
	w := a.NewWindow("News App")
	w.Resize(fyne.NewSize(600, 800))

	// Create a list widget to display news articles
	newsList := widget.NewList(
		func() int {
			return len(articles)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(index int, item fyne.CanvasObject) {
			label := item.(*widget.Label)
			article := articles[index]
			label.SetText(fmt.Sprintf("%s\n%s\n%s", article.Title, article.Source.Name, article.Description))
		},
	)

	// Create a button to refresh the news
	refreshButton := widget.NewButton("Refresh", func() {
		// Fetch the latest news articles from the News API
		newsAPIKey := "73eb2432b02347b79ad0e151de8a5586"
		newsAPIURL := "https://newsapi.org/v2/top-headlines?country=us&apiKey=" + newsAPIKey
		resp, err := http.Get(newsAPIURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		// Parse the response from the News API
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		var newsAPIResponse NewsAPIResponse
		err = json.Unmarshal(body, &newsAPIResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Update the news list with the latest news articles
		articles = newsAPIResponse.Articles
		newsList.Refresh()
	})

	// Create a box widget to contain the refresh button and the news list
	contentBox := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		refreshButton,
		newsList,
	)

	// Set the content of the window to the content box
	w.SetContent(contentBox)

	// Show the window and start the Fyne event loop
	w.ShowAndRun()
}

// The list of news articles
var articles []Article
