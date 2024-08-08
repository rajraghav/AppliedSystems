package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBfsShortestPath(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}

	// Test existing path
	path := bfsShortestPath(graph, "A", "F")
	expectedPath := []string{"A", "C", "F"}
	assert.Equal(t, expectedPath, path)

	// Test no path
	path = bfsShortestPath(graph, "A", "Z")
	assert.Nil(t, path)

	// Test same start and end
	path = bfsShortestPath(graph, "A", "A")
	expectedPath = []string{"A"}
	assert.Equal(t, expectedPath, path)
}
