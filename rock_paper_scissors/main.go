package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rockPaperScissors()
}

func rockPaperScissors() {
	scanner := bufio.NewScanner(os.Stdin)
	rockPaperScissors := []string{
		"rock",
		"paper",
		"scissors",
	}
	fmt.Println("Let's play Rock Paper Scissors!")

	currentScore := 0
	roundsPlayed := 0
newRound:
	for {
		fmt.Println("Make your move...")
		for {
			fmt.Println(" ")
			fmt.Print("Choice > ")
			if !scanner.Scan() {
				fmt.Println("Could not read input")
			}
			userChoice := strings.ToLower(strings.TrimSpace(scanner.Text()))
			randChoice := rockPaperScissors[rand.Intn(len(rockPaperScissors))]

			firstLine := fmt.Sprintf("You chose: %s\n", userChoice)
			secondLine := fmt.Sprintf("Computer chose: %s\n", randChoice)
			pause := time.Second / 2

			fmt.Println(" ")
			fmt.Println("Rock!")
			time.Sleep(pause)
			fmt.Println("Paper!")
			time.Sleep(pause)
			fmt.Println("Scissors!")
			time.Sleep(pause)
			fmt.Println("SHOOT!")
			fmt.Println(" ")

			if userChoice == randChoice {
				fmt.Printf("You chose: %s\n", userChoice)
				fmt.Printf("Computer chose: %s\n", randChoice)
				fmt.Println("It's a draw!")
				break
			}

			switch userChoice {
			case "rock":
				switch randChoice {
				case "paper":
					fmt.Print(firstLine)
					fmt.Print(secondLine)
					fmt.Println("You lose!")
					fmt.Println(" ")
				case "scissors":
					fmt.Print(firstLine)
					fmt.Print(secondLine)
					fmt.Println("You won!")
					fmt.Println(" ")
				}
			case "paper":
				switch randChoice {
				case "scissors":
					fmt.Print(firstLine)
					fmt.Print(secondLine)
					fmt.Println("You lose!")
					fmt.Println(" ")
				case "rock":
					fmt.Print(firstLine)
					fmt.Print(secondLine)
					fmt.Println("You won!")
					fmt.Println(" ")
				}
			case "scissors":
				switch randChoice {
				case "rock":
					fmt.Print(firstLine)
					fmt.Print(secondLine)
					fmt.Println("You lose!")
					fmt.Println(" ")
				case "paper":
					fmt.Print(firstLine)
					fmt.Print(secondLine)
					fmt.Println("You won!")
					fmt.Println(" ")
				}
			}
			break
		}

		fmt.Println("Do you want to go again?")
		fmt.Print("Y/n > ")
		if !scanner.Scan() {
			fmt.Println("Could not read input")
		}
		if scanner.Text() == "n" || scanner.Text() == "no" {
			currentScore++
			roundsPlayed++
			fmt.Printf("You won %d rounds out of %d\n", currentScore, roundsPlayed)
			fmt.Println("Thanks for playing!")
			return
		}
		if scanner.Text() == "y" || scanner.Text() == "yes" {
			currentScore++
			roundsPlayed++
			fmt.Println("Buckle up buttercup!")
			fmt.Println("Starting new round...")
			fmt.Println(" ")
			continue newRound
		}
	}
}
