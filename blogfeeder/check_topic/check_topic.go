package check_topic

import (

	"github.com/remotejob/go_cv/domains"

)

func CheckTopicExist(topic string, items []domains.BlogItem) bool {
	

	if len(items) == 0 {
		
		return true		
		
	}
	
	return false
	
}