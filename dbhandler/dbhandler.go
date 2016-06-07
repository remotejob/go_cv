package dbhandler

import (
	"github.com/remotejob/go_cv/domains"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//	"time"
)

func GetAllStitle(session mgo.Session, locale string, themes string) map[string]struct{} {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cv").C("cv")

	var results []domains.Md

	err := c.Find(bson.M{"locale": bson.M{"$exists": true}, "themes": bson.M{"$exists": true}}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	uniqstitle := make(map[string]struct{})

	for _, result := range results {

		fmt.Println(result.Site)

		for _, item := range result.Items {

			uniqstitle[item.Stitle] = struct{}{}

		}

	}

	return uniqstitle

}

func InsertRecord(session mgo.Session, locale string, themes string, site string, menu string, stopic string,topic string, item domains.BlogItem) {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cv").C("cv")

	result := domains.Md{}

	//	fmt.Println(locale, themes, site, menu, stopic)

	count, err := c.Find(bson.M{"locale": locale, "themes": themes, "site": site, "menu": menu, "stopic": stopic}).Limit(1).Count()

	//    fmt.Println(count)

	if err != nil {

		log.Fatal(err)
	}
	if count == 0 {
//		fmt.Println("not exists")

		toinsert := []domains.BlogItem{item}

		toinsertmd := domains.Md{locale, themes, site, menu, stopic,topic, toinsert}

		err := c.Insert(toinsertmd)
		if err != nil {
			log.Fatal(err)
		}

	} else if count == 1 {
//		fmt.Println("exist", count)

		err := c.Find(bson.M{"locale": locale, "themes": themes, "site": site, "menu": menu, "stopic": stopic}).One(&result)

		if err != nil {
			log.Fatal(err)
		}

		result.Items = append(result.Items, item)

		err = c.Update(bson.M{"locale": locale, "themes": themes, "site": site, "menu": menu, "stopic": stopic}, result)
		if err != nil {
			log.Fatal(err)
		}

	} else {

		fmt.Println("Records must be 1 !!!", count)

	}

}
