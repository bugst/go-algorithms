package f

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterIter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		matcher  Matcher[int]
		expected []int
	}{
		{
			name:     "even numbers",
			input:    []int{1, 2, 3, 4, 5},
			matcher:  func(x int) bool { return x%2 == 0 },
			expected: []int{2, 4},
		},
		{
			name:     "odd numbers",
			input:    []int{1, 2, 3, 4, 5},
			matcher:  func(x int) bool { return x%2 == 1 },
			expected: []int{1, 3, 5},
		},
		{
			name:     "none",
			input:    []int{2, 4, 6},
			matcher:  func(x int) bool { return x < 0 },
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := slices.Values(tt.input)
			result := slices.Collect(FilterIter(values, tt.matcher))
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMapIter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		mapper   Mapper[int, int]
		expected []int
	}{
		{
			name:     "double",
			input:    []int{1, 2, 3},
			mapper:   func(x int) int { return x * 2 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "negate",
			input:    []int{1, -2, 3},
			mapper:   func(x int) int { return -x },
			expected: []int{-1, 2, -3},
		},
		{
			name:     "identity",
			input:    []int{1, 2, 3},
			mapper:   func(x int) int { return x },
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := slices.Values(tt.input)
			result := slices.Collect(MapIter(values, tt.mapper))
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReduceIter(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		reducer      Reducer[int]
		initialValue int
		expected     int
	}{
		{
			name:         "sum",
			input:        []int{1, 2, 3, 4},
			reducer:      func(a, b int) int { return a + b },
			initialValue: 0,
			expected:     10,
		},
		{
			name:         "product",
			input:        []int{1, 2, 3, 4},
			reducer:      func(a, b int) int { return a * b },
			initialValue: 1,
			expected:     24,
		},
		{
			name:         "subtract",
			input:        []int{10, 2, 3},
			reducer:      func(a, b int) int { return a - b },
			initialValue: 20,
			expected:     5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := slices.Values(tt.input)
			result := ReduceIter(values, tt.reducer, tt.initialValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}
