package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toInt(input string) int {
	switch input {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		val, err := strconv.Atoi(input)
		check(err)
		return val
	}
}

func extractDigitPair(input string, part1_mode bool) int {
	var tenstr string
	var onestr string

	if part1_mode {
		pattern := `\d`
		re, err := regexp.Compile(pattern)
		check(err)
		matches := re.FindAllString(input, -1)
		tenstr = matches[0]
		onestr = matches[len(matches)-1]
	} else {
		numwords := [19]string{
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		}
		pos := -1
		word := ""
		for i := 0; i < len(numwords); i++ {
			found := strings.Index(input, numwords[i])
			if pos == -1 || (found > -1 && found < pos) {
				pos = found
				word = numwords[i]
			}
		}
		tenstr = word

		pos = -1
		word = ""
		for i := 0; i < len(numwords); i++ {
			found := strings.LastIndex(input, numwords[i])
			if pos == -1 || (found > -1 && found > pos) {
				pos = found
				word = numwords[i]
			}
		}
		onestr = word
	}

	tens := toInt(tenstr)
	ones := toInt(onestr)

	return 10*tens + ones
}

func main() {
	input_fname := "./input.txt"

	input, err := os.ReadFile(input_fname)
	check(err)
	lines := strings.Split(string(input), "\n")

	total := 0
	for i := 0; i < len(lines); i++ {
		total += extractDigitPair(lines[i], true)
	}
	fmt.Printf("Part 1 Total: %d\n", total)

	total = 0
	for i := 0; i < len(lines); i++ {
		total += extractDigitPair(lines[i], false)
	}
	fmt.Printf("Part 2 Total: %d\n", total)
}
