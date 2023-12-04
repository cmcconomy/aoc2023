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

type Engine struct {
	symbols      []string    // the actual engine view
	mask         [][]int     // rows of masks.. nondigits are -1, digits will be part of group 0, group 1, group 2, group 3...
	maskValueMap map[int]int // map of
}

// ------------------------------------------
// Parsing
// ------------------------------------------

func parseEngine(input string) *Engine {
	symbols := strings.Split(input, "\n")
	var mask [][]int
	maskValueMap := make(map[int]int)

	masknum := 0
	start := -1
	for _, line := range symbols {
		var linemask []int
		for i, char := range []byte(line) {
			if char > 47 && char < 58 { // digit
				if start == -1 {
					// starting a number
					start = i
				}
				linemask = append(linemask, masknum)
			} else {
				if start != -1 {
					// ending a number because the current character is non-digit
					num, err := strconv.Atoi(line[start:i])
					check(err)
					maskValueMap[masknum] = num
					masknum++
					start = -1
				}
				linemask = append(linemask, -1)
			}
		}
		if start != -1 {
			// ending a number with the very last character
			num, err := strconv.Atoi(line[start:])
			check(err)
			maskValueMap[masknum] = num
			masknum++
			start = -1
		}
		mask = append(mask, linemask)
	}

	return &Engine{
		symbols:      symbols,
		mask:         mask,
		maskValueMap: maskValueMap,
	}
}

// ------------------------------------------
// Part 1
// ------------------------------------------

func findSymbolPositions(engine *Engine, gearsOnly bool) [][2]int {
	const period byte = 46
	var positions [][2]int

	for row, line := range engine.symbols {
		for col, char := range []byte(line) {
			var isSymbol bool
			if gearsOnly {
				isSymbol = char == 42
			} else {
				isSymbol = char != period && char < 48 || char > 57
			}

			if isSymbol {
				positions = append(positions, [2]int{row, col})
			}
		}
	}

	return positions
}

func findSymbolAdjacentMasks(engine *Engine, symbolPosition [2]int) []int {
	adjacentMaskNums := make(map[int]bool)

	rowstart := max(0, symbolPosition[0]-1)
	rowend := min(len(engine.symbols[0]), symbolPosition[0]+2)
	colstart := max(0, symbolPosition[1]-1)
	colend := min(len(engine.symbols), symbolPosition[1]+2)
	for row := rowstart; row < rowend; row++ {
		for col := colstart; col < colend; col++ {
			mask := engine.mask[row][col]
			if mask > -1 {
				adjacentMaskNums[mask] = true
			}
		}
	}

	masks := make([]int, 0, len(adjacentMaskNums))
	for mask, _ := range adjacentMaskNums {
		masks = append(masks, mask)
	}

	return masks
}

func findAllSymbolAdjacentNums(engine *Engine, symbolPositions [][2]int) []int {
	adjacentMaskNums := make(map[int]bool)
	for _, symbolPos := range symbolPositions {
		masks := findSymbolAdjacentMasks(engine, symbolPos)
		for _, mask := range masks {
			adjacentMaskNums[mask] = true
		}
	}

	values := make([]int, 0, len(adjacentMaskNums))
	for mask, _ := range adjacentMaskNums {
		value := engine.maskValueMap[mask]
		values = append(values, value)
	}

	return values
}

func getPower(engine *Engine) int {
	sp := findSymbolPositions(engine, false)
	values := findAllSymbolAdjacentNums(engine, sp)
	power := 0
	for _, value := range values {
		power += value
	}

	return power
}

// ------------------------------------------
// Part 2
// ------------------------------------------

func findValidGears(engine *Engine) [][2]int {
	var validGears [][2]int
	gearPositions := findSymbolPositions(engine, true)
	for _, gearPos := range gearPositions {
		masks := findSymbolAdjacentMasks(engine, gearPos)
		if len(masks) == 2 {
			validGears = append(validGears, gearPos)
		}
	}

	return validGears
}

func getGearRatio(engine *Engine) int {
	gearRatio := 0
	validGearPositions := findValidGears(engine)
	for _, validGearPos := range validGearPositions {
		masks := findSymbolAdjacentMasks(engine, validGearPos)
		ratio := engine.maskValueMap[masks[0]] * engine.maskValueMap[masks[1]]
		gearRatio += ratio
	}

	return gearRatio
}

// ------------------------------------------
// Main
// ------------------------------------------

func main() {
	input_fname := "./input.txt"
	input, err := os.ReadFile(input_fname)
	check(err)

	engine := parseEngine(string(input))
	fmt.Printf("Part 1: %d\n", getPower(engine))
	fmt.Printf("Part 2: %d\n", getGearRatio(engine))
}
