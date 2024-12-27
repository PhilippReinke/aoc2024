package main

import (
	"flag"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"

	"github.com/PhilippReinke/aoc2024/shared"
)

const (
	plus  = '+'
	times = '*'
	pipe  = '|'
)

func main() {
	path := flag.String("input", "input.txt", "path to puzzle input")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	var sol1, sol2 int
	for _, line := range lines {
		if line == "" {
			continue
		}

		colon := strings.Index(line, ":")
		if colon == -1 {
			panic("colon not found")
		}

		testVal := shared.ParseSeperatedNums(line[:colon], " ")[0]
		nums := shared.ParseSeperatedNums(line[colon+2:], " ")

		// part I
		for operators := range generateCombinations(len(nums)-1, plus, times) {
			result := nums[0]
			for i, r := range operators {
				switch r {
				case plus:
					result += nums[i+1]
				case times:
					result *= nums[i+1]
				}
			}
			if result == testVal {
				sol1 += testVal
				break
			}
		}

		// part II
		for operators := range generateCombinations(len(nums)-1, plus, times, pipe) {
			result := nums[0]
			for i, r := range operators {
				switch r {
				case plus:
					result += nums[i+1]
				case times:
					result *= nums[i+1]
				case pipe:
					resultStr := strconv.Itoa(result)
					numStr := strconv.Itoa(nums[i+1])
					combined := resultStr + numStr
					result, err = strconv.Atoi(combined)
					if err != nil {
						panic("pipe failed")
					}
				}
			}
			if result == testVal {
				sol2 += testVal
				break
			}
		}
	}

	fmt.Println("Solution I:", sol1)
	fmt.Println("Solution II:", sol2)
}

func generateCombinations(length int, operators ...rune) iter.Seq[[]rune] {
	return func(yield func(s []rune) bool) {
		indices := make([]int, length)

		for {
			var opRunes []rune
			for _, i := range indices {
				opRunes = append(opRunes, operators[i])
			}
			if !yield(opRunes) {
				return
			}
			if finished(indices, len(operators)) {
				break
			}

			// next indices
			currentIdx, carry := 0, 1
			for carry != 0 {
				if indices[currentIdx] == len(operators)-1 {
					indices[currentIdx] = 0
					currentIdx++
					continue
				}
				indices[currentIdx] += carry
				carry = 0
			}
		}
	}
}

func finished(indices []int, numOfOperators int) bool {
	for _, i := range indices {
		if i != numOfOperators-1 {
			return false
		}
	}
	return true
}
