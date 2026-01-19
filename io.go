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
	p.buffer = append(p.buffer, data...)
	for {
		idx := bytes.IndexByte(p.buffer, '\n')
		if idx == -1 {
			break
		}
		line := p.buffer[:idx] // Do not include \n
		p.buffer = p.buffer[idx+1:]
		p.callback(string(line))
	}
	return len(data), nil
}

func (p *callbackWriter) Close() error {
	if len(p.buffer) > 0 {
		p.callback(string(p.buffer))
		p.buffer = p.buffer[:0]
	}
	return nil
}
