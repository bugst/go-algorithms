//
// This file is part of go-algorithms.
//
// Copyright 2024 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f

import "sync"

// DiscardCh consumes all incoming messages from the given channel until it's closed.
func DiscardCh[T any](ch <-chan T) {
	for range ch {
	}
}

// Future is an object that holds a result value. The value may be read and
// written asynchronously.
type Future[T any] struct {
	lock  sync.Mutex
	set   bool
	cond  *sync.Cond
	value T
}

// Send a result in the Future. Threads waiting for result will be unlocked.
func (f *Future[T]) Send(value T) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.set = true
	f.value = value
	if f.cond != nil {
		f.cond.Broadcast()
		f.cond = nil
	}
}

// Await for a result from the Future, blocks until a result is available.
func (f *Future[T]) Await() T {
	f.lock.Lock()
	defer f.lock.Unlock()
	for !f.set {
		if f.cond == nil {
			f.cond = sync.NewCond(&f.lock)
		}
		f.cond.Wait()
	}
	return f.value
}
