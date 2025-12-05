package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/05"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("05")

	fmt.Println(puzzles.CountTotalFreshIds(input))
}
