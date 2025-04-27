package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCorrectness(t *testing.T) {
	tests := []struct {
		answer      string
		guess       string
		correctness [5]Correctness
	}{
		{"speed", "abide", [5]Correctness{Wrong, Wrong, Misplaced, Wrong, Misplaced}},
		{"speed", "erase", [5]Correctness{Misplaced, Wrong, Misplaced, Misplaced, Wrong}},
		{"speed", "steal", [5]Correctness{Correct, Wrong, Correct, Wrong, Wrong}},
		{"speed", "crepe", [5]Correctness{Wrong, Misplaced, Correct, Misplaced, Wrong}},
		{"abcde", "abcde", [5]Correctness{Correct, Correct, Correct, Correct, Correct}},
		{"aacde", "abcde", [5]Correctness{Correct, Wrong, Correct, Correct, Correct}},
		{"aaddd", "baccc", [5]Correctness{Wrong, Correct, Wrong, Wrong, Wrong}},
		{"aaabb", "azzaz", [5]Correctness{Correct, Misplaced, Wrong, Wrong, Wrong}},
	}

	for _, tt := range tests {
		correctness := check(tt.answer, tt.guess)
		require.Equal(t, tt.correctness, correctness)
	}
}
