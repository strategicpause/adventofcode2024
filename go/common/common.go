package common

import (
	"fmt"
	"iter"
	"os"
	"strings"
	"time"
)

func RunAndMeasure[T int](part string, f func(string) T, fileName string) T {
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

func RunAndAssert[T int](expected T, f func(string) T, fileName string) {
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
	for s := range SplitItr(str, splitChar) {
		if s != "" {
			nums = append(nums, Atoi(strings.TrimSpace(s)))
		}
	}
	return nums
}

func SplitLines(str string) iter.Seq[string] {
	return SplitItr(str, '\n')
}

func SplitItr(str string, splitChar byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		splitStrs := strings.Split(str, string(splitChar))
		for _, s := range splitStrs {
			if s != "" {
				if !yield(strings.TrimSpace(s)) {
					return
				}
			}
		}
	}
}

func Zip[T, U, V any](first []T, second []U, f func(T, U) V) []V {
	length := Max(len(first), len(second))
	var zipped []V
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

func Abs[T int](n T) T {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Max[T int](n, m T) T {
	if n > m {
		return n
	}
	return m
}

func All[T any](s iter.Seq[T], f func(T) bool) bool {
	for v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[T any](s iter.Seq[T], f func(T) bool) bool {
	for v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func Range(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Remove[T any](vals []T, i int) []T {
	s := make([]T, len(vals))
	copy(s, vals)
	return append(s[:i], s[i+1:]...)
}

func Window[T any](s []T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		l := len(s)
		if size > l {
			yield(s)
			return
		}
		for i := 0; i <= l-size; i++ {
			windowStart := i
			windowEnd := i + size
			if windowEnd > l {
				windowEnd = l
			}
			if !yield(s[windowStart:windowEnd]) {
				return
			}
		}
	}
}
