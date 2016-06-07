package phrases

import (
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"github.com/remotejob/go_cv/mark/dbgetall"
	"github.com/remotejob/go_cv/mark/phrases/insertnewphrases"	
)

var mapold map[string]struct{}
var newinsert []string

func Elaborate(locale string, themes string, db sql.DB, records [][]string) {

	mapold = make(map[string]struct{})
	old := dbgetall.GetAll(db, locale, themes, "phrases", "phrase")

	for _, val := range old {

		mapold[val] = struct{}{}

	}
	for _, record := range records {

		if _, ok := mapold[record[0]]; ok {

		} else {

			newinsert = append(newinsert, record[0])

		}
	}
	
	insertnewphrases.InsertAll(db,locale,themes,newinsert) 

}
