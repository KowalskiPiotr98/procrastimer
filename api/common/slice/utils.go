package slice

// Map transforms a slice of elements into a new slice by applying the mapper func to all elements.
func Map[T interface{}, E interface{}](slice []T, mapper func(T) E) []E {
	results := make([]E, len(slice))
	for _, t := range slice {
		results = append(results, mapper(t))
	}
	return results
}
