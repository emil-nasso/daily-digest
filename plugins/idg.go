package plugins

import (
	"time"

	"github.com/emil-nasso/daily-digest/daily_digest"
	"github.com/mmcdole/gofeed"
)

var idgRssParser gofeed.Parser

var idgScraper = func() []*daily_digest.Entry {

	//items := getRssFeedItems("https://www.idg.se/rss/mest+l%C3%A4st%3A+senaste+dygnet")
	items := getRssFeedItemsFromFile("idg.xml")
	return rssItemsToEntriesUsingCurrentTime(items)
}

func init() {
	duration, _ := time.ParseDuration("5s")
	daily_digest.RegisterSource("idg", "IDG", "Tech news in swedish", []string{"news", "tech"}, idgScraper, duration, noopPersister)
}
