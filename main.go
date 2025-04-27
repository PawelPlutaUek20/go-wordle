package main

import (
	"github.com/stretchr/testify/require"
)

type Correctness uint8

const (
	Wrong Correctness = iota
	Misplaced
	Correct
)

type Guess struct {
	word string
	mask [5]Correctness
}

type Guesser interface {
	guess(history []Guess) string
}

func check(answer string, guess string) [5]Correctness {
	require.Len(nil, answer, 5)
	require.Len(nil, guess, 5)

	var correctness [5]Correctness
	var used [5]bool

	guessRunes := []rune(guess)
	answerRunes := []rune(answer)

	for i := 0; i < 5; i++ {
		if answerRunes[i] == guessRunes[i] {
			correctness[i] = Correct
			used[i] = true
		}
	}

	for i := 0; i < 5; i++ {
		if correctness[i] == Correct {
			continue
		}

		for j := 0; j < 5; j++ {
			if used[j] {
				continue
			}

			if answerRunes[i] == guessRunes[j] {
				correctness[i] = Misplaced
				used[j] = true
				break
			}
		}
	}

	return correctness
}

func play(answer string, guesser Guesser) int {
	history := make([]Guess, 0)
	for i := 1; ; i++ {
		guess := guesser.guess(history)
		if guess == answer {
			return i
		}
		correctness := check(answer, guess)
		history = append(history, Guess{guess, correctness})
	}
}

func main() {
}
