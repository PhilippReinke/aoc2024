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
	split := strings.Split(string(data), "\n\n")
	rules := strings.Split(split[0], "\n")
	updates := strings.Split(split[1], "\n")

	// parse rules
	var rulesParsed [][]int
	for _, rule := range rules {
		if rule == "" {
			continue
		}
		var num1, num2 int
		_, err := fmt.Sscanf(rule, "%d|%d", &num1, &num2)
		if err != nil {
			panic(err)
		}
		rulesParsed = append(rulesParsed, []int{num1, num2})
	}

	var sol1, sol2 int
	for _, update := range updates {
		if update == "" {
			continue
		}

		nums := shared.ParseSeperatedNums(update, ",")

		// check rules
		var ruleViolated bool
		for _, rule := range rulesParsed {
			firstPos := -1
			secondPos := -1

			// determine position of nums
			for i, num := range nums {
				if num == rule[0] {
					firstPos = i
				}
				if num == rule[1] {
					secondPos = i
				}
			}

			if (firstPos != -1 && secondPos != -1) && secondPos < firstPos {
				ruleViolated = true
				break
			}
		}

		if !ruleViolated {
			sol1 += nums[(len(nums)+1)/2-1]
		} else {
			numsCpy := make([]int, len(nums))
			copy(numsCpy, nums)
			numsCorrected := applyAllRules(numsCpy, rulesParsed)
			sol2 += numsCorrected[(len(numsCorrected)+1)/2-1]
		}
	}

	fmt.Println("Solution I:", sol1)
	fmt.Println("Solution II:", sol2)
}

func applyAllRules(nums []int, rules [][]int) []int {
	var allRulesApply bool
	for !allRulesApply {
		var ruleViolated bool
		for _, rule := range rules {
			firstPos := -1
			secondPos := -1

			// determine position of nums
			for i, num := range nums {
				if num == rule[0] {
					firstPos = i
				}
				if num == rule[1] {
					secondPos = i
				}
			}

			if (firstPos != -1 && secondPos != -1) && secondPos < firstPos {
				ruleViolated = true
				nums[firstPos], nums[secondPos] = nums[secondPos], nums[firstPos]
			}
		}
		if !ruleViolated {
			allRulesApply = true
		}
	}
	return nums
}
