package mgenerator

import (
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"mark/dbgetall"
	"testing"
	"log"
	"fmt"
)

func TestGenerate(t *testing.T) {

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	oldphrases := dbgetall.GetAll(*db, "en_US", "programming", "phrases", "phrase")
	oldkeywords := dbgetall.GetAll(*db, "en_US", "programming", "keywords", "keyword")	

	result :=Generate(oldkeywords,oldphrases)
	
	fmt.Println(result)

}
