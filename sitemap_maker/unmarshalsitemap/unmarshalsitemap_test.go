package unmarshalsitemap

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	sitemapfile := "/home/juno/git/go_cv/version_00/maps/sitemap_127.0.0.1.xml"

	result := Get(sitemapfile)

	for _, obj := range result {

		fmt.Println(obj)
	}

}
