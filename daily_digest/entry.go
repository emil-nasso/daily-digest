package daily_digest

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Entry struct {
	ID          string
	PublishedAt string
	Title       string
	Excerpt     string
	URL         string
}

func generateUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func NewEntryForCurrentTime(title, description, url string) *Entry {
	currentDate := time.Now().Format("2006-01-02")
	return &Entry{
		ID:          generateUUID(),
		PublishedAt: currentDate,
		Title:       title,
		Excerpt:     description,
		URL:         url,
	}
}

func SeedEntries() {

	rss := GetSourceById("rss")
	svt := GetSourceById("svtnyheter")

	svt.AddEntry(&Entry{
		ID:          "1",
		PublishedAt: "2018-01-01",
		Title:       "The russians are coming",
		Excerpt:     "Hide yoself, the russains are invading.",
		URL:         "https://www.svtnyheter.se/theruskies.html",
	})

	svt.AddEntry(&Entry{
		ID:          "2",
		PublishedAt: "2018-01-01",
		Title:       "The russians are retreated",
		Excerpt:     "Hide yoself, the russian where just passing by.",
		URL:         "https://www.svtnyheter.se/wearesafe.html",
	})

	svt.AddEntry(&Entry{
		ID:          "3",
		PublishedAt: "2018-01-02",
		Title:       "The germans are coming",
		Excerpt:     "Hide yoself, the germans are invading.",
		URL:         "https://www.svtnyheter.se/germans.html",
	})

	rss.AddEntry(&Entry{
		ID:          "4",
		PublishedAt: "2018-01-01",
		Title:       "Monday",
		Excerpt:     "First day of the week!",
		URL:         "http://www.example.com/blog/1.html",
	})

	rss.AddEntry(&Entry{
		ID:          "5",
		PublishedAt: "2018-01-02",
		Title:       "Tuesday",
		Excerpt:     "Second day of the week!",
		URL:         "http://www.example.com/blog/2.html",
	})

	rss.AddEntry(&Entry{
		ID:          "6",
		PublishedAt: "2018-01-03",
		Title:       "Wednesday",
		Excerpt:     "Third day of the week!",
		URL:         "http://www.example.com/blog/3.html",
	})

}
