package plugins

import "github.com/emil-nasso/daily-digest/daily_digest"

func init() {
	daily_digest.RegisterSource("reddit", "Reddit", "The top posts of any subreddit.", []string{"news", "misc"}, func() string {
		return ""
	})
}
