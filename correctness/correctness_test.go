package correctness

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const W = Wrong
const C = Correct
const M = Misplaced

func TestCorrectness(t *testing.T) {
	tests := []struct {
		answer      string
		guess       string
		correctness [5]Correctness
	}{
		{"abcde", "abcde", [5]Correctness{C, C, C, C, C}},
		{"aacde", "abcde", [5]Correctness{C, W, C, C, C}},
		{"aaddd", "baccc", [5]Correctness{W, C, W, W, W}},
		{"aaabb", "azzaz", [5]Correctness{C, W, W, M, W}},

		{"wacko", "weary", [5]Correctness{C, W, M, W, W}},
		{"wahoo", "weary", [5]Correctness{C, W, M, W, W}},
		{"waldo", "weary", [5]Correctness{C, W, M, W, W}},
		{"wangs", "weary", [5]Correctness{C, W, M, W, W}},
		{"watch", "weary", [5]Correctness{C, W, M, W, W}},
		{"wazoo", "weary", [5]Correctness{C, W, M, W, W}},
		{"wokka", "weary", [5]Correctness{C, W, M, W, W}},
	}

	for _, tt := range tests {
		correctness := Check(tt.answer, tt.guess)
		require.Equal(t, tt.correctness, correctness)
	}
}
