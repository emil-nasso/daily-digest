package sources

// Source represents a digest source
type Source struct {
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

func registerSource(name, description string, tags []string, scraper scraper) {
	source := Source{
		Name:        name,
		Description: description,
		Tags:        tags,
		scraper:     scraper,
	}
	sources = append(sources, source)
}
