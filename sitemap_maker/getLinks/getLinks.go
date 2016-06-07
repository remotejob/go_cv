package getLinks

import (

	"fmt"
	"log/syslog"
	"os"
	//	"strconv"
	"encoding/csv"
	"path/filepath"
	//	"strings"
)

var links []string
var maplinks map[string][]string

// Walker file dir
//func walkpath(pathstr string, f os.FileInfo, err error) error {
//
//	if f.IsDir() {
//
//		dirscplit := strings.Split(pathstr, "/")
//
//		link := "http:/"
//
//		for i := 6; i < len(dirscplit); i++ {
//
//			link = link + "/" + dirscplit[i]
//
//			if i == len(dirscplit)-1 {
//				if i == 6 {
//					link = link + "/index.html"
//					links =append(links,link)
//				} else {
//					link = link + ".html"
//					links =append(links,link)
//				}
//
//			}
//
//		}
//
//	}
//	return nil
//}

func walkpath(pathstr string, f os.FileInfo, err error) error {

	if !f.IsDir() {

        site := f.Name()
		csvfile, err := os.Open(pathstr)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer csvfile.Close()
		reader := csv.NewReader(csvfile)

		records, err := reader.ReadAll()
		if err != nil {

			fmt.Println(err)
			return err
		} else {
			
			for _,record := range records {
				maplinks[site]=append(maplinks[site],record[0])
			}
			
		}

	}

	return nil
}

func GetAllLinks(golog syslog.Writer, linksdir string)  map[string][]string{
	
	maplinks=make(map[string][]string)
	
	filepath.Walk(linksdir, walkpath)

	return maplinks

}
