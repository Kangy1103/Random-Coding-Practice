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
	max := 0

	fmt.Println("Select a difficulty, easy, medium or hard")
	fmt.Print("Select difficulty > ")
	if !lineRead.Scan() {
		fmt.Println("Could not read input")
	}
	difficulty := lineRead.Text()
	switch difficulty {
	case "easy":
		max = 10
	case "medium":
		max = 100
	case "hard":
		max = 1000
	default:
		max = 100
	}

	number := rand.IntN(max)
	numGuesses := 0
	fmt.Printf("I'm thinking of a number between 0 and %d\n", max)
	fmt.Println("Take a wild guess")
	for {
		fmt.Print("Guess > ")
		if !lineRead.Scan() {
			fmt.Println("Could not read input")
		}
		input := strings.TrimSpace(lineRead.Text())
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("This is a number game my dude")
			continue
		}
		if guess < 0 || guess > max {
			fmt.Printf("Please guess a number between 0 and %d\n", max)
			continue
		}
		if guess == number {
			numGuesses++
			fmt.Println("Well done, you got it!")
			fmt.Printf("Number of guesses: %d\n", numGuesses)
			return
		} else if guess < number {
			fmt.Println("Too low, try again")
			numGuesses++
		} else {
			fmt.Println("Too high, try again")
			numGuesses++
		}
	}
}
