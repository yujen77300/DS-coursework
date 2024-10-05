package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	randomNumber string
	isFinished   bool
	guessCount   int
}

func NewGame() *Game {
	return &Game{
		randomNumber: generateNumber(),
		isFinished:   false,
		guessCount:   0,
	}
}

func generateNumber() string {
	rand.Seed(time.Now().UnixNano())
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})
	return strings.Join(digits[:4], "")
}

func (g *Game) checkA(userInput string) int {
	count := 0
	for i := 0; i < 4; i++ {
		if g.randomNumber[i] == userInput[i] {
			count++
		}
	}
	return count
}

func (g *Game) checkB(userInput string) int {
	bCount := 0
	for i := 0; i < 4; i++ {
		if strings.Contains(g.randomNumber, string(userInput[i])) {
			bCount++
		}
	}
	return bCount - g.checkA(userInput)
}

func (g *Game) Play() {
	fmt.Println("Let's play a game of A and B")
	fmt.Println("I will generate a 4-digit number, and you have to guess the numbers one digit at a time.")
	fmt.Println("For every number in the wrong place, you get a B. For every one in the right place, you get an A.")
	fmt.Println("The game ends when you get 4 As!")
	fmt.Println("Type exit at any prompt to exit.")

	for !g.isFinished {
		fmt.Print("Give me your best guess: ")
		var userInput string
		fmt.Scanln(&userInput)

		userInput = strings.ToLower(userInput)
		if userInput == "exit" {
			fmt.Printf("Sorry to see you leave after %d guesses and the number was %s.\n", g.guessCount, g.randomNumber)
			return
		}

		if _, err := strconv.Atoi(userInput); err != nil || len(userInput) != 4 {
			fmt.Println("Please enter exactly 4 digits or 'exit' to end the game.")
			continue
		}

		g.guessCount++
		resultA := g.checkA(userInput)
		resultB := g.checkB(userInput)

		if resultA == 4 {
			g.isFinished = true
			fmt.Println("You win!")
			fmt.Printf("You made %d guesses.\n", g.guessCount)
			return
		}

		fmt.Printf("You have %d As, and %d Bs.\n", resultA, resultB)
		fmt.Println("Your guess isn't quite right, try again.")
	}
}

func main() {
	game := NewGame()
	game.Play()
}
