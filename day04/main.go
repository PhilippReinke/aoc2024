package main

import (
	"flag"
	"fmt"
	"iter"
	"os"
	"strings"
)

const (
	XMAS = "XMAS"
	SAMX = "SAMX"
	MAS  = "MAS"
	SAM  = "SAM"
)

func main() {
	path := flag.String("input", "input.txt", "path to puzzle input")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	var runes [][]rune
	for _, line := range lines {
		if line == "" {
			continue
		}
		runes = append(runes, []rune(line))
	}

	var sol1 int

	// horizontal
	for r := range horizontalIterator(runes) {
		sol1 += strings.Count(string(r), XMAS)
		sol1 += strings.Count(string(r), SAMX)
	}

	// vertical
	for r := range verticalIterator(runes) {
		sol1 += strings.Count(string(r), XMAS)
		sol1 += strings.Count(string(r), SAMX)
	}

	// diagonal top-left to bottom-right
	for r := range diagonalPrimaryIterator(runes) {
		sol1 += strings.Count(string(r), XMAS)
		sol1 += strings.Count(string(r), SAMX)
	}

	// diagonal top-right to bottom-left
	for r := range diagonalSecondaryIterator(runes) {
		sol1 += strings.Count(string(r), XMAS)
		sol1 += strings.Count(string(r), SAMX)
	}

	fmt.Println("Solution I:", sol1)
	fmt.Println("Solution II:", part2(runes))
}

func horizontalIterator(runes [][]rune) iter.Seq[[]rune] {
	return func(yield func(r []rune) bool) {
		for row := 0; row < len(runes); row++ {
			if !yield(runes[row]) {
				return
			}
		}
	}
}

func verticalIterator(runes [][]rune) iter.Seq[[]rune] {
	return func(yield func(r []rune) bool) {
		for col := 0; col < len(runes[0]); col++ {
			var colRunes []rune
			for row := 0; row < len(runes); row++ {
				colRunes = append(colRunes, runes[row][col])
			}
			if !yield(colRunes) {
				return
			}
		}
	}
}

func diagonalPrimaryIterator(runes [][]rune) iter.Seq[[]rune] {
	return func(yield func(r []rune) bool) {
		for startRow := len(runes) - 1; startRow >= 0; startRow-- {
			var diagonalRunes []rune

			for steps := 0; startRow+steps < len(runes); steps++ {
				diagonalRunes = append(diagonalRunes, runes[startRow+steps][steps])
			}

			if !yield(diagonalRunes) {
				return
			}
		}
		for startCol := 1; startCol < len(runes[0]); startCol++ {
			var diagonalRunes []rune

			for steps := 0; startCol+steps < len(runes[0]); steps++ {
				diagonalRunes = append(diagonalRunes, runes[steps][startCol+steps])
			}

			if !yield(diagonalRunes) {
				return
			}
		}
	}
}

func diagonalSecondaryIterator(runes [][]rune) iter.Seq[[]rune] {
	return func(yield func(r []rune) bool) {
		for startRow := len(runes) - 1; startRow >= 0; startRow-- {
			var diagonalRunes []rune

			for steps := 0; startRow+steps < len(runes); steps++ {
				diagonalRunes = append(diagonalRunes, runes[startRow+steps][len(runes)-1-steps])
			}

			if !yield(diagonalRunes) {
				return
			}
		}
		for startCol := len(runes) - 2; startCol >= 0; startCol-- {
			var diagonalRunes []rune

			for steps := 0; startCol-steps >= 0; steps++ {
				diagonalRunes = append(diagonalRunes, runes[steps][startCol-steps])
			}

			if !yield(diagonalRunes) {
				return
			}
		}
	}
}

func part2(runes [][]rune) int {
	// sadly, my fancy iterator approach didn't help with part 2. so fuck it and
	// do it like this
	var cnt int

	for row := 1; row < len(runes)-1; row++ {
		for col := 1; col < len(runes)-1; col++ {
			if runes[row][col] != 'A' {
				continue
			}
			dia1 := string([]rune{runes[row-1][col-1], runes[row][col], runes[row+1][col+1]})
			dia2 := string([]rune{runes[row-1][col+1], runes[row][col], runes[row+1][col-1]})
			if (dia1 == MAS || dia1 == SAM) && (dia2 == MAS || dia2 == SAM) {
				cnt++
			}
		}
	}

	return cnt
}
