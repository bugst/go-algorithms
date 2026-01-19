package f

import "iter"

// FilterIter takes an iterator and a matcher and returns only those elements that satisfy the matcher.
func FilterIter[T any](values iter.Seq[T], matcher Matcher[T]) iter.Seq[T] {
	return func(yield func(x T) bool) {
		for x := range values {
			if matcher(x) {
				if !yield(x) {
					return
				}
			}
		}
	}
}

// Map applies the Mapper function to each element of the iterator.
func MapIter[T, U any](values iter.Seq[T], mapper Mapper[T, U]) iter.Seq[U] {
	return func(yield func(x U) bool) {
		for x := range values {
			if !yield(mapper(x)) {
				return
			}
		}
	}
}

// Reducer is a function that reduces an iterator's elements to a single value.
func ReduceIter[T any](values iter.Seq[T], reducer Reducer[T], initialValue ...T) T {
	var result T
	if len(initialValue) > 1 {
		panic("initialValue must be a single value")
	} else if len(initialValue) == 1 {
		result = initialValue[0]
	}
	for v := range values {
		result = reducer(result, v)
	}
	return result
}
