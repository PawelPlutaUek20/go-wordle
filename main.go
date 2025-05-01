package main

import (
	"go-wordle/algorithms"
	"go-wordle/correctness"
	"strconv"
	"strings"

	"bufio"
	"log"
	"os"
)

func main() {
	dict := dictionary()
	guesser := algorithms.NewNaiveGuesser(dict)

	steps := play("weary", guesser)
	log.Println("Game completed in", steps, "guesses")
}

func play(answer string, guesser algorithms.Guesser) int {
	history := make([]algorithms.Guess, 0, 5)
	for i := 1; i <= 5; i++ {
		guess := guesser.Guess(history)
		log.Println("New guess", guess)
		if guess == answer {
			return i
		}
		correctness := correctness.Check(answer, guess)
		newGuess := algorithms.Guess{Word: guess, Mask: correctness}
		history = append(history, newGuess)
	}

	return -1
}

func dictionary() map[string]int {
	dict := make(map[string]int)

	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			log.Fatal("Every line is word + space + count")
		}

		word := fields[0]
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal("Every count is a number")
		}

		dict[word] = count
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dict
}
