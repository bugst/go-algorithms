//
// This file is part of go-algorithms.
//
// Copyright 2024 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	f "go.bug.st/f"
)

func TestFutures(t *testing.T) {
	{
		var futureInt f.Future[int]
		go func() {
			time.Sleep(100 * time.Millisecond)
			futureInt.Send(5)
		}()
		require.Equal(t, 5, futureInt.Await())
		go func() {
			require.Equal(t, 5, futureInt.Await())
		}()
		go func() {
			require.Equal(t, 5, futureInt.Await())
		}()
	}
	{
		var futureInt f.Future[int]
		futureInt.Send(5)
		require.Equal(t, 5, futureInt.Await())
		require.Equal(t, 5, futureInt.Await())
		require.Equal(t, 5, futureInt.Await())
	}
}
