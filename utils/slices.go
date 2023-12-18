package utils

func Contains[S ~[]E, E comparable](s S, e E) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func RemoveAt[S ~[]E, E any](s S, i int) S {
	return append(s[:i], s[i+1:]...)
}

func Remove[S ~[]E, E comparable](s S, e E) S {
	for i, v := range s {
		if v == e {
			return RemoveAt(s, i)
		}
	}
	return s
}
