package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	randNum()
}

func randNum() {
	lineRead := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Kangy's Magic Number Game!")
newRound:
	for {
		maxGuesses := 0
		max := 0
		fmt.Println(" ")
		fmt.Println("Select a difficulty, easy, medium or hard")
		fmt.Print("Select difficulty > ")
		if !lineRead.Scan() {
			fmt.Println("Could not read input")
		}
		difficulty := lineRead.Text()
		switch difficulty {
		case "easy":
			max = 10
			maxGuesses = 5
		case "medium":
			max = 100
			maxGuesses = 10
		case "hard":
			max = 1000
			maxGuesses = 15
		default:
			max = 100
			maxGuesses = 10
		}
		number := rand.IntN(max)
		remainingGuesses := 0
		fmt.Println(" ")
		fmt.Printf("I'm thinking of a number between 0 and %d\n", max)
		fmt.Printf("You get a total of %d attempts to get it right\n", maxGuesses)
		fmt.Println("Take a wild guess")
		fmt.Println(" ")

		lastGuess := 0
		hasLastGuess := false
	currentRound:
		for {

			fmt.Print("Guess > ")
			if !lineRead.Scan() {
				fmt.Println("Could not read input")
			}
			input := strings.TrimSpace(lineRead.Text())
			guess, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("This is a number game my dude")
				fmt.Println(" ")
				continue currentRound
			}
			if guess < 0 || guess > max {
				fmt.Printf("Please guess a number between 0 and %d\n", max)
				fmt.Println(" ")
				continue currentRound
			}
			if guess == number {
				remainingGuesses++
				fmt.Println("Well done, you got it!")
				fmt.Printf("Number of guesses: %d\n", remainingGuesses)
				fmt.Println(" ")
				break
			} else if guess < number {
				fmt.Println("Too low, try again")
				remainingGuesses++
				maxGuesses--

				if maxGuesses == 0 {
					fmt.Println("No guesses left, womp womp")
					fmt.Println(" ")
					break
				} else {
					fmt.Printf("You have %d guesses remaining\n", maxGuesses)
					fmt.Println(" ")
				}
			} else {
				fmt.Println("Too high, try again")
				remainingGuesses++
				maxGuesses--
				if maxGuesses == 0 {
					fmt.Println("No guesses left, womp womp")
					fmt.Println(" ")
					break
				} else {
					fmt.Printf("You have %d guesses remaining\n", maxGuesses)
					fmt.Println(" ")
				}
			}
			if hasLastGuess {
				lastDistance := lastGuess - number
				if lastDistance < 0 {
					lastDistance = -lastDistance
				}
				currentDistance := guess - number
				if currentDistance < 0 {
					currentDistance = -currentDistance
				}
				if currentDistance < lastDistance {
					fmt.Println("Warmer!")
					fmt.Println(" ")
				} else if currentDistance > lastDistance {
					fmt.Println("Colder!")
					fmt.Println(" ")
				}
			}
			lastGuess = guess
			hasLastGuess = true
		}
		fmt.Println("Try your hand at another round?")
		fmt.Print("Y/n > ")
		if !lineRead.Scan() {
			fmt.Println("Could not read input")
		}
		if lineRead.Text() == "n" || lineRead.Text() == "no" {
			fmt.Println("Thanks for playing!")
			return
		}
		if lineRead.Text() == "y" || lineRead.Text() == "yes" {
			fmt.Println("Buckle up buttercup!")
			fmt.Println("Starting new round...")
			fmt.Println(" ")
			continue newRound
		}
	}
}
