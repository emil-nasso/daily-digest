package plugins

import (
	"testing"

	"github.com/emil-nasso/daily-digest/util"
)

func TestScraper(t *testing.T) {
	//assert.Equal(t, "hello world", idgScraper())
	e := idgScraper()

	util.Dd(e)
	t.Fail()
}
