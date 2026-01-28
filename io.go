//
// This file is part of go-algorithms.
//
// Copyright (c) 2024-2026 go-algorithms (go.bug.st/f) authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package f

import (
	"bytes"
	"io"
)

const defaultBufferSize = 1024

type callbackWriter struct {
	callback func(line string)
	buffer   []byte
}

// NewCallbackWriter creates a WriterCloser that will buffer input and call the provided callback for each complete line.
// The last line (if not ending with a newline) will be processed when Close() is called.
func NewCallbackWriter(process func(line string)) io.WriteCloser {
	return &callbackWriter{
		callback: process,
		buffer:   make([]byte, 0, defaultBufferSize),
	}
}

// Write implements the io.Writer interface.
func (p *callbackWriter) Write(data []byte) (int, error) {
	l := len(data)
	for {
		if len(data) == 0 {
			return l, nil
		}

		idx := bytes.IndexByte(data, '\n')
		if idx == -1 {
			// No complete line found, buffer the data
			p.buffer = append(p.buffer, data...)
			return l, nil
		}

		if len(p.buffer) == 0 {
			// Fast path: no buffered data, process directly from input
			p.callback(string(data[:idx]))
			data = data[idx+1:]
			continue
		}

		// Append up to the newline to the buffer and process
		p.buffer = append(p.buffer, data[:idx]...)
		p.callback(string(p.buffer))

		// Clear the buffer and continue with remaining data
		p.buffer = p.buffer[:0]
		data = data[idx+1:]
	}
}

// Close implements the io.Closer interface.
func (p *callbackWriter) Close() error {
	if len(p.buffer) > 0 {
		p.callback(string(p.buffer))
		p.buffer = p.buffer[:0]
	}
	return nil
}
