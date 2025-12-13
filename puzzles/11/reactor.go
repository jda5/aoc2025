package puzzles

import (
	"fmt"
	"strings"
)

// -------------------------------------------------------------------------------- graph data structure

type Graph struct {
	vertices map[string][]string
}

func MakeGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

func (g *Graph) AddEdge(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

// -------------------------------------------------------------------------------- helpers

func createDeviceGraph(input []string) *Graph {
	g := MakeGraph()
	for _, row := range input {
		parts := strings.Split(row, ": ")
		from := parts[0]
		for to := range strings.SplitSeq(parts[1], " ") {
			g.AddEdge(from, to)
		}
	}
	return g
}

// -------------------------------------------------------------------------------- puzzle one

func (g *Graph) DFS(curr string, end *string) int {
	if curr == *end {
		return 1
	}
	count := 0
	for _, node := range g.vertices[curr] {
		count += g.DFS(node, end)
	}
	return count
}

func GetPathCount(input []string) int {
	g := createDeviceGraph(input)

	start := "you"
	end := "out"

	return g.DFS(start, &end)
}

// -------------------------------------------------------------------------------- puzzle two

var Cache map[string]int = make(map[string]int)

func (g *Graph) CachedDFS(curr string, end *string, fft bool, dac bool) int {

	cacheIdx := fmt.Sprintf("%v-%v-%v", curr, fft, dac)

	if count, exists := Cache[cacheIdx]; exists {
		return count
	}

	if curr == *end {
		if fft && dac {
			Cache[cacheIdx] = 1
			return 1
		}
		Cache[cacheIdx] = 0
		return 0
	}

	count := 0
	for _, node := range g.vertices[curr] {
		switch node {
		case "fft":
			count += g.CachedDFS(node, end, true, dac)
		case "dac":
			count += g.CachedDFS(node, end, fft, true)
		default:
			count += g.CachedDFS(node, end, fft, dac)
		}

	}
	Cache[cacheIdx] = count
	return count
}

func GetPathCountThroughNodes(input []string) int {
	g := createDeviceGraph(input)

	start := "svr"
	end := "out"

	res := g.CachedDFS(start, &end, false, false)
	return res
}
