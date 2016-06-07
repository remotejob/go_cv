package main

import (
	"database/sql"
	"github.com/remotejob/go_cv/domains"
//	"flag"
	"fmt"
	_ "github.com/mxk/go-sqlite/sqlite3"
	//	"io/ioutil"
	"gopkg.in/gcfg.v1"	
	"log"
	"log/syslog"
	"github.com/remotejob/go_cv/mark/dbgetall"
	"math/rand"
	"net/url"
	"path/filepath"
	"github.com/remotejob/go_cv/sitemap_maker/contents_feeder"
	"github.com/remotejob/go_cv/sitemap_maker/getLinks"
	"github.com/remotejob/go_cv/sitemap_maker/makemapfile"
	"github.com/remotejob/go_cv/sitemap_maker/unmarshalsitemap"
	"time"
)

//var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
//var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
//var dblocFlag = flag.String("dbloc", "", "must be somthing like en_US_programming.db")
//var contentsdirFlag = flag.String("contentsdir", "", "must dir location contents files")
//var linksdirFlag = flag.String("linksdir", "", "must dir location links files")
//var mapsdirFlag = flag.String("mapsdir", "", "must dir location sitemaps files")

var rootdir = ""
var backendrootdir = ""
var locale = ""
var themes = ""
var dbdir = ""
var changefreq = ""


func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
func init() {
	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		rootdir = cfg.Dirs.Rootdir
		locale = cfg.Main.Locale
		themes = cfg.Main.Themes
		backendrootdir = cfg.Dirs.Backendrootdir
		dbdir = cfg.Dirs.Dbdir
		changefreq =cfg.Main.Changefreq 

	}
}
func main() {

	linksdir := rootdir+"/links"
	mapsdir := backendrootdir+"/maps"
	contentsdir := rootdir+"/dist/www"
	dbloc := dbdir

	if (linksdir != "") && (mapsdir != "") && (contentsdir != "") && (locale != "") && (themes != "") && (dbloc != "") {
		golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}
		//

		db, err := sql.Open("sqlite3", dbloc)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		phrases := dbgetall.GetAll(*db, locale, themes, "phrases", "phrase")
		keywords := dbgetall.GetAll(*db, locale, themes, "keywords", "keyword")

		linksmap := getLinks.GetAllLinks(*golog, linksdir)


		var site string

		for key, vals := range linksmap {

			site = key
			filestr := mapsdir + "/sitemap_" + site + ".xml"

			fmt.Println("Site", site, filestr)

			sitemapObjs := unmarshalsitemap.Get(filestr)

			uniqlinkmap := make(map[string]struct{})

			var newlinks []string
			var durationfixed float64
			durationfixed = float64(0)

			for i, sitemapObj := range sitemapObjs {

				uniqlinkmap[sitemapObj.Loc] = struct{}{}

				if sitemapObj.Changefreq == "monthly" {

					durationfixed = float64(720)

				}
				
				if sitemapObj.Changefreq == "weekly" {

					durationfixed = float64(168)

				}

				if sitemapObj.Hoursduration > durationfixed {

					fmt.Println("need update ", sitemapObj.Loc)
					sitemapObjs[i].Lastmod = time.Now().Format(time.RFC3339)
					sitemapObjs[i].Changefreq = changefreq
					u, err := url.Parse(sitemapObj.Loc)
					if err != nil {
						panic(err)
					}
					path := u.Path

					contents_feeder.MakeContents(filepath.Join(contentsdir, site), path, keywords, phrases)
				}

			}

			for _, link := range vals {

				pageurl := "http://" + site + link

				if _, ok := uniqlinkmap[pageurl]; !ok {

					fmt.Println("new link", pageurl)
					newlinks = append(newlinks, pageurl)

				}


			}

			for _, link := range newlinks {

				newsitemapObj := domains.SitemapObj{changefreq, float64(0), link, time.Now().Format(time.RFC3339)}
				sitemapObjs = append(sitemapObjs, newsitemapObj)
				u, err := url.Parse(link)
				if err != nil {
					panic(err)
				}
				path := u.Path
				
				contents_feeder.MakeContents(filepath.Join(contentsdir, site), path, keywords, phrases)

			}

			makemapfile.Makefile(filestr, sitemapObjs)

		}

	} else {
		fmt.Println("check config.gcfg")
	}

}
