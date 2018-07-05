package plugins

import "github.com/emil-nasso/daily-digest/daily_digest"

func init() {
	daily_digest.RegisterSource("svtnyheter", "SvtNyheter", "News in swedish", []string{"news", "swedish"}, func() string {
		return ""
	})
}
