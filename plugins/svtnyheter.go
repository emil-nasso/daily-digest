package plugins

import "github.com/emil-nasso/daily-digest/server"

func init() {
	server.RegisterSource("svtnyheter", "SvtNyheter", "News in swedish", []string{"news", "swedish"}, func() string {
		return ""
	})
}
