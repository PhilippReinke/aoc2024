package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	up          = '^'
	obstruction = '#'
)

func main() {
	path := flag.String("input", "input.txt", "path to puzzle input")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	// make 2d rune array
	var runes [][]rune
	var rowStart, colStart int
	for row, line := range lines {
		if line == "" {
			continue
		}
		col := strings.Index(line, string(up))
		if col != -1 {
			rowStart, colStart = row, col
		}
		runes = append(runes, []rune(line))
	}

	// part 2
	var sol2 int
	var prevRune rune
	for r := range len(runes[0]) {
		for c := range len(runes) {
			prevRune = runes[r][c]
			runes[r][c] = obstruction

			if letGuardWalk(rowStart, colStart, runes) == -1 {
				sol2++
			}

			runes[r][c] = prevRune
		}
	}

	fmt.Println("Solution I:", letGuardWalk(rowStart, colStart, runes))
	fmt.Println("Solution II:", sol2)
}

func isRock(row, col int, runes [][]rune) bool {
	if !isInside(row, col, runes) {
		return false
	}
	return runes[row][col] == '#'
}

func isInside(row, col int, runes [][]rune) bool {
	if row < 0 || row >= len(runes[0]) || col < 0 || col >= len(runes) {
		return false
	}
	return true
}

func letGuardWalk(row, col int, runes [][]rune) int {
	visited := make(map[[2]int]struct{})
	visitedWithDirection := make(map[[4]int]struct{})

	dir := [2]int{-1, 0}

	visited[[2]int{row, col}] = struct{}{}
	visitedWithDirection[[4]int{row, col, dir[0], dir[1]}] = struct{}{}

	for {
		rowNext := row + dir[0]
		colNext := col + dir[1]
		if !isRock(rowNext, colNext, runes) {
			// move player
			row = rowNext
			col = colNext
		} else {
			// turn 90 degree right
			dir = [2]int{dir[1], -dir[0]}
		}

		if !isInside(row, col, runes) {
			break
		}
		visited[[2]int{row, col}] = struct{}{}

		// check for loops
		_, ok := visitedWithDirection[[4]int{row, col, dir[0], dir[1]}]
		if ok {
			return -1
		}
		visitedWithDirection[[4]int{row, col, dir[0], dir[1]}] = struct{}{}
	}

	return len(visited)
}
