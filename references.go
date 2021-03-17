package foodfans

import (
	"encoding/binary"
	"math/rand"
	"strings"
	"time"

	. "github.com/redsift/go-foodfans/lookup"
)

const sep = "_"

func init() {
	rand.Seed(Seed())
}

func New() string {
	var w strings.Builder

	w.WriteString(Verb(rand.Intn(int(LastVerb))).String())
	w.WriteString(sep)
	w.WriteString(Adjective(rand.Intn(int(LastAdjective))).String())
	w.WriteString(sep)
	w.WriteString(Noun(rand.Intn(int(LastNoun))).String())

	return w.String()
}

func Seed() int64 {
	var seed [8]byte
	if _, err := rand.Read(seed[:]); err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	return time.Now().UnixNano() * int64(binary.LittleEndian.Uint64(seed[:]))
}
