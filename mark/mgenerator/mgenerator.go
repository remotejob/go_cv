package mgenerator

import (
	"bytes"
	"github.com/remotejob/go_cv/comutils"
	"github.com/remotejob/go_cv/domains"
	"fmt"
	"github.com/remotejob/go_cv/mark/mgenerator/prtitlegen"
	"math/rand"
	"strings"
	"time"
)

type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Write(b []byte) {
	br := bytes.NewReader(b)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) Generate(n int, keyword string) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		if i == 0 {
			words = append(words, comutils.UpcaseInitial(keyword))
		} else {

			words = append(words, next)
		}
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func Generate(keywords []string, phrases []string) domains.Contents {

	rand.Seed(time.Now().UnixNano())
	prtitle := prtitlegen.Generate(keywords)
	title := comutils.UpcaseInitial(strings.Join(prtitle, " "))
	moto := comutils.UpcaseInitial(phrases[rand.Intn(len(phrases))]) + "."
//	fmt.Println(prtitle)
//	fmt.Println(moto)

	var buffer bytes.Buffer
	prefixLen := 1

	// Seed the random number generator.
	numWords := comutils.Random(300, 1000)

	c := NewChain(prefixLen)

	var sentence string
	for _, phrase := range phrases {

		sentence = sentence + phrase + " "

		if len(sentence) > 80 {

			sentence = sentence + phrase + ". "
			buffer.WriteString(comutils.UpcaseInitial(sentence))
			sentence = ""

		} else {

			sentence = sentence + phrase + " "
		}

	}
	c.Write(buffer.Bytes())
	text := c.Generate(numWords, prtitle[0])

	contents := domains.Contents{title, moto, text, time.Now(), time.Now()}
//	fmt.Println(text)
	return contents
}
