//
// This file is part of go-algorithms.
//
// Copyright 2024 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f

// Must should be used to wrap a call to a function returning a value and an error.
// Must returns the value if the errors is nil, or panics otherwise.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return val
}
