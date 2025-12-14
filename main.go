package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/12"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInput("12")

	fmt.Println(puzzles.CountRegions(input))
}
