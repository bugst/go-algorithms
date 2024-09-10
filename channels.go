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
type Future[T any] interface {
	Send(T)
	Await() T
}

type future[T any] struct {
	wg    sync.WaitGroup
	value T
}

// NewFuture creates a new Future[T]
func NewFuture[T any]() Future[T] {
	res := &future[T]{}
	res.wg.Add(1)
	return res
}

// Send a result in the Future. Threads waiting for result will be unlocked.
func (f *future[T]) Send(value T) {
	f.value = value
	f.wg.Done()
}

// Await for a result from the Future, blocks until a result is available.
func (f *future[T]) Await() T {
	f.wg.Wait()
	return f.value
}
