package server

import (
	"strconv"
)

// Subscription

type Subscription struct {
	ID     string
	Source *Source
}

type SubscriptionService struct {
	subscriptions []Subscription
}

func (s *SubscriptionService) Create(source *Source) *Subscription {
	sub := Subscription{
		ID:     strconv.Itoa(len(s.subscriptions) + 1),
		Source: source,
	}
	s.subscriptions = append(s.subscriptions, sub)
	return &sub
}

func (s *SubscriptionService) ListAll() []Subscription {
	return s.subscriptions
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

type Digest struct {
	Subscription Subscription
	Entries      []*Entry
}

type DigestService struct {
}

func (s *DigestService) Get(date string, subscriptionService *SubscriptionService) []Digest {
	digests := make([]Digest, 0)
	subscriptions := subscriptionService.ListAll()

	for i := range subscriptions {
		d := Digest{
			Subscription: subscriptions[i],
			Entries:      subscriptions[i].Source.EntriesForDate(date),
		}
		digests = append(digests, d)
	}

	return digests
}
