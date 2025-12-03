package puzzles

import (
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

func getLargestJoltage(bank string, joltage int) int {

	largest := make([]int, joltage)

	bankLength := len(bank)

	for i, num := range bank {
		n, err := strconv.Atoi(string(num))
		utils.Check(err)

		largestIndex := max(joltage-(bankLength-i), 0)

		for j := largestIndex; j < joltage; j++ {
			if n > largest[j] {
				largest[j] = n
				for k := j + 1; k < joltage; k++ {
					largest[k] = 0
				}
				break
			}
		}
	}

	var sb strings.Builder
	for _, x := range largest {
		sb.WriteString(strconv.Itoa(x))
	}
	res, err := strconv.Atoi(sb.String())
	utils.Check(err)

	return res
}

func TotalJoltage(banks []string) int {
	sum := 0
	for _, bank := range banks {
		v := getLargestJoltage(bank, 12)
		sum += v
	}
	return sum
}
