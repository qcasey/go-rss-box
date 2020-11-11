package main

import (
	"fmt"
	"time"

	"github.com/gorilla/feeds"
)

type Author struct {
	Name        string
	MailboxName string
	Domain      string
}

// RSSEntry corresponds to an email that should be parsed and returned in a feed
type RSSEntry struct {
	Messages map[string]string
	Subject  string
	Author   Author
	Date     time.Time
	ID       string

	seqNum uint32
}

func CreateFeed() (*feeds.Feed, error) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       config.FeedName,
		Link:        &feeds.Link{Href: config.ServerHostname},
		Description: config.FeedDescription,
		Created:     now,
	}

	entries, err := GetRecentEntries(config.Mail, 10)
	if err != nil {
		return feed, err
	}

	for _, entry := range entries {
		feed.Items = append(feed.Items, &feeds.Item{
			Id:      entry.ID,
			Title:   entry.Subject,
			Link:    &feeds.Link{Href: fmt.Sprintf("%s/feed/%s", config.ServerHostname, entry.ID)},
			Content: entry.Messages["text/html"],
			Author:  &feeds.Author{Name: entry.Author.Name, Email: entry.Author.Domain},
			Created: entry.Date,
		})
	}

	return feed, nil
}
