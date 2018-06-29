package server

import (
	"strconv"
)

// Digest

type Digest struct {
	ID     string
	Source *Source
}

type DigestService struct {
	digests []Digest
}

func (s *DigestService) Create(source *Source) *Digest {
	d := Digest{
		ID:     strconv.Itoa(len(s.digests) + 1),
		Source: source,
	}
	s.digests = append(s.digests, d)
	return &d
}

func (s *DigestService) ListAll() []Digest {
	return s.digests
}

// Entry

type Entry struct {
	ID          string
	PublishedAt string
	Title       string
	Excerpt     string
	URL         string
}

type EntryService struct {
}

func (s *EntryService) Seed() {

	rss := GetById("rss")
	svt := GetById("svtnyheter")

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

// Daily

type Daily struct {
	Date    string
	Digests []DailyDigest
}

type DailyDigest struct {
	Digest  Digest
	Entries []*Entry
}

type DailiesService struct {
}

func (s *DailiesService) Get(date string, digestService *DigestService) Daily {
	dailyDigests := make([]DailyDigest, 0)
	digests := digestService.ListAll()

	for i := range digests {
		d := DailyDigest{
			Digest:  digests[i],
			Entries: digests[i].Source.EntriesForDate(date),
		}
		dailyDigests = append(dailyDigests, d)
	}

	return Daily{
		Date:    date,
		Digests: dailyDigests,
	}
}
