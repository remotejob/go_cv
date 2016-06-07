package handlers

import (
    "testing"
     "log" 
    "net/http"
	"net/http/httptest" 
	 "github.com/zenazn/goji/web"    
)

func TestElaborate(t *testing.T) {
	
	req, err := http.NewRequest("GET", "http://www.test.com/contact.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	var c  web.C
	
	Elaborate(c,resp,req)
	
	log.Println(resp.Code)	
	
}

