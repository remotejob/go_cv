package main

import (
    "testing"
    "net/http"
	"net/http/httptest"
	  "log"
	  "fmt"    
)

func TestServer(t *testing.T) {
	_, err := http.NewRequest("GET", "http://www.test.com/sitemapp.xml", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	
	fmt.Println(resp.Code)
	log.Println(resp.Code)
	main()
	
	
}

