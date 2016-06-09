package unmarshalsitemap

import (
	"github.com/remotejob/go_cv/domains"
	"encoding/xml"
//	"fmt"
	"io/ioutil"
	"time"
	//	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Get(sitemapfile string) []domains.SitemapObj {

	dat, err := ioutil.ReadFile(sitemapfile)
	check(err)

	var sitemap domains.Pages
	xml.Unmarshal(dat, &sitemap)

	var sitemapObjs []domains.SitemapObj

	for _, page := range sitemap.Pages {

		layout := "2006-01-02T15:04:05+03:00"
		t, err := time.Parse(layout, page.Lastmod)
		check(err)
		duration := time.Since(t)

		sitemapObj := domains.SitemapObj{page.Changefreq, duration.Hours(),page.Loc,page.Lastmod}
		
		sitemapObjs =append(sitemapObjs,sitemapObj)

	}

  return sitemapObjs
}
