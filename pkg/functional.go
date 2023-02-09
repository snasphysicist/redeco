package redeco

// mapping provides functional slice mapping functionality
func mapping[T any, U any](a []T, f func(T) U) []U {
	fwe := func(e T) (U, error) {
		return f(e), nil
	}
	m, _ := mapWithError(a, fwe)
	return m
}

// mapWithError provides functional slice mapping functionality
func mapWithError[T any, U any](a []T, f func(T) (U, error)) ([]U, error) {
	m := make([]U, 0)
	for _, e := range a {
		nm, err := f(e)
		if err != nil {
			return make([]U, 0), err
		}
		m = append(m, nm)
	}
	return m, nil
}

// filter returns a new slice containing only the elements of a which satisfy the predicate f
func filter[T any](a []T, f func(T) bool) []T {
	m := make([]T, 0)
	for _, e := range a {
		if f(e) {
			m = append(m, e)
		}
	}
	return m
}

// anyMatch returns true if any element of tests true with f
func anyMatch[T any](a []T, f func(T) bool) bool {
	return len(filter(a, f)) > 0
}

// contains returns true iff e can be found in a
func contains[T comparable](a []T, e T) bool {
	return len(filter(a, func(i T) bool { return i == e })) > 0
}

// unique returns only those elements of a
// such that a[i] != a[j] for all i != j
func unique[T comparable](a []T) []T {
	u := make([]T, 0)
	for _, e := range a {
		if !contains(u, e) {
			u = append(u, e)
		}
	}
	return u
}
