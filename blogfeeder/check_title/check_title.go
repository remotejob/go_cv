package check_title

import (
	"github.com/remotejob/go_cv/domains"
)

func CheckIfExist(stitle string, items []domains.BlogItem) bool {

	stitlemap := make(map[string]struct{})

	for _, item := range items {

		stitlemap[item.Stitle] = struct{}{}

	}

	if _, ok := stitlemap[stitle]; !ok {
		
		return false

	}
	
	return true

}
