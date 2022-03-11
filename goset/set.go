package goset

type Set interface {
	Add(i string) bool
	Remove(i string)
	Contains(i ...string) bool
	Clear()
}

func NewSet() Set {
	//s := newThreadUnsafeSet()
	s := newThreadSafeSet()
	return &s
}
