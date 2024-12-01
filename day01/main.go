package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
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

	numsLeft := make([]int, 0, len(lines))
	numsRight := make([]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		var num1, num2 int
		_, err := fmt.Sscanf(line, "%d   %d", &num1, &num2)
		if err != nil {
			panic(err)
		}

		numsLeft = append(numsLeft, num1)
		numsRight = append(numsRight, num2)
	}

	sort.Ints(numsLeft)
	sort.Ints(numsRight)
	var difference int
	seen := make(map[int]int, len(numsLeft))

	for i := range len(numsLeft) {
		difference += shared.Abs(numsLeft[i] - numsRight[i])

		_, ok := seen[numsRight[i]]
		if !ok {
			seen[numsRight[i]] = 1
		} else {
			seen[numsRight[i]]++
		}
	}

	var similarityScore int
	for _, num := range numsLeft {
		similarityScore += num * seen[num]
	}

	fmt.Println("Solution I:", difference)
	fmt.Println("Solution II:", similarityScore)
}
