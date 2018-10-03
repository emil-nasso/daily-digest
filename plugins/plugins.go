package plugins

import (
	"os"

	"github.com/emil-nasso/daily-digest/daily_digest"
	"github.com/mmcdole/gofeed"
)

func noopPersister(e []*daily_digest.Entry) {

}

func getRssFeedItemsFromUrl(url string) []*gofeed.Item {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	return feed.Items
}

func getRssFeedItemsFromFile(fileName string) []*gofeed.Item {
	file, _ := os.Open(fileName)
	defer file.Close()
	fp := gofeed.NewParser()
	feed, _ := fp.Parse(file)
	return feed.Items
}

func rssItemsToEntriesUsingCurrentTime(items []*gofeed.Item) []*daily_digest.Entry {
	entires := make([]*daily_digest.Entry, 0)
	for _, item := range items {
		entires = append(entires, daily_digest.NewEntryForCurrentTime(
			item.Title,
			item.Description,
			item.Link,
		))
	}
	return entires
}
