package algorithms

import (
	"go-wordle/correctness"
	"math"
)

type NaiveGuesser struct {
	remaining map[string]int
}

func NewNaiveGuesser(dictionary map[string]int) NaiveGuesser {
	return NaiveGuesser{
		remaining: dictionary,
	}
}

func (g NaiveGuesser) Guess(history []Guess) string {
	if len(history) == 0 {
		return "tares"
	}

	guess := history[len(history)-1]
	for candidate := range g.remaining {
		if !guess.Matches(candidate) {
			delete(g.remaining, candidate)
		}
	}

	remainingCount := 0
	for _, count := range g.remaining {
		remainingCount += count
	}

	var best string
	var bestGoodness = -1.0

	for word := range g.remaining {
		goodness := 0.0
		for _, pattern := range correctness.Patterns() {
			inPatternTotal := 0
			for candidate, count := range g.remaining {
				guess := Guess{
					Word: word,
					Mask: pattern,
				}
				if guess.Matches(candidate) {
					inPatternTotal += count
				}
			}
			if inPatternTotal == 0 {
				continue
			}
			probability := float64(inPatternTotal) / float64(remainingCount)
			goodness += -(probability * math.Log2(probability))
		}

		if goodness > bestGoodness {
			best = word
			bestGoodness = goodness
		}
	}

	return best
}
