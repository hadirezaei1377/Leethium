package hads_bezan

import (
	"fmt"
	"math/rand"
	"time"
)

type Game interface {
	CheckNumber(n int) string
}

func GuessMyNumber(game Game) string {
	// initial guess
	guess := 180
	maxAttempts := 10

	for i := 0; i < maxAttempts; i++ {
		result := game.CheckNumber(guess)
		if result == "CORRECT" {
			return fmt.Sprintf("Your Number was %d", guess)
		} else if result == "My Number is Greater" {
			guess -= 30
		} else {
			guess += 30
		}
	}

	return "Failed to find the number within the maximum attempts"
}

func (g *GameSample) CheckNumber(n int) string {
	if n == g.Number {
		return "CORRECT"
	} else if n < g.Number {
		return "My Number is Greater"
	} else {
		return "My Number is Lower"
	}
}

type GameSample struct {
	Number int
}

func main() {

	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(360) + 1

	game := &GameSample{Number: number}

	result := GuessMyNumber(game)

	fmt.Println(result)
}
