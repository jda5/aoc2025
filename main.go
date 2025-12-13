package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/11"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("11")

	fmt.Println(puzzles.CountPathsThroughNodes(input))
}
