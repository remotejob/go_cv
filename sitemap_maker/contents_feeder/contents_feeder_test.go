package contents_feeder

import (
    "testing"
)

func TestMakeContents(t *testing.T) {
	
	rootdir :="/home/juno/git/go_cv/www"
	pathstr :="www.test.com/contact.html"	
	
	MakeContents(rootdir,pathstr)	

}

