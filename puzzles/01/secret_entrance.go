package puzzles

import (
	"strconv"

	"github.com/jda5/aoc2025/utils"
)

func CountZeros(input []string) int {

	pos := 50
	count := 0

	for _, line := range input {
		direction := string(line[0])
		stringDistance := string(line[1:])
		distance, err := strconv.Atoi(stringDistance)
		utils.Check(err)

		switch direction {
		case "L":
			pos -= distance
		case "R":
			pos += distance
		default:
			panic("Unrecognised distance")
		}

		pos = pos % 100

		if pos == 0 {
			count += 1
		}
	}
	return count
}

func CountClicks(input []string) int {
	pos := 50
	count := 0

	for _, line := range input {
		direction := string(line[0])
		stringDistance := string(line[1:])
		distance, err := strconv.Atoi(stringDistance)
		utils.Check(err)

		for range distance {
			switch direction {
			case "L":
				pos--
				if pos < 0 {
					pos = 99
				}
			case "R":
				pos++
				if pos > 99 {
					pos = 0
				}
			}

			if pos == 0 {
				count++
			}
		}
	}
	return count
}
