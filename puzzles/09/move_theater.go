package puzzles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

type Coordinate [2]int

type Tiles [][]bool

// -------------------------------------------------------------------------------- helpers

func abs(x *int) int {
	if *x < 0 {
		return -(*x)
	}
	return *x
}

func formatCoordinates(input []string) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for _, row := range input {
		c := strings.Split(row, ",")

		x, err := strconv.Atoi(c[0])
		utils.Check(err)
		y, err := strconv.Atoi(c[1])
		utils.Check(err)

		coordinates = append(coordinates, Coordinate{x, y})
	}
	return coordinates
}

// useful for debugging
func (t *Tiles) print() {
	for _, row := range *t {
		fmtRow := make([]string, 0)
		for _, elem := range row {
			if !elem {
				fmtRow = append(fmtRow, ".")
			} else {
				fmtRow = append(fmtRow, "#")
			}
		}
		fmt.Println(fmtRow)
	}
}

// -------------------------------------------------------------------------------- puzzle one

func calculateArea(a, b *Coordinate) int {
	dy := (*b)[1] - (*a)[1]
	dx := (*b)[0] - (*a)[0]
	return (abs(&dy) + 1) * (abs(&dx) + 1)
}

func CalculateLargestRectangle(input []string) int {
	coordinates := formatCoordinates(input)
	maxArea := -1

	for i := range coordinates {
		a := &coordinates[i]
		for j := i + 1; j < len(coordinates); j++ {

			b := &coordinates[j]
			maxArea = max(calculateArea(a, b), maxArea)

		}
	}
	return maxArea
}

// -------------------------------------------------------------------------------- puzzle two

// determine the dimensions of the tile grid
func determineDimensions(coords []Coordinate) (int, int) {
	maxX, maxY := 0, 0
	for _, c := range coords {
		if c[0] > maxX {
			maxX = c[0]
		}
		if c[1] > maxY {
			maxY = c[1]
		}
	}
	return maxX + 2, maxY + 2
}

// draws the boundary between two coordinates
func drawBoundary(tiles *Tiles, curr *Coordinate, prev *Coordinate) {

	var start, end int

	if curr[1] == prev[1] {

		if curr[0] < prev[0] {
			start, end = curr[0]+1, prev[0]
		} else {
			start, end = prev[0]+1, curr[0]
		}
		for i := start; i < end; i++ {
			(*tiles)[curr[1]][i] = true
		}

	} else {
		if curr[1] < prev[1] {
			start, end = curr[1]+1, prev[1]
		} else {
			start, end = prev[1]+1, curr[1]
		}
		for i := start; i < end; i++ {
			(*tiles)[i][curr[0]] = true
		}
	}
}

// checks if the area between a and b is fully bounded by the boundary
func evaluateBoundedArea(tiles *Tiles, a *Coordinate, b *Coordinate) (int, error) {

	var directions [4][2]int

	if a[1] > b[1] {
		a, b = b, a
	}

	if a[0] < b[0] {
		directions = [4][2]int{
			{1, 0},  // right
			{0, 1},  // up
			{-1, 0}, // left
			{0, -1}, // down
		}
	} else {
		directions = [4][2]int{
			{-1, 0}, // left
			{0, 1},  // up
			{1, 0},  // right
			{0, -1}, // down
		}
	}

	maxY := max(a[1], b[1])
	minY := min(a[1], b[1])
	maxX := max(a[0], b[0])
	minX := min(a[0], b[0])

	curr := *a

	for _, dir := range directions {
		for {
			x := curr[0] + dir[0]
			y := curr[1] + dir[1]

			if x > maxX || x < minX {
				break
			}

			if y > maxY || y < minY {
				break
			}

			c := Coordinate{x, y}

			if c == *a {
				return calculateArea(a, b), nil
			}

			if !(*tiles)[y][x] {
				return -1, fmt.Errorf("invalid rectangle")
			}

			curr = c
		}
	}
	return -1, fmt.Errorf("invalid rectangle")

}

func CalculateLargestBoundedRectangle(input []string) int {

	coordinates := formatCoordinates(input)
	maxArea := -1

	maxX, maxY := determineDimensions(coordinates)
	tiles := make(Tiles, maxY)
	for i := range tiles {
		tiles[i] = make([]bool, maxX)
	}

	// draw the boundary made by the tiles

	prev := coordinates[0]
	tiles[prev[1]][prev[0]] = true

	for _, curr := range coordinates[1:] {
		tiles[curr[1]][curr[0]] = true
		drawBoundary(&tiles, &curr, &prev)
		prev = curr
	}
	drawBoundary(&tiles, &coordinates[0], &coordinates[len(coordinates)-1])

	// flood fill the inside of the boundary

	for i, row := range tiles {

		inBoundary := false
		prev := row[0]

		for j, curr := range row[1:] {
			if !curr && inBoundary {
				prev = curr
				tiles[i][j+1] = true
				continue
			}
			if !prev && curr && !inBoundary {
				inBoundary = true
			} else if inBoundary && curr && !prev {
				inBoundary = false
			}
			prev = curr
		}
	}

	// determine the largest bounded rectangle

	tilePtr := &tiles
	for i := range coordinates {
		a := &coordinates[i]
		for j := i + 1; j < len(coordinates); j++ {

			b := &coordinates[j]
			area, err := evaluateBoundedArea(tilePtr, a, b)
			if err == nil {
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea

}
