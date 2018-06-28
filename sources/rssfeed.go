package sources

func init() {
	registerSource("rss", "RSS Feed", "Any RSS-feed. Great for blogs, podcasts and more.", []string{"blog", "podcast", "general"}, func() string {
		return ""
	})
}
