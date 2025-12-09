package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/09"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("09")

	fmt.Println(puzzles.CalculateLargestRectangle(input))
}
