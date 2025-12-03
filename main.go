package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/03"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("03")

	fmt.Println(puzzles.TotalJoltage(input))
}
