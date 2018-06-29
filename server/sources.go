package server

import (
	"fmt"
)

// Source represents a digest source
type Source struct {
	ID          string
	Name        string
	Description string
	Tags        []string
	scraper     scraper
	entries     []*Entry
}

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

type scraper func() string

var sources []*Source

// Get all initialized sources
func Get() []*Source {
	return sources
}

func GetById(id string) *Source {
	for _, s := range sources {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func RegisterSource(id, name, description string, tags []string, scraper scraper) {
	fmt.Printf("Registering source: %s\n", id)
	source := &Source{
		ID:          id,
		Name:        name,
		Description: description,
		Tags:        tags,
		scraper:     scraper,
	}
	sources = append(sources, source)
}
