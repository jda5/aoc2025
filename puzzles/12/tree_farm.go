package puzzles

import (
	"regexp"
	"strconv"
	"strings"
)

type Present struct {
	area int
}

type Region struct {
	width         int
	length        int
	presentCounts []int
}

// -------------------------------------------------------------------------------- helpers

func parseInput(input string) ([]Present, []Region) {
	parts := strings.Split(input, "\n\n")
	presents := make([]Present, 0)
	for _, shape := range parts[:len(parts)-1] {
		presents = append(presents, Present{area: strings.Count(shape, "#")})
	}

	dimensionsRegex := regexp.MustCompile(`(\d+)x(\d+)`)
	numbersRegex := regexp.MustCompile(` \d+`)
	regions := make([]Region, 0)

	for region := range strings.SplitSeq(parts[len(parts)-1], "\n") {

		dimensionMatches := dimensionsRegex.FindStringSubmatch(region)
		width, _ := strconv.Atoi(dimensionMatches[1])
		length, _ := strconv.Atoi(dimensionMatches[2])

		r := Region{width: width, length: length}

		counts := numbersRegex.FindAllString(region, -1)
		for _, count := range counts {
			c, _ := strconv.Atoi(strings.TrimSpace(count))
			r.presentCounts = append(r.presentCounts, c)
		}

		regions = append(regions, r)
	}

	return presents, regions
}

// -------------------------------------------------------------------------------- part one

func CountRegions(input string) int {
	res := 0
	presents, regions := parseInput(input)

	for _, region := range regions {
		regionArea := region.length * region.width
		presentArea := 0

		for i, count := range region.presentCounts {
			presentArea += presents[i].area * count
		}
		if regionArea >= presentArea {
			res += 1
		}
	}

	return res
}
