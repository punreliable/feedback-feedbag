package utils

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Login(username, password, url string) (*rod.Browser, *rod.Page) {
	dir := "~/.config/google-chrome"

	u := launcher.New().
		UserDataDir(dir).
		Leakless(true).
		NoSandbox(true).
		Headless(false).
		Set("js-flags", "--max_old_space_size=65536").
		Set("disable-gpu", "true").
		Set("disable-software-rasterizer", "true").
		Set("disable-site-isolation-trials", "true").
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect().NoDefaultDevice()
	page := browser.MustPage(url).MustWaitLoad().MustWindowMaximize()

	if page.MustHasX("/html/body/div[1]/div/div[1]/div/div[5]/div/div/div[1]/div/div[2]/div/div/div") {
		page.MustElementX("/html/body/div[1]/div/div[1]/div/div[5]/div/div/div[1]/div/div[2]/div/div/div/div[2]/form/div/div[3]/div/div/label/div/input").MustInput(username)
		page.MustElementX("/html/body/div[1]/div/div[1]/div/div[5]/div/div/div[1]/div/div[2]/div/div/div/div[2]/form/div/div[4]/div/div/label/div/input").MustInput(password)
		page.MustElementX("/html/body/div[1]/div/div[1]/div/div[5]/div/div/div[1]/div/div[2]/div/div/div/div[2]/form/div/div[5]/div").MustClick()
		time.Sleep(1 * time.Minute)
		page.MustNavigate(url).MustWaitLoad().MustWaitDOMStable()
	}

	return browser, page
}
