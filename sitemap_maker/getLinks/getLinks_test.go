package getLinks

import (
//	"fmt"
//	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
//	"startones"
	"testing"
)

func TestGetAllLinks(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	rootpath := "/home/juno/git/cv/version_desk_react_00/links"

//	res := GetAllLinks(golog, c, site)
	maplinks :=GetAllLinks(*golog, rootpath)
	
	log.Println(maplinks) 

//	for _, character := range res {
//
//		fmt.Println(character.Moto)
//
//	}

}
