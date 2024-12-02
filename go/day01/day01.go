package main

import (
	"github.com/strategicpause/adventofcode2024/common"
	"sort"
)

func main() {
	common.RunAndAssert(11, PartA, "day01/sample.txt")
	common.RunAndAssert(31, PartB, "day01/sample.txt")
	common.RunAndMeasure("A", PartA, "day01/input.txt")
	common.RunAndMeasure("B", PartB, "day01/input.txt")
}

func PartA(input string) int {
	var leftList []int
	var rightList []int

	for line := range common.SplitLines(input) {
		vals := common.SplitAtoi(line, ' ')
		leftList = append(leftList, vals[0])
		rightList = append(rightList, vals[1])
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return common.Sum(common.Zip(leftList, rightList, func(a, b int) int {
		return common.Abs(b - a)
	}))
}

func PartB(input string) int {
	var leftList []int
	rightMap := make(map[int]int)

	for line := range common.SplitLines(input) {
		vals := common.SplitAtoi(line, ' ')
		leftList = append(leftList, vals[0])
		rightMap[vals[1]] += 1
	}

	return common.Sum(common.Map(leftList, func(leftVal int) int {
		return leftVal * rightMap[leftVal]
	}))
}
