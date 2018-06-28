package plugins

import "github.com/emil-nasso/daily-digest/server"

func init() {
	server.RegisterSource("rss", "RSS Feed", "Any RSS-feed. Great for blogs, podcasts and more.", []string{"blog", "podcast", "general"}, func() string {
		return ""
	})
}
