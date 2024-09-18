package f_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	f "go.bug.st/f"
)

func TestPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 12, *f.Ptr(12))
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "hello", *f.Ptr("hello"))
	})
}

func TestUnwrapOrDefault(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		given := 12
		assert.Equal(t, 12, f.UnwrapOrDefault(&given))
	})

	t.Run("nil", func(t *testing.T) {
		assert.Equal(t, "", f.UnwrapOrDefault[string](nil))
	})
}
