package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestUnits(t *testing.T) {
	r := *parseRound("3 blue, 5 red")
	if len(r) != 2 {
		t.Errorf(`Item was %d long`, len(r))
	}

	g := *parseGame(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`)
	if g.num != 1 {
		t.Errorf(`Item Number was %d long`, g.num)
	}
	if len(g.rounds) != 3 {
		t.Errorf(`Item was %d long`, len(r))
	}
	if g.rounds[0]["blue"] != 3 {
		t.Errorf(`Blue counts was %d`, g.rounds[0]["blue"])
	}
}

func TestPart1(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	lines := strings.Split(input, "\n")
	var games []Game

	for _, line := range lines {
		games = append(games, *parseGame(line))
	}

	if len(games) != 5 {
		t.Fatalf(`Wrong number of games %d`, len(games))
	}
	if !gameValid(games[0], maxCubes) {
		t.Fatalf(`Game %q should be valid`, games[0])
	}
	if gameValid(games[2], maxCubes) {
		t.Fatalf(`Game %q should be NOT valid`, games[0])
	}

	if sumValidGames(games, maxCubes) != 8 {
		t.Fatalf(`Games %q should have added up to %d`, games, 8)
	}

	val := map[string]int{"red": 4, "green": 2, "blue": 6}
	if !reflect.DeepEqual(fewestCubes(games[0]), val) {
		t.Fatalf("Expected %q but got %q\n", val, fewestCubes(games[0]))
	}

	if gamePower(games[0]) != 48 {
		t.Fatalf("Expected power of %d but got %d\n", 48, gamePower(games[0]))
	}
}
