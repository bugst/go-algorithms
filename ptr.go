package f

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// UnwrapOrDefault returns the ptr value if it is not nil, otherwise returns the zero value.
func UnwrapOrDefault[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}

	var zero T
	return zero
}
