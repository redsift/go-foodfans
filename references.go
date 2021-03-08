package foodfans

import (
	"encoding/binary"
	"math/rand"
	"strings"
	"sync"
	"time"

	. "github.com/redsift/go-foodfans/lookup"
)

const sep = "_"

var _default = NewFoodFans(rand.NewSource(Seed()))

func New() string {
	return _default.New()
}

type FoodFans struct {
	rnd *sync.Pool
}

func NewFoodFans(src rand.Source) *FoodFans {
	return &FoodFans{
		rnd: &sync.Pool{New: func() interface{} { return rand.New(src) }},
	}
}

func (f *FoodFans) New() string {
	rng := f.rnd.Get().(*rand.Rand)

	var w strings.Builder

	w.WriteString(Verb(rng.Intn(int(LastVerb))).String())
	w.WriteString(sep)
	w.WriteString(Adjective(rng.Intn(int(LastAdjective))).String())
	w.WriteString(sep)
	w.WriteString(Noun(rng.Intn(int(LastNoun))).String())

	f.rnd.Put(rng)

	return w.String()
}

func Seed() int64 {
	var seed [8]byte
	if _, err := rand.Read(seed[:]); err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	return time.Now().UnixNano() * int64(binary.LittleEndian.Uint64(seed[:]))
}
