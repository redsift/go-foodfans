package foodfans

import (
	"math/rand"
	"strings"
	"time"

	. "github.com/redsift/go-foodfans/lookup"
)

const sep = "_"

var _default=NewFoodFans(rand.NewSource(time.Now().UnixNano()))

func New() string {
	return _default.New()
}

type FoodFans struct {
	rand *rand.Rand
}

func NewFoodFans(src rand.Source) *FoodFans {
	return &FoodFans{rand: rand.New(src)}
}

func (f* FoodFans) New() string {
	var w strings.Builder
	w.WriteString(Verb(f.rand.Intn(int(LastVerb))).String())
	w.WriteString(sep)
	w.WriteString(Adjective(f.rand.Intn(int(LastVerb))).String())
	w.WriteString(sep)
	w.WriteString(Noun(f.rand.Intn(int(LastVerb))).String())
	return w.String()
}