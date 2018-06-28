package sources

func init() {
	registerSource("svtnyheter", "SvtNyheter", "News in swedish", []string{"news", "swedish"}, func() string {
		return ""
	})
}
