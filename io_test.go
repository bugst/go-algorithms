//
// This file is part of go-algorithms.
//
// Copyright (c) 2024-2026 go-algorithms (go.bug.st/f) authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCallbackWriter(t *testing.T) {
	t.Run("ProcessLines", func(t *testing.T) {
		var lines []string
		w := NewCallbackWriter(func(line string) { lines = append(lines, line) })
		require.NotNil(t, w, "NewCallbackWriter returned nil")

		// Write with two complete lines and one partial
		chunk1 := []byte("first\nsecond\nthird")
		n, err := w.Write(chunk1)
		require.NoError(t, err)
		require.Equal(t, len(chunk1), n)
		require.Equal(t, []string{"first", "second"}, lines)

		n, err = w.Write([]byte(""))
		require.NoError(t, err)
		require.Equal(t, 0, n)

		// Complete the partial and add another full line
		chunk2 := []byte("\nfourth\n\n")
		n, err = w.Write(chunk2)
		require.NoError(t, err)
		require.Equal(t, len(chunk2), n)

		// Write a partial line and then close
		n, err = w.Write([]byte("fifth"))
		require.NoError(t, err)
		require.Equal(t, 5, n)
		require.NoError(t, w.Close())

		wants := []string{"first", "second", "third", "fourth", "", "fifth"}
		require.Equal(t, wants, lines)
	})

	t.Run("BufferGrowth", func(t *testing.T) {
		w := NewCallbackWriter(func(line string) {})
		initialCap := cap(w.(*callbackWriter).buffer)
		// Ensure initial capacity is already greater than the bigger chunk we are testing
		require.Greater(t, initialCap, 100)

		// Run a series of writes to test buffer growth
		for i := 0; i < initialCap*10; i++ {
			n, err := w.Write([]byte("12345"))
			require.NoError(t, err)
			require.Equal(t, 5, n)

			n, err = w.Write([]byte("\n12345"))
			require.NoError(t, err)
			require.Equal(t, 6, n)

			n, err = w.Write([]byte("6347856387657823\n12345"))
			require.NoError(t, err)
			require.Equal(t, 22, n)
		}

		finalCap := cap(w.(*callbackWriter).buffer)
		fmt.Println(finalCap)
		require.Equal(t, initialCap, finalCap)
	})
}
