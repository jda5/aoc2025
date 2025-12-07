package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/07"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("07")

	fmt.Println(puzzles.CountTimelines(input))
}
