package main

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message"
	"github.com/google/uuid"
)

// GetRecentEntries will try to fetch numOfEntries of RSSEntry using the given IMAP config
func GetRecentEntries(config MailConfig, numOfEntries uint32) (map[uint32]*RSSEntry, error) {
	rssMap := make(map[uint32]*RSSEntry, numOfEntries)

	// Connect to server
	c, err := client.DialTLS(config.IMAPServer, nil)
	if err != nil {
		return rssMap, err
	}
	defer c.Logout()

	// Login
	if err := c.Login(config.Username, config.Password); err != nil {
		return rssMap, err
	}

	mbox, err := c.Select(config.IMAPInboxName, false)
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > numOfEntries {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - numOfEntries
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, numOfEntries)
	done := make(chan error, 1)
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	for msg := range messages {
		r := msg.GetBody(section)
		if r == nil {
			log.Fatal("Server didn't returned message body")
		}

		m, err := message.Read(r)
		if message.IsUnknownCharset(err) {
			// This error is not fatal
			log.Println("Unknown encoding:", err)
		} else if err != nil {
			return rssMap, err
		}

		if mr := m.MultipartReader(); mr != nil {
			// This is a multipart message
			//log.Println("This is a multipart message containing:")
			messageMap := make(map[string]string, 3)

			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					return rssMap, err
				}

				t, _, _ := p.Header.ContentType()
				//log.Println("A part with type", t)
				if t == "text/html" || t == "text/plain" || t == "multipart/alternative" {
					body, _ := ioutil.ReadAll(p.Body)
					messageMap[t] = string(body)
				}
			}
			if len(messageMap) > 0 {
				rssMap[msg.SeqNum] = &RSSEntry{Messages: messageMap, seqNum: msg.SeqNum}
			}

		} else {
			t, _, _ := m.Header.ContentType()
			log.Println("This is a non-multipart message with type", t)
		}

	}

	if err := <-done; err != nil {
		return rssMap, err
	}

	for _, rss := range rssMap {
		seqsetNew := new(imap.SeqSet)
		seqsetNew.AddNum(rss.seqNum)
		messageHeaders := make(chan *imap.Message, 1)

		err := c.Fetch(seqsetNew, []imap.FetchItem{imap.FetchEnvelope}, messageHeaders)
		if err != nil {
			return rssMap, err
		}

		mh := <-messageHeaders
		rss.Subject = mh.Envelope.Subject
		rss.Author = Author{
			MailboxName: mh.Envelope.From[0].MailboxName,
			Domain:      mh.Envelope.From[0].HostName,
			Name:        mh.Envelope.From[0].PersonalName,
		}
		rss.Date = mh.Envelope.Date
		rss.ID = uuid.NewMD5(config.Secret, []byte(mh.Envelope.MessageId)).String()
	}

	return rssMap, nil
}
