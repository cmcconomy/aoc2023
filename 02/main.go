package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	num    int              // Game #
	rounds []map[string]int // Each round is a map of colorname->count
}

// ------------------------------------------
// Parsing
// ------------------------------------------

func parseRound(input string) *map[string]int {
	movemap := make(map[string]int)

	moves := strings.Split(input, ", ")
	for i := 0; i < len(moves); i++ {
		move := strings.Split(moves[i], " ")
		num, err := strconv.Atoi(move[0])
		check(err)
		movemap[move[1]] = num
	}

	return &movemap
}

func parseGame(input string) *Game {
	parts := strings.Split(input, ": ")
	gameNum, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	check(err)
	round_strs := strings.Split(parts[1], "; ")

	var rounds []map[string]int
	for i := 0; i < len(round_strs); i++ {
		rounds = append(rounds, *parseRound(round_strs[i]))
	}

	game := Game{
		num:    gameNum,
		rounds: rounds,
	}
	return &game
}

// ------------------------------------------
// Part 1
// ------------------------------------------

func gameValid(game Game, maxCubes map[string]int) bool {
	for color, max := range maxCubes {
		for _, round := range game.rounds {
			if round[color] > max {
				return false
			}
		}
	}
	return true
}

func sumValidGames(games []Game, maxCubes map[string]int) int {
	sum := 0
	for _, game := range games {
		if gameValid(game, maxCubes) {
			sum += game.num
		}
	}
	return sum
}

// ------------------------------------------
// Part 2
// ------------------------------------------

func fewestCubes(game Game) map[string]int {
	minmoves := make(map[string]int)
	for _, round := range game.rounds {
		for color, num := range round {
			if num > minmoves[color] {
				minmoves[color] = num
			}
		}
	}

	return minmoves
}

func gamePower(game Game) int {
	power := 1
	for _, val := range fewestCubes(game) {
		power *= val
	}
	return power
}

func sumGamePower(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += gamePower(game)
	}
	return sum
}

func main() {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	input_fname := "./input.txt"

	input, err := os.ReadFile(input_fname)
	check(err)
	lines := strings.Split(string(input), "\n")

	var games []Game
	for _, line := range lines {
		games = append(games, *parseGame(line))
	}

	fmt.Printf("Part 1: %d\n", sumValidGames(games, maxCubes))
	fmt.Printf("Part 2: %d\n", sumGamePower(games))
}
