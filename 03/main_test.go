package main

import (
	"testing"
)

func testData() string {
	return `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
}

func TestEngineCreation(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	if len(e.maskValueMap) != 10 {
		t.Errorf(`Item was %d long`, len(e.maskValueMap))
	}
}

func TestSymbolPositions(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	sp := findSymbolPositions(e, false)
	if len(sp) != 6 {
		t.Errorf(`Item was %d long`, len(sp))
	}
}

func TestFindSymbolAdjacentMasks(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	m := findSymbolAdjacentMasks(e, [2]int{1, 3})
	if len(m) != 2 {
		t.Errorf(`Item was %d long`, len(m))
	}
}

func TestFindAllSymbolAdjacentNums(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	sp := findSymbolPositions(e, false)
	v := findAllSymbolAdjacentNums(e, sp)
	if len(v) != 8 {
		t.Errorf(`Item was %d long`, len(v))
	}
}

func TestGetPower(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	power := getPower(e)
	if power != 4361 {
		t.Errorf(`Power was %d`, power)
	}
}

func TestFindGears(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	gearPositions := findSymbolPositions(e, true)
	if len(gearPositions) != 3 {
		t.Errorf(`Item was %d long`, len(gearPositions))
	}
}

func TestFindValidGears(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	validGearPositions := findValidGears(e)
	if len(validGearPositions) != 2 {
		t.Errorf(`Item was %d long`, len(validGearPositions))
	}
}

func TestGetGearRatio(t *testing.T) {
	input := testData()
	e := parseEngine(input)
	gearRatio := getGearRatio(e)
	if gearRatio != 467835 {
		t.Errorf(`ratio was %d`, gearRatio)
	}
}
