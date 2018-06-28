package server

import (
	"strconv"

	"github.com/emil-nasso/daily-digest/util"
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
	util.Dd(d)
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
	entries []Entry
}

func (s *EntryService) Seed() {

	s.entries = []Entry{
		Entry{
			ID:          "1",
			PublishedAt: "2018-01-01 10:00:01",
			Title:       "The russians are coming",
			Excerpt:     "Hide yoself, the ruskies are invading.",
			URL:         "https://www.svtnyheter.se/theruskies.html",
		},
		Entry{
			ID:          "2",
			PublishedAt: "2018-01-02 10:00:01",
			Title:       "The germans are coming",
			Excerpt:     "Hide yoself, the germans are invading.",
			URL:         "https://www.svtnyheter.se/zegermans.html",
		},
	}
}

// Daily

type Daily struct {
	Date    string
	Digests []*DailyDigest
}

type DailyDigest struct {
	Digest  *Digest
	Entries []*Entry
}
type DailiesService struct {
}

func (s *DailiesService) Get(date string, digestService *DigestService) Daily {
	dailyDigests := make([]*DailyDigest, 0)

	for _, digest := range digestService.ListAll() {
		//entries := digest.Source.
		entries := []*Entry{
			&Entry{
				ID:          "1",
				PublishedAt: "2018-01-01 10:00:01",
				Title:       "The russians are coming",
				Excerpt:     "Hide yoself, the ruskies are invading.",
				URL:         "https://www.svtnyheter.se/theruskies.html",
			},
			&Entry{
				ID:          "2",
				PublishedAt: "2018-01-02 10:00:01",
				Title:       "The germans are coming",
				Excerpt:     "Hide yoself, the germans are invading.",
				URL:         "https://www.svtnyheter.se/zegermans.html",
			},
		}
		dailyDigests = append(dailyDigests, &DailyDigest{
			Digest:  &digest,
			Entries: entries,
		})
	}

	return Daily{
		Date:    date,
		Digests: dailyDigests,
	}
}
