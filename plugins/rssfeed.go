package plugins

import "github.com/emil-nasso/daily-digest/daily_digest"

func init() {
	daily_digest.RegisterSource("rss", "RSS Feed", "Any RSS-feed. Great for blogs, podcasts and more.", []string{"blog", "podcast", "general"}, func() string {
		return ""
	})
}
