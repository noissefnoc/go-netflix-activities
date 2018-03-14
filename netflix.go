package main

import (
	"github.com/sclevine/agouti"
	"fmt"
)

// Netflix is struct for netflix page scraping
type Netflix struct {
	LoginURL           string
	ViewingHistoryURL  string
	ViewingHistoryHTML []byte
	Debug              bool
}

var netflixURL = "https://www.netflix.com"

// Following CSS selectors may change various reasons.
// So I extract those to variables.
var loginFormIDSelector = "#email"
var loginFormPasswordSelector = "#password"
var loginFormSubmitButtonSelector = ".login-button"
var viewingHistoryListSelector = "li.retableRow"
var viewingHistoryDateSelector = ".col.date.nowrap"
var viewingHistoryTitleSelector = ".col.title a"

// FetchViewingHistory is API to fetch Netflix viewing history by scraping
func (n *Netflix) FetchViewingHistory(email string, password string) (error) {
	// TODO: enable debug option
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
			"--allow-insecure-localhost",
		}),
	)

	if err := driver.Start(); err != nil {
		return fmt.Errorf("failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()

	if err != nil {
		return fmt.Errorf("faild to open page:%v", err)
	}

	if err := page.Navigate(n.LoginURL); err != nil {
		return fmt.Errorf("failed to navigate:%v", err)
	}

	id := page.Find(loginFormIDSelector)
	pass := page.Find(loginFormPasswordSelector)
	id.Fill(email)
	pass.Fill(password)

	if err := page.Find(loginFormSubmitButtonSelector).Submit(); err != nil {
		return fmt.Errorf("failed to login:%v", err)
	}

	if err := page.Navigate(n.ViewingHistoryURL); err != nil {
		return fmt.Errorf("failed to navigate:%v", err)
	}

	html, err := page.HTML()

	if err != nil {
		return fmt.Errorf("failed to get html:%v", err)
	}

	n.ViewingHistoryHTML = []byte(html)

	return err
}
