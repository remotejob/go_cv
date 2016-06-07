package formfeeder

import (
	"github.com/remotejob/go_cv/domains"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
	"path/filepath"
	//	"strings"
	"github.com/remotejob/go_cv/blogfeeder/addlink"
	"github.com/remotejob/go_cv/blogfeeder/addnewblogitem"
	"github.com/remotejob/go_cv/blogfeeder/check_title"
//	"blogfeeder/check_topic"
	"github.com/gosimple/slug"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func HandleForm(c web.C, w http.ResponseWriter, r *http.Request) {
	rootdir := c.Env["rootdir"].(string)
	locale := c.Env["locale"].(string)
	themes := c.Env["themes"].(string)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/feederform.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		topic := r.Form["topic"][0]
		title := r.Form["title"][0]
		contents := r.Form["contents"][0]

		if (rootdir != "") && (locale != "") && (themes != "") && (topic != "") && (title != "") && (contents != "") {

			now := time.Now()
			filestr := filepath.Join(rootdir, "dist", locale+"_"+themes+"_"+"blog.json")

			linksdir := filepath.Join(rootdir, "links")

			stitle := slug.Make(title)
			stopic := slug.Make(topic)
			
			blogItems := make(map[string][]domains.BlogItem)
			item := domains.BlogItem{stopic,topic,stitle, title, contents, now, now}

			if _, err := os.Stat(filestr); os.IsNotExist(err) {

				addnewblogitem.Addnew(blogItems, item, stopic, true, linksdir, filestr)

			} else {

				var blogObj map[string]*json.RawMessage
				file, e := ioutil.ReadFile(filestr)
				if e != nil {
					fmt.Printf("File error: %v\n", e)
					os.Exit(1)
				}
				err := json.Unmarshal(file, &blogObj)
				if err != nil {
					fmt.Println("error:", err)
				} else {

					for keytopic, val := range blogObj {

						var items []domains.BlogItem
						err := json.Unmarshal(*val, &items)
						if err != nil {
							fmt.Println("error:", err)
						} else {

							blogItems[keytopic] = items

						}

					}
					
					key :=stopic				
					stitleOK := check_title.CheckIfExist(stitle, blogItems[key])
//					topicOK := check_topic.CheckTopicExist(topic, blogItems[key])

					if !stitleOK {

						blogItems[key] = append(blogItems[key], item)

						b, err := json.Marshal(blogItems)
						if err != nil {
							log.Println(err)
						}
						ioutil.WriteFile(filestr, b, 0644)

						addlink.AddLinktoAllfiles(linksdir, stopic, stitle)

					}

				}

			}

		}
		http.Redirect(w, r, "/formfeeder", http.StatusFound)
	}

}
