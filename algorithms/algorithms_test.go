package algorithms

import (
	"go-wordle/correctness"

	"github.com/stretchr/testify/require"
	"testing"
)

const W = correctness.Wrong
const C = correctness.Correct
const M = correctness.Misplaced

func TestMatch(t *testing.T) {
	tests := []struct {
		prev   string
		mask   [5]correctness.Correctness
		next   string
		allows bool
	}{
		{"abcde", [5]correctness.Correctness{C, C, C, C, C}, "abcde", true},
		{"abcdf", [5]correctness.Correctness{C, C, C, C, C}, "abcde", false},
		{"abcde", [5]correctness.Correctness{W, W, W, W, W}, "fghij", true},
		{"abcde", [5]correctness.Correctness{M, M, M, M, M}, "eabcd", true},
		{"baaaa", [5]correctness.Correctness{W, C, M, W, W}, "caacc", false},
		{"baaaa", [5]correctness.Correctness{W, C, M, W, W}, "aaccc", true},
		{"aaabb", [5]correctness.Correctness{C, M, W, W, W}, "accaa", false},
		{"abcde", [5]correctness.Correctness{W, W, W, W, W}, "bcdea", false},

		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "wacko", true},
		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "wahoo", true},
		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "waldo", true},
		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "wangs", true},
		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "watch", true},
		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "wazoo", true},
		{"weary", [5]correctness.Correctness{C, W, M, W, W}, "wokka", true},
	}

	for _, tt := range tests {
		guess := Guess{
			Word: tt.prev,
			Mask: tt.mask,
		}

		require.Equal(t, guess.Matches(tt.next), tt.allows)
	}
}
