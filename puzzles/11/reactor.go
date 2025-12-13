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

func (g *Graph) AddNode(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

// -------------------------------------------------------------------------------- helper

func createGraph(input []string) *Graph {
	g := MakeGraph()
	for _, row := range input {
		parts := strings.Split(row, ": ")
		from := parts[0]
		for to := range strings.SplitSeq(parts[1], " ") {
			g.AddNode(from, to)
		}
	}
	return g
}

// -------------------------------------------------------------------------------- puzzle one

var Cache map[string]int

func (g *Graph) CachedDFS(curr string, end *string) int {
	if count, exists := Cache[curr]; exists {
		return count
	}
	if curr == *end {
		Cache[curr] = 1
		return 1
	}
	count := 0
	for _, node := range g.vertices[curr] {
		count += g.CachedDFS(node, end)
	}
	Cache[curr] = count
	return count
}

func CountPaths(input []string) int {
	g := createGraph(input)

	Cache = make(map[string]int)

	end := "out"
	return g.CachedDFS("you", &end)
}

// -------------------------------------------------------------------------------- puzzle two

func (g *Graph) CachedDFSThroughNodes(curr string, end *string, fft bool, dac bool) int {

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
			count += g.CachedDFSThroughNodes(node, end, true, dac)
		case "dac":
			count += g.CachedDFSThroughNodes(node, end, fft, true)
		default:
			count += g.CachedDFSThroughNodes(node, end, fft, dac)
		}
	}
	Cache[cacheIdx] = count
	return count
}

func CountPathsThroughNodes(input []string) int {
	g := createGraph(input)
	Cache = make(map[string]int)
	end := "out"
	return g.CachedDFSThroughNodes("svr", &end, false, false)
}
