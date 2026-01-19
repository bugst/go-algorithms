package f_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	f "go.bug.st/f"
)

func TestNewCallbackWriter_Write_ProcessesLines(t *testing.T) {
	var lines []string
	w := f.NewCallbackWriter(func(line string) { lines = append(lines, line) })
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
}
