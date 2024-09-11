//
// This file is part of go-algorithms.
//
// Copyright 2024 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// Filter takes a slice of type []T and a Matcher[T]. It returns a newly
// allocated slice containing only those elements of the input slice that
// satisfy the matcher.
func Filter[T any](values []T, matcher Matcher[T]) []T {
	var res []T
	for _, x := range values {
		if matcher(x) {
			res = append(res, x)
		}
	}
	return res
}

// Map applies the Mapper function to each element of the slice and returns
// a new slice with the results in the same order.
func Map[T, U any](values []T, mapper Mapper[T, U]) []U {
	res := make([]U, len(values))
	for i, x := range values {
		res[i] = mapper(x)
	}
	return res
}

// ParallelMap applies the Mapper function to each element of the slice and returns
// a new slice with the results in the same order. This is executed among multilple
// goroutines in parallel. If jobs is specified it will indicate the maximum number
// of goroutines to be spawned.
func ParallelMap[T, U any](values []T, mapper Mapper[T, U], jobs ...int) []U {
	res := make([]U, len(values))
	var j int
	if len(jobs) == 0 {
		j = runtime.NumCPU()
	} else if len(jobs) == 1 {
		j = jobs[0]
	} else {
		panic("jobs must be a single value")
	}
	j = min(j, len(values))

	var idx atomic.Int64
	idx.Store(-1)
	var wg sync.WaitGroup
	wg.Add(j)
	for count := 0; count < j; count++ {
		go func() {
			i := int(idx.Add(1))
			for i < len(values) {
				res[i] = mapper(values[i])
				i = int(idx.Add(1))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return res
}

// Reduce applies the Reducer function to all elements of the input values
// and returns the result.
func Reduce[T any](values []T, reducer Reducer[T], initialValue ...T) T {
	var result T
	if len(initialValue) > 1 {
		panic("initialValue must be a single value")
	} else if len(initialValue) == 1 {
		result = initialValue[0]
	}
	for _, v := range values {
		result = reducer(result, v)
	}
	return result
}

// Equals return a Matcher that matches the given value
func Equals[T comparable](value T) Matcher[T] {
	return func(x T) bool {
		return x == value
	}
}

// NotEquals return a Matcher that does not match the given value
func NotEquals[T comparable](value T) Matcher[T] {
	return func(x T) bool {
		return x != value
	}
}

// Uniq return a copy of the input array with all duplicates removed
func Uniq[T comparable](in []T) []T {
	have := map[T]bool{}
	var out []T
	for _, v := range in {
		if have[v] {
			continue
		}
		out = append(out, v)
		have[v] = true
	}
	return out
}
