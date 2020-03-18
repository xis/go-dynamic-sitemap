package sitemap

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	var s Sitemap
	err := Create("www.example.com/path/", []string{"userid", "contentid", "something", "somethingelse"}).Result(&s)
	if err != nil {
		panic(err)
	}
}
