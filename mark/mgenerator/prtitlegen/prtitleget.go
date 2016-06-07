package prtitlegen

import (
	"github.com/remotejob/go_cv/comutils"

)

func Generate(keywordsarr []string) []string {

	var titlearr []string
	for i := 0; i < 2; i++ {
	
		titlearr =append(titlearr,keywordsarr[comutils.Random(0, len(keywordsarr))])
	
	}
	
	
	return titlearr

}
