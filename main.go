package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/xuri/excelize/v2"
)

type Comment struct {
	CommentText          string
	CommentorName        string
	CommentorProfileLink string
}

func main() {

	fmt.Print("Here we go...\n")
	// postURL := "https://web.facebook.com/photo/?fbid=1298842747871125&set=a.835084834246921"
	//_, page := utils.Login("", "", postURL)
	//chromiumPath := "/usr/bin/google-chrome"
	chromiumPath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	postURL := "https://www.facebook.com/100064306171261/posts/pfbid02rhBKwS16aMJyqbTyFNqfrfcJD3cdKjCV93NTTLeSJMgMAA3FFvtFSpf5m3os1VWNl"
	l := launcher.New().
		Bin(chromiumPath).
		Headless(false).
		Devtools(false).
		Set("js-flags", "--max_old_space_size=65536").
		Set("disable-gpu", "true").
		Set("disable-software-rasterizer", "true").
		Set("disable-site-isolation-trials", "true")

	// fmt.Print("...I need to start a Browser...\n")
	browser := rod.New().ControlURL(l.MustLaunch()).MustConnect().NoDefaultDevice()
	// fmt.Print("...ok, Browser is running...\n")
	page := browser.MustPage(postURL).MustWindowMaximize().MustWaitLoad().MustWaitDOMStable()
	// fmt.Print("...ok, Page is running...\n")
	page.MustElement(`div.x6s0dn4.x78zum5.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.xe0p6wg`).MustClick()
	// fmt.Print("...pfeewww I'm sleepy, boy...\n")
	time.Sleep(3 * time.Second)
	// fmt.Print("...I just slept 3 seconds...\n")
	menuItems := page.MustElements(`div[role="menuitem"]`)
	for _, menuItem := range menuItems {
		text := menuItem.MustText()
		if strings.Contains(text, "All comments") {
			menuItem.MustClick()
		}
	}
	// fmt.Print("...pfeewww I'm sleepy, boy...\n")
	time.Sleep(3 * time.Second)
	// fmt.Print("...I just slept for 3 seconds...\n")

	// fmt.Print("...Let's do replies...\n")
	maxReplyIterations := 20
	replyIter := 0
	for {
		if replyIter >= maxReplyIterations {
			fmt.Println("Max reply iterations reached")
			break
		}
		replyIter++

		commentReplies := page.MustElements(`div.x1i10hfl.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x2lwn1j.xeuugli.xexx8yu.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x3nfvp2.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xi81zsa.x1q0g3np.x1iyjqo2.xs83m0k.xsyo7zv.x1mnrxsn`)
		totalReplies := len(commentReplies)

		fmt.Println("Total replies:", totalReplies)

		if totalReplies == 0 {
			fmt.Println("No more replies")
			break
		}

		fmt.Println("Handle replies")

		for range totalReplies {
			page.MustEval(`() => document.querySelector('div.x1i10hfl.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x2lwn1j.xeuugli.xexx8yu.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x3nfvp2.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xi81zsa.x1q0g3np.x1iyjqo2.xs83m0k.xsyo7zv.x1mnrxsn').click()`)
			fmt.Println("Replies button clicked")
		}

		commentReplies = nil

		time.Sleep(10 * time.Second)
	}

	fmt.Print("...About to print the number of articles...")
	fmt.Print("\n")
	fmt.Print(len(page.MustElements(`div[role="article"]`)))
	fmt.Print("\n")

	// complementary := page.MustElement(`div[role="complementary"]`)
	// complementary.MustClick()

	fmt.Print("Max Iterations set to 10, starting now...\n")
	maxIterations := 10
outer:

	for i := 0; i < maxIterations; i++ {

		// fmt.Print("...beginning loop...\n")

		pageHasViewMoreCommentsButton := page.MustHas(`div.x1i10hfl.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x2lwn1j.xeuugli.xexx8yu.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x3nfvp2.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xi81zsa.x1q0g3np.x1iyjqo2.xs83m0k.xsyo7zv`)

		// fmt.Print("...I wonder if the pageHasViewMoreCommentsButton is displayed...\n")
		if pageHasViewMoreCommentsButton {

			// fmt.Print("\n")
			// fmt.Print("Page has view more comments button...")
			// fmt.Print("\n")

			timeout := time.After(1 * time.Minute)
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-timeout:
					fmt.Print("\n")
					fmt.Print("Timeout reached! Exiting loop.")
					fmt.Print("\n")
					break outer
				case <-ticker.C:
					//loadingComments := page.MustHas(`span.html-span.xdj266r.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x78zum5.x1w0mnb.xeuugli`)
					loadingComments := page.MustHas(`span.xdj266r.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1hl2dhg.x16tdsg8.x1vvkbs.x78zum5.x1w0mnb.xeuugli`)
					if loadingComments {
						fmt.Print("\n")
						fmt.Print("Loading comments...")
						fmt.Print("\n")
						time.Sleep(20 * time.Millisecond)
						continue
					}
					fmt.Print("Comments loaded.")
				}
				break
			}
			// Wait for the button to be clickable

			// Block image requests
			err := proto.NetworkSetBlockedURLs{
				Urls: []string{"*.jpg", "*.jpeg", "*.png", "*.gif", "*.svg", "*.webp"}, // Add other image extensions as needed
			}.Call(page)

			if err != nil {
				log.Fatal(err)
			}

			page.MustEval(`() => document.querySelector('div.x1i10hfl.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x2lwn1j.xeuugli.xexx8yu.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x3nfvp2.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xi81zsa.x1q0g3np.x1iyjqo2.xs83m0k.xsyo7zv').click()`)

			// fmt.Print("\n")
			// fmt.Print("View more comments button clicked")
			// fmt.Print("\n")

			timeout2 := time.After(3 * time.Minute)
			ticker2 := time.NewTicker(1 * time.Second)
			defer ticker2.Stop()

			for {
				select {
				case <-timeout2:
					fmt.Print("\n")
					fmt.Print("Timeout reached! Exiting loop.")
					fmt.Print("\n")
				case <-ticker2.C:
					loadingComments := page.MustHas(`div.html-div.xdj266r.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x78zum5.x1w0mnb.xeuugli`)
					if loadingComments {
						fmt.Print("\n")
						fmt.Print("Loading comments...")
						fmt.Print("\n")
						time.Sleep(20 * time.Millisecond)
						continue
					}
					fmt.Print("\n")
					fmt.Print("Comments loaded.")
					fmt.Print("\n")
				}

				break
			}

			if page.MustHas(`div.html-div.xdj266r.xat24cr.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x78zum5.x13a6bvl.x1d52u69.xktsk01`) {
				fmt.Print(page.MustElement(`div.html-div.xdj266r.xat24cr.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x78zum5.x13a6bvl.x1d52u69.xktsk01`).MustText())
			}

		} else {
			fmt.Print("Page does not have view more comments button")
			break
		}

		for {
			commentReplies := page.MustElements(`div.x1i10hfl.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x2lwn1j.xeuugli.xexx8yu.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x3nfvp2.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xi81zsa.x1q0g3np.x1iyjqo2.xs83m0k.xsyo7zv.x1mnrxsn`)

			totalReplies := len(commentReplies)
			fmt.Print("\n")
			fmt.Print("Total replies:", totalReplies)
			fmt.Print("\n")
			if totalReplies == 0 {
				fmt.Print("\n")
				fmt.Print("No more replies")
				fmt.Print("\n")
				break
			}
			fmt.Print("\n")
			fmt.Print("Handle replies")
			fmt.Print("\n")

			for range totalReplies {
				page.MustEval(`() => document.querySelector('div.x1i10hfl.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x2lwn1j.xeuugli.xexx8yu.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x3nfvp2.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xi81zsa.x1q0g3np.x1iyjqo2.xs83m0k.xsyo7zv.x1mnrxsn').click()`)
				fmt.Print("\n")
				fmt.Print("Replies button clicked")
				fmt.Print("\n")
			}

			commentReplies = nil

			time.Sleep(10 * time.Second)
		}

		fmt.Print("\n")
		fmt.Print(len(page.MustElements(`div[role="article"]`)))
		fmt.Print("\n")
	}

	articles := page.MustElements(`div[role="article"]`)

	fmt.Print(len(articles))
	fmt.Print("\n")
	fmt.Print(len(articles))
	fmt.Print("\n")

	var comments []Comment

	fmt.Print("...About to start another loop...\n")

	for _, article := range articles {
		if !article.MustHas(`div.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x1vvkbs`) {
			continue
		}

		re := regexp.MustCompile(`[\r\n]+`)

		commentText := strings.TrimSpace(re.ReplaceAllString(article.MustElement(`div.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x1vvkbs`).MustText(), " "))
		fmt.Println(commentText)

		if commentText == "" {
			fmt.Println("Comment is empty")
			continue
		}

		a, err := article.Element(`a.x1i10hfl.xjbqb8w.x1ejq31n.xd10rxx.x1sy0etr.x17r0tee.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xt0psk2.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x16tdsg8.x1hl2dhg.xggy1nq.x1a2a7pz.x1heor9g.xkrqix3.x1sur9pj.x1s688f`)
		if err != nil || a == nil {
			fmt.Println("...No commenter name found, skipping...")
			continue
		}

		commentorName := strings.TrimSpace(a.MustText())

		// a := article.MustElement(`a.x1i10hfl.xjbqb8w.x1ejq31n.xd10rxx.x1sy0etr.x17r0tee.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xt0psk2.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x16tdsg8.x1hl2dhg.xggy1nq.x1a2a7pz.x1heor9g.xkrqix3.x1sur9pj.x1s688f`)
		//commentorName := strings.TrimSpace(a.MustText())
		//fmt.Println(commentorName)

		href := *a.MustAttribute(`href`)
		href = fmt.Sprintf(`%s`, href)
		//fmt.Println(href)
		// p := browser.MustPage(href).MustWaitLoad()
		// a = p.MustElement(`a[aria-label="View profile"]`)
		// href = *a.MustAttribute(`href`)
		// p.MustNavigate(href).MustWaitLoad()
		// href = p.MustInfo().URL
		// commentorProfileLink := href
		// fmt.Println(commentorProfileLink)
		// p.MustClose()

		comments = append(comments, Comment{
			CommentText:          commentText,
			CommentorName:        commentorName,
			CommentorProfileLink: href,
		})
	}

	fmt.Print("...I'm about to attempt exporting to Excel...\n")
	if err := exportCommentsToExcel(comments, "comments.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func scrollToBottomByKeyboard(page *rod.Page) error {
	// Use Page.Keyboard.Press("End") to scroll to the end of the page using the "End" key
	err := page.Keyboard.Press(input.End)
	return err
}

func clearCache(page *rod.Page) {
	err := proto.NetworkClearBrowserCache{}.Call(page)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Browser cache cleared")
}

func removeComments(page *rod.Page) {
	elements, err := page.Elements(`div.x16hk5td.x12rz0ws`)
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range elements {
		if err := element.Remove(); err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Removed an element")
	}

	fmt.Println("Comments removed")

	clearCache(page)
}

func exportCommentsToExcel(comments []Comment, filename string) error {
	const sheetName = "Comments"

	// Create a new Excel file
	file := excelize.NewFile()

	// Create a new sheet and set the sheet name
	index, err := file.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("error creating new sheet: %v", err)
	}

	// Define a style for bold text
	boldStyle, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err != nil {
		return fmt.Errorf("error creating bold style: %v", err)
	}

	// Set headers for the columns (UPPERCASE and BOLD)
	headers := []string{"Comment Text", "Commentor Name", "Commentor Profile Link"}
	for col, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+col) // A1, B1, C1
		file.SetCellValue(sheetName, cell, strings.ToUpper(header))
		err = file.SetCellStyle(sheetName, cell, cell, boldStyle)
		if err != nil {
			return fmt.Errorf("error setting style for cell %s: %v", cell, err)
		}
	}

	// Fill the sheet with data from the comments
	for i, comment := range comments {
		row := i + 2 // Starting from row 2
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), comment.CommentText)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), comment.CommentorName)

		// Make the URL clickable using the AddHyperlink function
		cell := fmt.Sprintf("C%d", row)
		linkValue := comment.CommentorProfileLink // Use the actual link as the value to display
		if comment.CommentorProfileLink != "" {
			err := file.SetCellHyperLink(sheetName, cell, comment.CommentorProfileLink, "External", excelize.HyperlinkOpts{
				Display: &linkValue,
			})
			if err != nil {
				return fmt.Errorf("error setting hyperlink for cell %s: %v", cell, err)
			}

			// Use the original link as the displayed value in the cell
			file.SetCellValue(sheetName, cell, linkValue) // Display the actual link
		} else {
			file.SetCellValue(sheetName, cell, "") // or some placeholder
		}

	}

	// Adjust column width for better readability
	file.SetColWidth(sheetName, "A", "C", 35)  // Increased width for better link visibility
	file.SetColWidth(sheetName, "C", "C", 100) // Increase column C width

	// Set the active sheet
	file.SetActiveSheet(index)

	// Save the file
	if err := file.SaveAs(filename); err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	fmt.Printf("Excel file saved as %s\n", filename)
	return nil
}
