package utils

// MapK return all keys as slice
func MapK[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, len(m))

	for k := range m {
		r = append(r, k)
	}
	return r
}

// MapV return all values as slice
func MapV[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, len(m))

	for _, v := range m {
		r = append(r, v)
	}

	return r
}
