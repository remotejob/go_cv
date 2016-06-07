package addlink

import (
	"bufio"
	"github.com/remotejob/go_cv/comutils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//	"bytes"
	//	"bufio"
)

func AddLinktoAllfiles(dir string, stopic string, stitle string) string {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	mp := make(map[int]string)

	for i, file := range files {

		mp[i] = file.Name()

	}

	random_file_num := comutils.Random(0, len(mp))

	site :=mp[random_file_num]
		
	linkfile := filepath.Join(dir, site)

	//	fileHandle, _ := os.OpenFile(linkfile, os.O_APPEND|os.O_WRONLY, 0666)
	fileHandle, _ := os.OpenFile(linkfile, os.O_APPEND|os.O_RDWR, 0666)
	
	uniqlinks :=make(map[string]struct{})
	
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
//		log.Println(scanner.Text())
		uniqlinks[scanner.Text()] =struct{}{}
	}

	// check for errors
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	linktoadd := filepath.Join("/blog/", stopic, stitle+".html")

	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	linktopicadd := filepath.Join("/blog/", stopic)
	
	if _, ok := uniqlinks[linktopicadd]; !ok {	
		
		fmt.Fprintln(writer, linktopicadd)

	}

//	if topicOK {
//		linktopicadd := filepath.Join("/blog/", stopic)
//		fmt.Fprintln(writer, linktopicadd)
//	}
	fmt.Fprintln(writer, linktoadd)

	writer.Flush()

return site

}
