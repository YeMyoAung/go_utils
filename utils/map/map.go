package map_utils

// Keys returns a slice of all the keys in the given map.
//
// source map[K]V
// []K
func Keys[K comparable, V any](source map[K]V) []K {
	values := make([]K, 0, len(source))

	for key := range source {
		values = append(values, key)
	}

	return values
}

// Values returns a slice of all the values in the given map.
//
// source map[K]V
// []V
func Values[K comparable, V any](source map[K]V) []V {
	values := make([]V, 0, len(source))

	for _, value := range source {
		values = append(values, value)
	}

	return values
}
