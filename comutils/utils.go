package comutils

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"unicode"
	
//    "time"
)

func UpcaseInitial(str string) string {

	runes := []rune(str)

	for i, v := range runes {
	
		return string(unicode.ToUpper(v)) + string(runes[i+1:])

	}

	return ""

}

func Clone(a, b interface{}) {

	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(a)
	dec.Decode(b)
}

func Random(min, max int) int {	
//	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
