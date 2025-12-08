package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/08"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("08")

	fmt.Println(puzzles.CalculateCircuits(input, 1000))
}
