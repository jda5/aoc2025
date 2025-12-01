package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/01"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	fmt.Println(puzzles.CountClicks(utils.ReadInputLines("01")))
}
