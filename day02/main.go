package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	var cntSafeReport int
	var cntSafeReport2 int
	for _, line := range lines {
		if line == "" {
			continue
		}

		var nums []int64
		numStrings := strings.Split(line, " ")
		for _, numString := range numStrings {
			num, _ := strconv.ParseInt(numString, 10, 64)
			nums = append(nums, num)
		}

		if isSafe(nums) {
			cntSafeReport++
		}

		if isSafe2(nums) {
			cntSafeReport2++
		}
	}

	fmt.Println("Solution I:", cntSafeReport)
	fmt.Println("Solution II:", cntSafeReport2)
}

func isSafe(nums []int64) bool {
	isIncreasing := nums[1]-nums[0] > 0

	for i := 0; i < len(nums)-1; i++ {
		dist := nums[i+1] - nums[i]
		if !(1 <= shared.Abs(dist) && shared.Abs(dist) <= 3) {
			return false
		}
		if (dist > 0) != isIncreasing {
			return false
		}
	}
	return true
}

func isSafe2(nums []int64) bool {
	for i := range nums {
		tempNums := make([]int64, 0, len(nums)-1)
		tempNums = append(tempNums, nums[:i]...)
		tempNums = append(tempNums, nums[i+1:]...)

		if isSafe(tempNums) {
			return true
		}
	}
	return false
}
