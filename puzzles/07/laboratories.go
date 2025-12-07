package puzzles

import (
	"fmt"

	"github.com/jda5/aoc2025/utils"
)

// -------------------------------------------------------------------------------- coordinate data structure

type Coordinate [2]int

func (c *Coordinate) move() {
	(*c)[0] += 1
}

// -------------------------------------------------------------------------------- queue data structure

type Queue []Coordinate

// Enqueue adds an element to the rear of the queue
func (q *Queue) Enqueue(value Coordinate) {
	*q = append(*q, value)
}

// Dequeue removes and returns an element from the front of the queue
func (q *Queue) Dequeue() (Coordinate, error) {
	if q.IsEmpty() {
		return Coordinate{}, fmt.Errorf("empty queue")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// -------------------------------------------------------------------------------- diagram data structure

type Diagram [][]string

func (d *Diagram) get(c *Coordinate) (string, error) {
	if (*c)[0] >= 0 && (*c)[0] < len(*d) &&
		(*c)[1] >= 0 && (*c)[1] < len((*d)[0]) {
		return (*d)[(*c)[0]][(*c)[1]], nil
	}
	return "", fmt.Errorf("Coordinate is not on the diagram")
}

func (d *Diagram) update(c *Coordinate) {
	(*d)[(*c)[0]][(*c)[1]] = "|"
}

// print outputs the diagram to the console - for debugging purposes
func (d *Diagram) print() {
	for _, row := range *d {
		fmt.Println(row)
	}
}

func createDiagram(input []string) *Diagram {

	d := make(Diagram, 0)

	for _, row := range input {
		diagramRow := make([]string, 0)
		for _, val := range row {
			diagramRow = append(diagramRow, string(val))
		}
		d = append(d, diagramRow)
	}

	return &d
}

func locateStart(d *Diagram) Coordinate {
	for idx, val := range (*d)[0] {
		if string(val) == "S" {
			return Coordinate{0, idx}
		}
	}
	panic("Unable to find starting point")
}

// -------------------------------------------------------------------------------- puzzles

func CountSplits(input []string) int {

	diagram := createDiagram(input)

	splitCount := 0

	q := make(Queue, 0)
	q.Enqueue(locateStart(diagram))

	for !q.IsEmpty() {
		beam, err := q.Dequeue()
		utils.Check(err)

		for {
			beam.move()
			tile, err := diagram.get(&beam)

			if err != nil {
				// the beam is no longer on the diagram
				break
			}

			if string(tile) == "|" {
				// the beam has hit an existing beam path
				break
			}

			if string(tile) == "^" {

				newBeams := [2]Coordinate{
					{beam[0], beam[1] - 1},
					{beam[0], beam[1] + 1},
				}

				for _, new := range newBeams {

					newPtr := &new

					startTile, err := diagram.get(newPtr)
					if err != nil {
						// the new beam is not on the diagram
						continue
					}

					if startTile == "." {
						q.Enqueue(new)
						diagram.update(newPtr)
					}
				}

				splitCount++

				// the old beam ceases to exist
				break

			} else {
				// update the diagram with the beam path
				diagram.update(&beam)
			}

		}

	}

	return splitCount
}

// func CountTimelines(input []string) int {

// 	diagram := createDiagram(input)

// 	timelineCount := 0

// 	q := make(Queue, 0)
// 	q.Enqueue(locateStart(diagram))

// 	for !q.IsEmpty() {
// 		beam, err := q.Dequeue()
// 		utils.Check(err)

// 		for {
// 			beam.move()
// 			tile, err := diagram.get(&beam)

// 			if err != nil {
// 				// the beam is no longer on the diagram
// 				timelineCount++
// 				break
// 			}

// 			if string(tile) == "^" {

// 				newBeams := [2]Coordinate{
// 					{beam[0], beam[1] - 1},
// 					{beam[0], beam[1] + 1},
// 				}

// 				for _, start := range newBeams {

// 					startPtr := &start

// 					startTile, err := diagram.get(startPtr)
// 					if err != nil {
// 						// the new beam is not on the diagram
// 						continue
// 					}

// 					if startTile == "." {
// 						q.Enqueue(start)
// 					}
// 				}

// 				// the old beam ceases to exist
// 				break
// 			}

// 		}

// 	}

// 	return timelineCount
// }

func CountTimelines(input []string) int {

	diagram := createDiagram(input)
	origins := make(map[Coordinate]int, 0)

	q := make(Queue, 0)

	for rowIdx := len(*diagram) - 1; rowIdx >= 0; rowIdx-- {
		for colIdx, val := range (*diagram)[rowIdx] {
			if val == "^" {
				for _, coor := range [2]Coordinate{{rowIdx, colIdx + 1}, {rowIdx, colIdx - 1}} {
					if _, ok := origins[coor]; !ok {
						origins[coor] = 0
						q.Enqueue(coor)
					}
				}
			}
		}
	}

	for !q.IsEmpty() {
		beam, err := q.Dequeue()
		utils.Check(err)

		start := beam
		timelines := 0

		for {
			beam.move()
			tile, err := diagram.get(&beam)
			if err != nil {
				// the beam is no longer on the diagram
				origins[start] = timelines + 1
				break
			}

			if string(tile) == "^" {
				newBeams := [2]Coordinate{
					{beam[0], beam[1] - 1},
					{beam[0], beam[1] + 1},
				}
				for _, new := range newBeams {
					timelines += origins[new]
				}
				origins[start] = timelines
				break
			}
		}
	}

	beam := locateStart(diagram)
	totalTimelines := 0

	for {
		beam.move()
		tile, _ := diagram.get(&beam)
		if string(tile) == "^" {
			newBeams := [2]Coordinate{
				{beam[0], beam[1] - 1},
				{beam[0], beam[1] + 1},
			}
			for _, new := range newBeams {
				totalTimelines += origins[new]
			}
			return totalTimelines
		}
	}
}
