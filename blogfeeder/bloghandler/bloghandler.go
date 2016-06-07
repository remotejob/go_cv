package bloghandler

import (
	"fmt"
	"os"
)

func Bhandler(locale string, themes string, file string,title string, contents string) {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println(file, "not EXIST")
	

	}

}
