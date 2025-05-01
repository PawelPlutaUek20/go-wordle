package algorithms

import (
	"go-wordle/correctness"
)

type Guesser interface {
	Guess(history []Guess) string
}

type Guess struct {
	Word string
	Mask [5]correctness.Correctness
}

func (g Guess) Matches(word string) bool {
	mask := correctness.Check(word, g.Word)
	return mask == g.Mask
}
