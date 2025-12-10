package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/10"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("10")

	fmt.Println(puzzles.FewestButtonPresses(input))
}
