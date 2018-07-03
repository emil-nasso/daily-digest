package plugins

import "github.com/emil-nasso/daily-digest/server"

func init() {
	server.RegisterSource("reddit", "Reddit", "The top posts of any subreddit.", []string{"news", "misc"}, func() string {
		return ""
	})
}
