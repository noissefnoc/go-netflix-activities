package main

import (
	"github.com/sclevine/agouti"
	"fmt"
)

// Netflix is struct for netflix page scraping
type Netflix struct {
	Driver             agouti.WebDriver
	LoginURL           string
	ViewingHistoryURL  string
}

var netflixURL = "https://www.netflix.com"

// Following CSS selectors may change various reasons.
// So I extract those to variables.
// Login page selector
var loginFormIDSelector = "#email"
var loginFormPasswordSelector = "#password"
var loginFormSubmitButtonSelector = ".login-button"

// viewing history page selector
var viewingHistoryListSelector = "li.retableRow"
var viewingHistoryDateSelector = ".col.date.nowrap"
var viewingHistoryTitleSelector = ".col.title a"

// FetchViewingHistory is API to fetch Netflix viewing history by scraping
func (n *Netflix) FetchViewingHistory(email string, password string) ([]byte, error) {
	// / start login
	if err := n.Driver.Start(); err != nil {
		return nil, fmt.Errorf("failed to start driver:%v", err)
	}
	defer n.Driver.Stop()

	page, err := n.Driver.NewPage()

	if err != nil {
		return nil, fmt.Errorf("faild to open page:%v", err)
	}

	if err := page.Navigate(n.LoginURL); err != nil {
		return nil, fmt.Errorf("failed to navigate:%v", err)
	}

	id := page.Find(loginFormIDSelector)
	pass := page.Find(loginFormPasswordSelector)
	id.Fill(email)
	pass.Fill(password)

	if err := page.Find(loginFormSubmitButtonSelector).Submit(); err != nil {
		return nil, fmt.Errorf("failed to login:%v", err)
	}
	// end login

	// viewing history page
	if err := page.Navigate(n.ViewingHistoryURL); err != nil {
		return nil, fmt.Errorf("failed to navigate:%v", err)
	}

	html, err := page.HTML()

	if err != nil {
		return nil, fmt.Errorf("failed to get html:%v", err)
	}

	return []byte(html), err
}
