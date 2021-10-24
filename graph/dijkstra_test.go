package graph

import (
	"testing"
)

var testCases = []struct {
	name     string
	edges    [][]int
	N        int
	node0    int
	node1    int
	expected int
}{
	{
		"straight line graph",
		[][]int{{0, 1, 5}, {1, 2, 2}},
		3, 0, 2, 7,
	},
	{
		"unconnected node",
		[][]int{{0, 1, 5}},
		3, 0, 2, -1,
	},
	{
		"double paths",
		[][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}},
		4, 0, 3, 9,
	},
	{
		"double paths extended",
		[][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}, {3, 4, 1}},
		5, 0, 4, 10,
	},
}

func TestDijkstra(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Dijkstra(tc.N, tc.edges, tc.node0, tc.node1)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d, from node %d to %d, with %v",
					tc.expected, actual, tc.node0, tc.node1, tc.edges)
			}
		})
	}
}
