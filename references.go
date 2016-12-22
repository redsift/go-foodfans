package foodfans

import (
	"math/rand"
	"time"

	. "github.com/redsift/go-foodfans/lookup"
)

const sep = "_"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func New() string {

	v := Verb(rand.Intn(int(LastVerb)))
	a := Adjective(rand.Intn(int(LastAdjective)))
	n := Noun(rand.Intn(int(LastNoun)))
	return v.String() + sep + a.String() + sep + n.String()
}
