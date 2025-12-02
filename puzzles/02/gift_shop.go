package puzzles

import (
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

func isSymmetric(n int) bool {
	// puzzle one
	numString := strconv.Itoa(n)
	if len(numString)%2 == 1 {
		return false
	}
	i := len(numString) / 2
	return numString[:i] == numString[i:]
}

func isRepeating(n int) bool {
	// puzzle two
	num := strconv.Itoa(n)

	for size := 1; size < 1+(len(num)/2); size++ {
		group := num[:size]

		if len(num)%size != 0 {
			continue
		}

		repeating := true
		for i := size; i < len(num); i += size {
			g := num[i : i+size]
			if g != group {
				repeating = false
				break
			}
		}
		if repeating {
			return true
		}
	}
	return false
}

func formatInput(s string) []string {
	return strings.Split(s, ",")
}

func SumInvalidID(s string) int {

	res := 0

	idRanges := formatInput(s)
	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")
		l, r := ids[0], ids[1]

		left, err := strconv.Atoi(l)
		utils.Check(err)

		right, err := strconv.Atoi(r)
		utils.Check(err)

		for left <= right {
			if isRepeating(left) {
				res += left
			}
			left++
		}
	}

	return res
}
