package goset

type threadUnsafeSet map[string]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (s *threadUnsafeSet) Add(i string) bool {
	// Return false if found
	if _, found := (*s)[i]; found {
		return false
	}
	(*s)[i] = struct{}{}
	return true
}

func (s *threadUnsafeSet) Remove(i string) {
	delete(*s, i)
}

func (s *threadUnsafeSet) Contains(i ...string) bool {
	// Return false if not found
	for _, val := range i {
		if _, found := (*s)[val]; !found {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet) Clear() {
	*s = threadUnsafeSet{}
}
