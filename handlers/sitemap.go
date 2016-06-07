package handlers

import (
	"bytes"
//	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	//	"time"
)

func CheckServeSitemap(w http.ResponseWriter, r *http.Request) {

	sitefull := r.Host
	site := strings.Split(sitefull, ":")[0]
	//	site :=r.URL.String()
	filestr := "maps/sitemap_" + site + ".xml"
//	fmt.Println("site", filestr)
	if _, err := os.Stat(filestr); os.IsNotExist(err) {

		http.NotFound(w, r)

	} else {
		f, _ := os.Open(filestr)
		buf := bytes.NewBuffer(nil)
		io.Copy(buf,f)

		w.Header().Add("Content-type", "application/xml")
		w.Write(buf.Bytes())

	}

}
