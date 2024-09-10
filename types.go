//
// This file is part of go-algorithms.
//
// Copyright 2024 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f

// Matcher is a function that tests if a given value matches a certain criteria.
type Matcher[T any] func(T) bool

// Reducer is a function that combines two values of the same type and return
// the combined value.
type Reducer[T any] func(T, T) T

// Mapper is a function that converts a value of one type to another type.
type Mapper[T, U any] func(T) U
