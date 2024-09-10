//
// This file is part of go-algorithms.
//
// Copyright 2024 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	f "go.bug.st/f"
)

func TestFilter(t *testing.T) {
	a := []string{"aaa", "bbb", "ccc"}
	require.Equal(t, []string{"bbb", "ccc"}, f.Filter(a, func(x string) bool { return x > "b" }))
	b := []int{5, 9, 15, 2, 4, -2}
	require.Equal(t, []int{5, 9, 15}, f.Filter(b, func(x int) bool { return x > 4 }))
}

func TestEqualsAndNotEquals(t *testing.T) {
	require.True(t, f.Equals(int(0))(0))
	require.False(t, f.Equals(int(1))(0))
	require.True(t, f.Equals("")(""))
	require.False(t, f.Equals("abc")(""))

	require.False(t, f.NotEquals(int(0))(0))
	require.True(t, f.NotEquals(int(1))(0))
	require.False(t, f.NotEquals("")(""))
	require.True(t, f.NotEquals("abc")(""))
}

func TestMap(t *testing.T) {
	value := []string{"hello", " world ", " how are", "you? "}
	{
		parts := f.Map(value, strings.TrimSpace)
		require.Equal(t, 4, len(parts))
		require.Equal(t, "hello", parts[0])
		require.Equal(t, "world", parts[1])
		require.Equal(t, "how are", parts[2])
		require.Equal(t, "you?", parts[3])
	}
	{
		parts := f.ParallelMap(value, strings.TrimSpace)
		require.Equal(t, 4, len(parts))
		require.Equal(t, "hello", parts[0])
		require.Equal(t, "world", parts[1])
		require.Equal(t, "how are", parts[2])
		require.Equal(t, "you?", parts[3])
	}
}
