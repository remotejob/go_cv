package addnewblogitem

import (
	"github.com/remotejob/go_cv/blogfeeder/addlink"
	"github.com/remotejob/go_cv/domains"
	"encoding/json"
	"io/ioutil"
	"log"
)

func Addnew(blogItems map[string][]domains.BlogItem, item domains.BlogItem, stopic string,topicOK bool,linksdir string, filestr string) {

	key :=stopic
	 
	blogItems[key] = append(blogItems[key], item)

	b, err := json.Marshal(blogItems)
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(filestr, b, 0644)

	addlink.AddLinktoAllfiles(linksdir,stopic, item.Stitle)

}
