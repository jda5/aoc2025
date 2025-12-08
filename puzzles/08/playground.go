package puzzles

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

// -------------------------------------------------------------------------------- data structures and helpers

type Box struct {
	X int
	Y int
	Z int
}

type BoxDistance struct {
	Boxes    [2]*Box
	Distance float64
}

type Circuit map[*Box]struct{}

func distance(a, b *Box) float64 {
	dx := float64(b.X - a.X)
	dy := float64(b.Y - a.Y)
	dz := float64(b.Z - a.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func createBox(data string) *Box {
	coordinates := strings.Split(data, ",")

	x, err := strconv.Atoi(coordinates[0])
	utils.Check(err)

	y, err := strconv.Atoi(coordinates[1])
	utils.Check(err)

	z, err := strconv.Atoi(coordinates[2])
	utils.Check(err)

	return &Box{
		X: x,
		Y: y,
		Z: z,
	}
}

// -------------------------------------------------------------------------------- puzzle one

func CalculateCircuits(input []string, numConnections int) int {
	res := 1

	boxes := make([]*Box, 0)
	circuits := make(map[*Box]*Circuit, 0)
	for _, data := range input {
		box := createBox(data)
		boxes = append(boxes, box)

		circuit := make(Circuit, 0)
		circuit[box] = struct{}{}
		circuits[box] = &circuit
	}

	distances := make([]BoxDistance, 0)
	for i := 0; i < len(boxes); i++ {
		b1 := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			b2 := boxes[j]
			distances = append(distances, BoxDistance{
				Boxes:    [2]*Box{b1, b2},
				Distance: distance(b1, b2),
			})
		}
	}

	slices.SortFunc(distances, func(a, b BoxDistance) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	for i := range numConnections {
		distance := distances[i]
		b1 := distance.Boxes[0]
		b2 := distance.Boxes[1]

		c1 := circuits[b1]
		c2 := circuits[b2]

		if c1 == c2 {
			continue
		}

		// Always merge the smaller circuit into the larger one to do less work.
		if len(*c1) < len(*c2) {
			c1, c2 = c2, c1
		}

		// Merge c2 into c1
		for box := range *c2 {
			(*c1)[box] = struct{}{}
			circuits[box] = c1
		}

	}

	circuitMap := make(map[*Circuit]struct{}, 0)
	cirtcuitArray := make([]*Circuit, 0)
	for _, val := range circuits {
		if _, exists := circuitMap[val]; !exists {
			circuitMap[val] = struct{}{}
			cirtcuitArray = append(cirtcuitArray, val)
		}
	}

	slices.SortFunc(cirtcuitArray, func(a, b *Circuit) int {
		return cmp.Compare(len(*b), len(*a))
	})

	for i := range 3 {
		res *= len(*cirtcuitArray[i])
	}

	return res
}
