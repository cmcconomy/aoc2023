package main

import (
	"strings"
	"testing"
)

func TestExtractDigitPairPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	lines := strings.Split(input, "\n")
	if len(lines) != 4 {
		t.Fatalf(`Wrong number of lines %d`, len(lines))
	}

	total := 0
	for i := 0; i < len(lines); i++ {
		total += extractDigitPair(lines[i], true)
	}
	if total != 142 {
		t.Fatalf(`Sum should be %d but got %d`, 142, total)
	}
}

func TestExtractDigitPairPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	lines := strings.Split(input, "\n")
	if len(lines) != 7 {
		t.Fatalf(`Wrong number of lines %d`, len(lines))
	}

	total := 0
	for i := 0; i < len(lines); i++ {
		total += extractDigitPair(lines[i], false)
	}
	if total != 281 {
		t.Fatalf(`Sum should be %d but got %d`, 281, total)
	}
}
