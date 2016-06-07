package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/zenazn/goji/web"
	"github.com/remotejob/go_cv/godevice"
	"io"
	"net/http"
	"os"
	"strings"
)

type Query struct {
	XMLName xml.Name `xml:"urlset"`
	Locs    []string `xml:"url>loc"`
}
type Loc string

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func Elaborate(c web.C, w http.ResponseWriter, r *http.Request) {

	rootdir := c.Env["rootdir"].(string)
	rootdirm := c.Env["rootdirm"].(string)	
	backendrootdir := c.Env["backendrootdir"].(string)

	exist := false
	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]
	path := r.URL.Path
	urlstr := "http://" + site + path

	fmt.Println(urlstr)
	deviceType := godevice.GetType(r)
	fmt.Println(deviceType)

	sitemapfile := backendrootdir + "/maps/sitemap_" + site + ".xml"
	//	fmt.Println("map", sitemapfile)

	if _, err := os.Stat(sitemapfile); os.IsNotExist(err) {

		fmt.Println("NOT EXIST")
		// path/to/whatever does not exist
	} else {

		if strings.HasSuffix(path, ".json") {

			fmt.Println(rootdir + "/dist" + path)
			if deviceType == "Mobile" {
				http.ServeFile(w, r, rootdirm+"/dist"+path)
			} else {
				http.ServeFile(w, r, rootdir+"/dist"+path)
			}

		} else {

			if path == "/" || path == "/index.html" {

				exist = true

			} else {

				f, _ := os.Open(sitemapfile)
				buf := bytes.NewBuffer(nil)
				io.Copy(buf, f)

				var q Query
				xml.Unmarshal(buf.Bytes(), &q)

				if stringInSlice(urlstr, q.Locs) {

					exist = true
				}
			}

		}
	}

	if !exist {

//		fmt.Println("return 404")
		http.NotFound(w, r)

	} else {

		if deviceType == "Mobile" {

			http.ServeFile(w, r, "/home/juno/git/cv/version_mobile_react_00/dist/index.html")

		} else {
			http.ServeFile(w, r, rootdir+"/dist/index.html")
		}

	}


}
