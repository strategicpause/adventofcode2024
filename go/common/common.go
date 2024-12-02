package common

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func RunAndMeasure(part string, f func(string) int, fileName string) int {
	var input, err = ReadInput(fileName)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	answer := f(input)
	duration := time.Since(now).Microseconds()
	fmt.Printf("Part %s (%s): %d (%d us)\n", part, fileName, answer, duration)

	return answer
}

func RunAndAssert(expected int, f func(string) int, fileName string) {
	var input, err = ReadInput(fileName)
	if err != nil {
		panic(err)
	}
	answer := f(input)
	if answer != expected {
		panic(fmt.Sprintf("Expected %d but got %d", expected, answer))
	}
}

func ReadInput(name string) (string, error) {
	path := fmt.Sprintf("../../data/%s", name)
	content, err := os.ReadFile(path)
	return string(content), err
}

func Atoi(str string) int {
	n := 0
	l := len(str)
	for i := 0; i < l; i++ {
		n = n*10 + CharAtoi(str[i])
	}
	return n
}

func CharAtoi(c byte) int {
	return int(c - '0')
}

// SplitAtoi will split the given string by the given split character.
// Each of the resulting elements will then be converted to an integer.
func SplitAtoi(str string, splitChar byte) []int {
	var nums []int
	SplitItr(str, splitChar, func(s string) {
		if s != "" {
			nums = append(nums, Atoi(strings.TrimSpace(s)))
		}
	})
	return nums
}

func SplitLines(str string, f func(string)) {
	SplitItr(str, '\n', f)
}

// SplitItr will split the given string by the given split character.
// The resulting elements will passed as a parameter to the given
// function.
func SplitItr(str string, splitChar byte, f func(string)) {
	strs := strings.Split(str, string(splitChar))
	for _, str := range strs {
		if str != "" {
			f(strings.TrimSpace(str))
		}
	}
}

func Zip(first []int, second []int, f func(int, int) int) []int {
	length := Max(len(first), len(second))
	var zipped []int
	for i := 0; i < length; i++ {
		result := f(first[i], second[i])
		zipped = append(zipped, result)
	}
	return zipped
}

func Sum[T int](window []T) T {
	var sum T
	for _, n := range window {
		sum += n
	}
	return sum
}

func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
