package f_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	f "go.bug.st/f"
)

func TestMust(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.Equal(t, 12, f.Must(12, nil))
	})

	t.Run("error", func(t *testing.T) {
		want := fmt.Errorf("this is an error")
		assert.PanicsWithValue(t, want.Error(), func() {
			f.Must(0, want)
		})
	})
}

func TestAssert(t *testing.T) {
	t.Run("true", func(_ *testing.T) {
		f.Assert(true, "should not panic")
	})

	t.Run("false", func(t *testing.T) {
		assert.PanicsWithValue(t, "should panic", func() {
			f.Assert(false, "should panic")
		})
	})
}
