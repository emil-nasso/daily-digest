package sources

import "fmt"

// Source represents a digest source
type Source struct {
	ID          string
	Name        string
	Description string
	Tags        []string
	scraper     scraper
}

type scraper func() string

var sources []Source

// Get all initialized sources
func Get() []Source {
	return sources
}

func GetById(id string) *Source {
	for _, s := range sources {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

func registerSource(id, name, description string, tags []string, scraper scraper) {
	fmt.Printf("Registering source: %s\n", id)
	source := Source{
		ID:          id,
		Name:        name,
		Description: description,
		Tags:        tags,
		scraper:     scraper,
	}
	sources = append(sources, source)
}
