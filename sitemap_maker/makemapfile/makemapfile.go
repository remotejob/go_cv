package makemapfile

import (
	"github.com/remotejob/go_cv/domains"
	"encoding/xml"
//	"math/rand"
//	"time"
	"log"
	"io/ioutil"	
)

func Makefile(filestr string, sitemapobjs []domains.SitemapObj) {

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for _, sitemapobj := range sitemapobjs {

		doc := new(domains.Page)
		doc.Loc = sitemapobj.Loc
		//		now := time.Now()
		//		intrand := random(100, 50000)
		//		minback := time.Duration(intrand)
		//		lastmod := now.Add(-minback * time.Second)
		doc.Lastmod = sitemapobj.Lastmod
		doc.Changefreq = sitemapobj.Changefreq
		docList.Pages = append(docList.Pages, doc)

	}

	resultXml, err := xml.MarshalIndent(docList, "", "  ")
	if err != nil {

//		golog.Crit(err.Error())
		log.Println(err.Error())
	}
	docList.Pages = nil

	ioutil.WriteFile(filestr, resultXml, 0644)
	if err != nil {

//		golog.Crit(err.Error())
		log.Println(err.Error())		
	}

}
