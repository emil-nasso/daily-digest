package daily_digest

import (
	"fmt"
	"time"
)

// Source represents a digest source
type Source struct {
	ID               string
	Name             string
	Description      string
	Tags             []string
	scraper          scraper
	scrapingInterval time.Duration
	persister        persister
	entries          []*Entry
}

type scraper func() []*Entry

type persister func([]*Entry)

var sources []*Source

func (s *Source) EntriesForDate(date string) []*Entry {
	entries := make([]*Entry, 0)
	for _, entry := range s.entries {
		if entry.PublishedAt == date {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (s *Source) AddEntry(entry *Entry) {
	s.entries = append(s.entries, entry)
}

// Get all initialized sources
func GetSources() []*Source {
	return sources
}

func GetSourceById(id string) *Source {
	for _, s := range sources {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func RegisterSource(id, name, description string, tags []string, scraper scraper, scrapingInterval time.Duration, persister persister) {
	fmt.Printf("Registering source: %s\n", id)
	source := &Source{
		ID:               id,
		Name:             name,
		Description:      description,
		Tags:             tags,
		scraper:          scraper,
		scrapingInterval: scrapingInterval,
		persister:        persister,
	}
	sources = append(sources, source)
}
