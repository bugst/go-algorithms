package f

import (
	"bytes"
	"io"
)

const defaultBufferSize = 1024

// callbackWriter is a custom writer that processes each line calling the callback.
type callbackWriter struct {
	callback func(line string)
	buffer   []byte
}

// NewCallbackWriter creates a new CallbackWriter.
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

func (p *callbackWriter) Close() error {
	if len(p.buffer) > 0 {
		p.callback(string(p.buffer))
		p.buffer = p.buffer[:0]
	}
	return nil
}
