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
	line := strings.Split(string(data), "\n")[0]

	// parse numbers to int slice
	nums := shared.ParseSeperatedNums(line, "")

	var expanded, expandedBlanks []int
	files := make(map[int]Block)
	var blanks []Block

	var curFileID, curIdx int
	for i, n := range nums {
		isFile := i%2 == 0

		if isFile {
			files[curFileID] = Block{curIdx, n}
			for range n {
				expanded = append(expanded, curFileID)
			}
			curFileID++
		}

		if !isFile && n > 0 {
			blanks = append(blanks, Block{curIdx, n})
			for i := range n {
				expanded = append(expanded, -1)
				expandedBlanks = append(expandedBlanks, curIdx+i)
			}
		}

		curIdx += n
	}

	// part 1
	for _, blankIdx := range expandedBlanks {
		for expanded[len(expanded)-1] == -1 {
			expanded = expanded[:len(expanded)-1]
		}
		if blankIdx >= len(expanded)-1 {
			break
		}
		expanded[blankIdx] = expanded[len(expanded)-1]
		expanded = expanded[:len(expanded)-1]
	}

	var checksumPart1 int
	for i, num := range expanded {
		checksumPart1 += i * num
	}

	// part 2
	for curFileID > 0 {
		curFileID--
		file := files[curFileID]

		for i, blank := range blanks {
			if blank.Idx >= file.Idx {
				// we are past the current file and drop the blank
				blanks = blanks[:i]
				break
			}
			if file.Len <= blank.Len {
				// file fits in blank
				files[curFileID] = Block{blank.Idx, file.Len}

				if file.Len == blank.Len {
					// drop blank
					blanks = append(blanks[:i], blanks[i+1:]...)
					break
				} else {
					blanks[i] = Block{blank.Idx + file.Len, blank.Len - file.Len}
				}
				break
			}
		}
	}

	var checksumPart2 int
	for fileID, file := range files {
		for i := range file.Len {
			checksumPart2 += fileID * (file.Idx + i)
		}
	}

	fmt.Println("Solution I:", checksumPart1)
	fmt.Println("Solution II:", checksumPart2)
}

type Block struct {
	Idx int
	Len int
}
