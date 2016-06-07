package dbgetall

import (
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"log"
	"fmt"
)

func GetAll(db sql.DB,locale string, themes string, table string,colomns string) []string{


	sqlstr := "select "+colomns+" from "+table+" where locale='" + locale + "' and themes='" + themes + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var outarray []string
	for rows.Next() {
		var keyword string
		rows.Scan(&keyword)
//		fmt.Println(keyword)
		outarray =append(outarray,keyword)
		 
	}
	rows.Close()
	
	fmt.Println("getAll",len(outarray))
	
	return outarray
}
