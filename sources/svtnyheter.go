package sources

func init() {
	registerSource("SvtNyheter", "News in swedish", []string{"news", "swedish"}, func() string {
		return ""
	})
}
