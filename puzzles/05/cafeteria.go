package puzzles

import (
	"cmp"
	"slices"
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

func parseInput(db []string) ([][2]int, []int) {
	ranges := make([][2]int, 0)
	ids := make([]int, 0)

	sepIdx := 0
	for rowIdx, row := range db {
		if row == "" {
			sepIdx = rowIdx
			break
		}
		items := strings.Split(row, "-")
		var formattedRow [2]int

		for i, item := range items {
			itemInt, err := strconv.Atoi(item)
			utils.Check(err)
			formattedRow[i] = itemInt
		}
		ranges = append(ranges, formattedRow)
	}

	for _, row := range db[sepIdx+1:] {
		id, err := strconv.Atoi(row)
		utils.Check(err)
		ids = append(ids, id)
	}
	return ranges, ids
}

func isFresh(ranges [][2]int, id int) bool {
	for _, ids := range ranges {
		if id >= ids[0] && id <= ids[1] {
			return true
		}
	}
	return false
}

// -------------------------------------------------------------------------------- part one

func CountFresh(db []string) int {
	ranges, ids := parseInput(db)
	count := 0

	for _, id := range ids {
		if isFresh(ranges, id) {
			count++
		}
	}

	return count
}

// -------------------------------------------------------------------------------- part two

func CountTotalFreshIds(db []string) int {
	ranges, _ := parseInput(db)
	count := 0

	// sort by first range value
	slices.SortFunc(ranges, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })

	// remove all overlapping
	uniqueRanges := make([][2]int, 0)
	prev := ranges[0]

	for _, curr := range ranges[1:] {
		if curr[0] > prev[1] {
			uniqueRanges = append(uniqueRanges, prev)
			prev = curr
			continue
		}
		if curr[1] > prev[1] {
			prev[1] = curr[1]
		}
	}
	uniqueRanges = append(uniqueRanges, prev)

	for _, row := range uniqueRanges {
		count += row[1] - row[0] + 1
	}

	return count
}
