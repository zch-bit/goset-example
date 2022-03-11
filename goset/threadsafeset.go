package goset

import "sync"

type threadSafeSet struct {
	s threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() threadSafeSet {
	return threadSafeSet{
		s: newThreadUnsafeSet(),
	}
}

func (s *threadSafeSet) Add(i string) bool {
	s.Lock()
	ret := s.s.Add(i)
	s.Unlock()
	return ret
}
func (s *threadSafeSet) Remove(i string) {
	s.Lock()
	s.s.Remove(i)
	s.Unlock()

}
func (s *threadSafeSet) Contains(i ...string) bool {
	s.RLock()
	ret := s.s.Contains(i...)
	s.RUnlock()
	return ret
}
func (s *threadSafeSet) Clear() {
	s.Lock()
	s.s = newThreadUnsafeSet()
	s.Unlock()
}
