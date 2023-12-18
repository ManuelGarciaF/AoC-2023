// Why the fuck is this not part of the language
package utils

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

func (s Set[T]) Remove(t T) {
	delete(s, t)
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}
