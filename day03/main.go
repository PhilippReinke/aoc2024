package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	path := flag.String("input", "input.txt", "path to puzzle input")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	line := strings.Replace(string(data), "\n", "", -1)

	var sol1, sol2 int
	reMul := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	reDo := regexp.MustCompile("do\\(\\)")
	reDont := regexp.MustCompile("don't\\(\\)")

	matchesMul := reMul.FindAllStringIndex(line, -1)
	matchesDo := reDo.FindAllStringIndex(line, -1)
	matchesDont := reDont.FindAllStringIndex(line, -1)

	doMul := func(mulStart int) bool {
		var mostRecentDo, mostRecentDont int
		for _, indices := range slices.Backward(matchesDo) {
			doStart := indices[0]
			if mulStart >= doStart {
				mostRecentDo = doStart
				break
			}
		}
		for _, indices := range slices.Backward(matchesDont) {
			dontStart := indices[0]
			if mulStart >= dontStart {
				mostRecentDont = dontStart
				break
			}
		}
		if mostRecentDo >= mostRecentDont {
			return true
		}
		return false
	}

	for _, indices := range matchesMul {
		start := indices[0]
		end := indices[1]
		match := line[start:end]

		var num1, num2 int
		_, err := fmt.Sscanf(match, "mul(%d,%d)", &num1, &num2)
		if err != nil {
			panic(err)
		}
		sol1 += num1 * num2

		if doMul(start) {
			sol2 += num1 * num2
		}
	}

	fmt.Println("Solution I:", sol1)
	fmt.Println("Solution II:", sol2)
}
