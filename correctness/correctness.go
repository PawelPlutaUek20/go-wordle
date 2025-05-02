package correctness

type Correctness uint8

const (
	Wrong Correctness = iota
	Misplaced
	Correct
)

func Check(answer string, guess string) [5]Correctness {
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

			if answerRunes[j] == guessRunes[i] {
				correctness[i] = Misplaced
				used[j] = true
				break
			}
		}
	}

	return correctness
}

var cachedPatterns [][5]Correctness = nil

func Patterns() [][5]Correctness {
	if cachedPatterns != nil {
		return cachedPatterns
	}

	values := [3]Correctness{Correct, Misplaced, Wrong}
	patterns := make([][5]Correctness, 0, 243)

	for i0 := 0; i0 < 3; i0++ {
		for i1 := 0; i1 < 3; i1++ {
			for i2 := 0; i2 < 3; i2++ {
				for i3 := 0; i3 < 3; i3++ {
					for i4 := 0; i4 < 3; i4++ {
						pattern := [5]Correctness{
							values[i0],
							values[i1],
							values[i2],
							values[i3],
							values[i4],
						}
						patterns = append(patterns, pattern)
					}
				}
			}
		}
	}

	cachedPatterns = patterns
	return cachedPatterns
}
