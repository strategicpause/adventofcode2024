package main

import (
	"github.com/strategicpause/adventofcode2024/common"
)

func main() {
	common.RunAndAssert(2, PartA, "day02/sample.txt")
	common.RunAndAssert(4, PartB, "day02/sample.txt")
	// 287
	common.RunAndMeasure("A", PartA, "day02/input.txt")
	// 354
	common.RunAndMeasure("B", PartB, "day02/input.txt")
}

func PartA(input string) int {
	safe := 0
	for line := range common.SplitLines(input) {
		vals := common.SplitAtoi(line, ' ')
		if isSafe(vals) {
			safe += 1
		}
	}
	return safe
}

func isSafe(vals []int) bool {
	isSafe := common.All(common.Window(vals, 2), func(v []int) bool {
		if v[0] < v[1] {
			return false
		}
		return true
	})
	isSafe = isSafe || common.All(common.Window(vals, 2), func(v []int) bool {
		if v[0] > v[1] {
			return false
		}
		return true
	})

	return isSafe && common.All(common.Window(vals, 2), func(v []int) bool {
		return common.Abs(v[0]-v[1]) >= 1 && common.Abs(v[0]-v[1]) <= 3
	})
}

func isAnySafe(vals []int) bool {
	return common.Any(common.Range(0, len(vals)), func(v int) bool {
		return isSafe(common.Remove(vals, v))
	})
}

func PartB(input string) int {
	safe := 0
	for line := range common.SplitLines(input) {
		vals := common.SplitAtoi(line, ' ')
		if isAnySafe(vals) {
			safe += 1
		}
	}
	return safe
}
