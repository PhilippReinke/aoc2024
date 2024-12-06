package shared

import (
	"strconv"
	"strings"
)

type Number interface {
	int | int64
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func ParseSeperatedNums(s, sep string) []int {
	var nums []int
	numsString := strings.Split(s, sep)
	for _, numString := range numsString {
		num, err := strconv.ParseInt(numString, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(num))
	}
	return nums
}
