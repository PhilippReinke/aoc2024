package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	numRowsCols = 0
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
	var zeroHeights []pos
	for row, line := range lines {
		if line == "" {
			continue
		}
		if numRowsCols == 0 {
			numRowsCols = len(line)
		}

		var runesRow []rune
		for col, r := range line {
			runesRow = append(runesRow, r)
			if r == '0' {
				zeroHeights = append(zeroHeights, pos{row, col})
			}
		}
		runes = append(runes, runesRow)
	}

	var sol1, sol2 int
	for _, zeroPos := range zeroHeights {
		var stepCount int
		currentPositions := []pos{zeroPos}

		for len(currentPositions) > 0 {
			nextPositions := []pos{}

			for _, cur := range currentPositions {
				for _, next := range cur.nextPossiblePositions() {
					if int(runes[next.row][next.col])-int(runes[cur.row][cur.col]) == 1 {
						nextPositions = append(nextPositions, next)
					}
				}
			}
			currentPositions = nextPositions
			stepCount++

			if len(currentPositions) == 0 || stepCount == 9 {
				break
			}
		}

		uniqueCurrentPositions := make(map[pos]struct{})
		for _, cur := range currentPositions {
			uniqueCurrentPositions[cur] = struct{}{}
		}
		sol1 += len(uniqueCurrentPositions)
		sol2 += len(currentPositions)
	}

	fmt.Println("Solution I:", sol1)
	fmt.Println("Solution II:", sol2)
}

type pos struct {
	row, col int
}

func (p pos) nextPossiblePositions() []pos {
	var validNext []pos

	up := pos{p.row - 1, p.col}
	if up.valid() {
		validNext = append(validNext, up)
	}

	down := pos{p.row + 1, p.col}
	if down.valid() {
		validNext = append(validNext, down)
	}

	left := pos{p.row, p.col - 1}
	if left.valid() {
		validNext = append(validNext, left)
	}

	right := pos{p.row, p.col + 1}
	if right.valid() {
		validNext = append(validNext, right)
	}

	return validNext
}

func (p pos) valid() bool {
	if p.row < 0 || p.row >= numRowsCols {
		return false
	}
	if p.col < 0 || p.col >= numRowsCols {
		return false
	}
	return true
}
