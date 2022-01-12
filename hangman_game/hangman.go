package main

import (
	"fmt"
	"math/rand"
	"regexp"
)

var lives int = 5

func getWord() string {
	var word string = "hello"
	return word
}

func printWord(word string, correctGuesses map[string]bool) {
	for _, letter := range word {
		actualLetter := string(letter)
		if correctGuesses[actualLetter] {
			fmt.Printf("%v ", actualLetter)
		} else {
			fmt.Printf(" _ ")
		}
	}
	fmt.Println()
}

func giveHint(word string, hints int, correctGuesses map[string]bool) int {
	if hints > 0 {
		for {
			index := rand.Intn(len(word)-0) + 0
			letter := string(word[index])

			if !correctGuesses[letter] {
				hints--
				correctGuesses[letter] = true
				break
			}
		}
	}
	return hints
}
func printGuessed(word string, correctGuesses map[string]bool, wrongGuesses map[string]bool, hints int) {
	printWord(word, correctGuesses)

	fmt.Print("correctGuesses = [")
	for key, _ := range correctGuesses {
		fmt.Printf("%v,", key)
	}
	fmt.Println("]")
	fmt.Print("wrongGuesses = [")
	for key, _ := range wrongGuesses {
		fmt.Printf("%v,", key)
	}
	fmt.Println("]")
	fmt.Println("lives remaining : ", lives)
	fmt.Println("hints remaining : ", hints)

}

func checkIsLost(lives int, hints int, wrongGuesses map[string]bool) bool {
	isLost := false
	if lives == 0 {
		isLost = true
	}

	return isLost
}

func checkLetter(word string, letter string) bool {
	res, _ := regexp.MatchString("[A-Za-z]", letter)
	if res {
		return true
	} else {
		return false
	}
}
func checkIsWon(lives int, hints int, word string, correctGuesses map[string]bool, wrongGuesses map[string]bool) bool {
	var isWon bool = true
	for _, letter := range word {

		actualLetter := string(letter)

		if !correctGuesses[actualLetter] {
			isWon = false
			break
		}
	}
	return isWon
}
func checkWord(word string, userInput string, correctGuesses map[string]bool, wrongGuesses map[string]bool) {
	userInputLength := len(userInput)
	if userInputLength == 1 {
		// check letter
		for _, letter := range word {
			actualLetter := string(letter)
			if actualLetter == userInput {
				correctGuesses[actualLetter] = true

			}

		}
		if correctGuesses[userInput] {

		} else {

			if !wrongGuesses[userInput] {
				wrongGuesses[userInput] = true
				lives--
			}
		}

	} else {
		// check word
		if word == userInput {
			for _, letter := range word {
				correctGuesses[string(letter)] = true
			}
		} else {
			fmt.Println("Wrong guess, try again")
			lives--
		}
	}

}

func main() {
	word := getWord()

	hints := int(len(word) / 3)

	correctGuesses := make(map[string]bool)
	wrongGuesses := make(map[string]bool)

outer:
	for {
		fmt.Print("Enter help for hint \nEnter stop for exit \nEnter character or word :")
		var userInput string
		fmt.Scanf("%s", &userInput)
		switch {
		case userInput == "stop" || userInput == "STOP":
			break outer

		case userInput == "help" || userInput == "HELP":
			hints = giveHint(word, hints, correctGuesses)
			printGuessed(word, correctGuesses, wrongGuesses, hints)

		case checkLetter(word, userInput):
			checkWord(word, userInput, correctGuesses, wrongGuesses)

		}
		isWon := checkIsWon(lives, hints, word, correctGuesses, wrongGuesses)
		printGuessed(word, correctGuesses, wrongGuesses, hints)

		if isWon {

			fmt.Println("Congratulations, you guessed the word correctly")
			break outer
		} else {
			isLost := checkIsLost(lives, hints, wrongGuesses)
			if isLost {

				fmt.Println("You lost...")
				break outer
			}

		}

	}

}
