package puzzles

import "strings"

// loc = rowIdx, columnIdx

// Set is a collection of unique elements
type Set map[[2]int]struct{}

func getAdjacentRolls(row int, col int, grid []string) [][2]int {

	positions := make([][2]int, 0)

	delta := [3]int{1, 0, -1}

	for _, i := range delta {
		for _, j := range delta {

			if i == 0 && j == 0 {
				continue
			}

			dy := row + i
			dx := col + j

			if dy < 0 || dy >= len(grid) {
				continue
			}

			if dx < 0 || dx >= len(grid[0]) {
				continue
			}

			val := string(grid[dy][dx])

			if val != "@" {
				continue
			}

			positions = append(positions, [2]int{dy, dx})
		}
	}

	return positions
}

// -------------------------------------------------------------------------------- part one

func CountRolls(grid []string) int {

	locations := make(Set, 0)

	for rowIdx, row := range grid {
		for colIdx, square := range row {

			if string(square) != "@" {
				continue
			}

			coors := getAdjacentRolls(rowIdx, colIdx, grid)

			if len(coors) >= 4 {
				continue
			}

			loc := [2]int{rowIdx, colIdx}

			locations[loc] = struct{}{}

		}
	}

	return len(locations)
}

// -------------------------------------------------------------------------------- part two

func RemoveRolls(input []string) int {

	removed := 0
	grid := input

	for {

		locations := make(Set, 0)
		removedRolls := false

		for rowIdx, row := range grid {
			for colIdx, square := range row {

				if string(square) != "@" {
					continue
				}

				coors := getAdjacentRolls(rowIdx, colIdx, grid)

				if len(coors) >= 4 {
					continue
				}

				loc := [2]int{rowIdx, colIdx}

				locations[loc] = struct{}{}
			}
		}

		// create a new grid with the rolls remove
		newGrid := make([]string, 0)
		for rowIdx, row := range grid {

			var sb strings.Builder

			for colIdx, square := range row {

				loc := [2]int{rowIdx, colIdx}

				_, ok := locations[loc]
				if ok {
					sb.WriteString(".")
					removed++
					removedRolls = true
				} else {
					sb.WriteString(string(square))
				}

			}
			newGrid = append(newGrid, sb.String())
		}

		if !removedRolls {
			break
		}

		grid = newGrid

	}

	return removed
}
