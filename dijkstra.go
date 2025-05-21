package main

import (
	"math"
)

func Dijkstra(g *Graph, start, end string) ([]string, int) {
	dist := make(map[string]int)
	prev := make(map[string]string)
	unvisited := make(map[string]bool)

	for node := range g.Nodes {
		dist[node] = math.MaxInt32
		unvisited[node] = true
	}

	dist[start] = 0

	for len(unvisited) > 0 {
		var current string
		minDist := math.MaxInt32
		for node := range unvisited {
			if dist[node] < minDist {
				current = node
				minDist = dist[node]
			}
		}

		if current == end {
			break
		}

		for neighbor, weight := range g.Nodes[current] {
			alt := dist[current] + weight
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = current
			}
		}

		delete(unvisited, current)
	}

	// Reconstrução do caminho
	path := []string{}
	for u := end; u != ""; u = prev[u] {
		path = append([]string{u}, path...)
		if u == start {
			break
		}
	}

	if len(path) == 0 || path[0] != start {
		return nil, 0
	}

	return path, dist[end]
}
