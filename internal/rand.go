package internal

import (
	"math/rand"
	"sync"
)

type SyncSource struct {
	lock *sync.Mutex
	src  rand.Source
}

func NewSyncSource(src rand.Source) *SyncSource {
	return &SyncSource{
		lock: new(sync.Mutex),
		src:  src,
	}
}

func (s *SyncSource) Int63() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.src.Int63()
}

func (s *SyncSource) Seed(seed int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.src.Seed(seed)
}
