package puzzles

import (
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

// -------------------------------------------------------------------------------- helpers

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

func multiply(nums []int) int {
	res := 1
	for _, num := range nums {
		res *= num
	}
	return res
}

func solve(res *int, operator *string, nums []int) {
	switch *operator {
	case "+":
		*res += sum(nums)
	case "*":
		*res += multiply(nums)
	}
}

// -------------------------------------------------------------------------------- problems

func PartOne(worksheet []string) int {

	res := 0

	values := make([][]int, 0)

	for _, line := range worksheet[:len(worksheet)-1] {

		strNums := strings.Fields(line)
		row := make([]int, 0)

		for _, num := range strNums {
			n, err := strconv.Atoi(num)
			utils.Check(err)
			row = append(row, n)
		}
		values = append(values, row)
	}

	operators := strings.Fields(worksheet[len(worksheet)-1])

	for col, operator := range operators {

		problemNums := make([]int, 0)

		for row := range values {
			value := values[row][col]
			problemNums = append(problemNums, value)
		}

		solve(&res, &operator, problemNums)
	}

	return res
}

func PartTwo(worksheet []string) int {

	res := 0
	problemNums := make([]int, 0)

	i := 0
	operators := strings.Fields(worksheet[len(worksheet)-1])

	// construct number top-to-bottom
	for col := range worksheet[0] {

		var sb strings.Builder

		for row := range worksheet[:len(worksheet)-1] {
			value := string(worksheet[row][col])
			if _, err := strconv.Atoi(value); err == nil {
				sb.WriteString(value)
			}
		}

		num, err := strconv.Atoi(sb.String())
		if err == nil {
			problemNums = append(problemNums, num)
		} else {
			// collected all numbers in the problem
			solve(&res, &operators[i], problemNums)
			i += 1
			problemNums = make([]int, 0)
		}
	}

	solve(&res, &operators[i], problemNums)
	return res
}
