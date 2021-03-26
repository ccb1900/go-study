package server

import (
	"sync"
)

type M int
type N *Client
type ClientStore struct {
	space map[M]N
	lock  sync.RWMutex
}

func NewClientStore() *ClientStore {
	return &ClientStore{
		space: make(map[M]N),
	}
}

func (s *ClientStore) All() map[M]N {
	return s.space
}

func (s *ClientStore) SetRaw(k M, v N) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.space[k] = v
}

func (s *ClientStore) GetRaw(k M) N {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if v, exit := s.space[k]; exit {
		return v
	}
	return nil
}
func (s *ClientStore) Set(k int, v *Client) {
	s.SetRaw(M(k), v)
}

func (s *ClientStore) Get(k int) *Client {
	return s.Get(k)
}

func (s *ClientStore) Remove(k int) {
	s.RemoveRaw(M(k))
}

func (s *ClientStore) RemoveRaw(k M) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.space, k)
}
