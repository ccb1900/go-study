package server

import "sync"

type K string
type V string
type Store struct {
	space map[K]V
	lock  sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		space: make(map[K]V),
	}
}

func (s *Store) All() map[K]V {
	return s.space
}

func (s *Store) SetRaw(k K, v V) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.space[k] = v
}
func (s *Store) Set(k string, v string) {
	s.SetRaw(K(k), V(v))
}

func (s *Store) GetRaw(k K) V {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if v, exit := s.space[k]; exit {
		return v
	}
	return "(empty)"
}

func (s *Store) Get(k string) string {
	return string(s.GetRaw(K(k)))
}
