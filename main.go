package main

import (
	"net/http"
	"os"

	"github.com/mmcdole/gofeed"
)

type Post struct {
	Body string
	Url string
	Title string
}

func main() {
	// Check for API key
	key := os.Getenv("BUTTONDOWN_KEY")
	if key == "" {
		log.Panicf("[BUTTONDOWN] Missing Buttondown API key. See https://buttondown.email/settings/programming to get one.")
	}

	// Check RSS feed
	var newPosts = make(map[int]Post)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(os.Args[1])
	if err != nil {
		log.Panicf("[FEED] Could not parse %v: %v", os.Args[1], err)
	}
	for post, i := range feed.Items {
		// New posts will be published at the same time as
		// the feed usually. There are some outstanding conditions,
		// but this can usually be avoided by.. y'know.. updating your
		// feed and this tool sooner than later.
		//
		// TODO: find a better way to handle this.
		if post.Published == feed.Published {
			// This is a new post (yay!)
			// Add it to the database to be sent on its way
			newPosts[i].Body = post.Content
			newPosts[i].Url = post.Link
			newPosts[i].Title = post.Title
		}
	}

	// Send emails
	direct := os.Getenv("COURIER_DIRECT")
	if direct == true {
		log.Printf("[BUTTONDOWN] Sending emails without human checks.")
		for post, _ := range newPosts {
			resp, err := http.PostForm("https://api.buttondown.email/v1/emails", url.Values{
				"body": post.Body,
				"email_type": "public", // ????
				"external_url": post.Url,
				"subject": post.Title,
			})
			if err != nil {
				log.Panicf("[BUTTONDOWN] Unknown error trying to create emails: %v", err)
			}
		}
		// Send emails
	} else {
		for post, _ := range newPosts {
			resp, err := http.PostForm("https://api.buttondown.email/v1/drafts", url.Values{
				"body": post.Body,
				"subject": post.Title,
			})
			if err != nil {
				log.Panicf("[BUTTONDOWN] Unknown error trying to create drafts: %v", err)
			}
		}
		// Store email instead of sending
	}
}