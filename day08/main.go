package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PhilippReinke/aoc2024/shared"
)

func main() {
	path := flag.String("input", "input.txt", "path to puzzle input")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	var maxRowAndCol int
	typeToPos := make(map[rune][][2]int)
	for row, line := range lines {
		if line == "" {
			continue
		}

		if maxRowAndCol == 0 {
			maxRowAndCol = len(line)
		}

		for col, r := range line {
			if r != '.' {
				_, ok := typeToPos[r]
				if !ok {
					typeToPos[r] = [][2]int{{row, col}}
					continue
				}
				typeToPos[r] = append(typeToPos[r], [2]int{row, col})
			}
		}
	}

	// part 1
	sol1 := make(map[[2]int]struct{})
	for _, positions := range typeToPos {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				row1, col1 := positions[i][0], positions[i][1]
				row2, col2 := positions[j][0], positions[j][1]

				rowA1, colA1 := 2*row1-row2, 2*col1-col2
				if 0 <= rowA1 && rowA1 < maxRowAndCol && 0 <= colA1 && colA1 < maxRowAndCol {
					sol1[[2]int{rowA1, colA1}] = struct{}{}
				}

				rowA2, colA2 := 2*row2-row1, 2*col2-col1
				if 0 <= rowA2 && rowA2 < maxRowAndCol && 0 <= colA2 && colA2 < maxRowAndCol {
					sol1[[2]int{rowA2, colA2}] = struct{}{}
				}
			}
		}
	}

	// part 2
	sol2 := make(map[[2]int]struct{})
	for _, positions := range typeToPos {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				row1, col1 := positions[i][0], positions[i][1]
				row2, col2 := positions[j][0], positions[j][1]

				rowVec, colVec := row2-row1, col2-col1
				gcdVec := gcd(rowVec, colVec)
				rowVecNor, colVecNor := rowVec/gcdVec, colVec/gcdVec

				// minus direction
				var mul int
				for {
					rowNew, colNew := row1+mul*rowVecNor, col1+mul*colVecNor
					if 0 <= rowNew && rowNew < maxRowAndCol && 0 <= colNew && colNew < maxRowAndCol {
						sol2[[2]int{rowNew, colNew}] = struct{}{}
						mul--
					} else {
						break
					}
				}

				// plus direction
				mul = 0
				for {
					rowNew, colNew := row1+mul*rowVecNor, col1+mul*colVecNor
					if 0 <= rowNew && rowNew < maxRowAndCol && 0 <= colNew && colNew < maxRowAndCol {
						sol2[[2]int{rowNew, colNew}] = struct{}{}
						mul++
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println("Solution I:", len(sol1))
	fmt.Println("Solution II:", len(sol2))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return shared.Abs(a)
}
