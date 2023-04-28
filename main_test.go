package main

import "testing"

func TestAStar(t *testing.T) {
	graph := NewGraph()

	// Add cities to the graph
	graph.AddCity("A", 52.2297, 21.0122)
	graph.AddCity("B", 51.5074, -0.1278)
	graph.AddCity("C", 48.8566, 2.3522)
	graph.AddCity("D", 40.7128, -74.0060)

	// Add edges to the graph
	graph.AddEdge("A", "B", 1448.79)
	graph.AddEdge("A", "C", 1053.81)
	graph.AddEdge("B", "C", 344.35)
	graph.AddEdge("B", "D", 5573.07)
	graph.AddEdge("C", "D", 5844.25)

	// Test case 1: start and goal nodes are the same
	path, cost := graph.AStar("A", "A")
	if len(path) != 1 || path[0] != "A" || cost != 0 {
		t.Errorf("Test case 1 failed: expected [A], 0, but got %v, %f", path, cost)
	}

	// Test case 2: start and goal nodes are adjacent
	path, cost = graph.AStar("A", "B")
	if len(path) != 2 || path[0] != "A" || path[1] != "B" || cost != 1448.79 {
		t.Errorf("Test case 2 failed: expected [A B], 1448.79, but got %v, %f", path, cost)
	}

	// Test case 3: start and goal nodes are adjacent
	path, cost = graph.AStar("A", "C")
	if len(path) != 3 || path[0] != "A" || path[1] != "B" || path[2] != "C" || cost != 1053.81 {
		t.Errorf("Test case 3 failed: expected [A B C], 1053.81, but got %v, %f", path, cost)
	}

	// Test case 4: start and goal nodes are not connected
	path, cost = graph.AStar("A", "D")
	if len(path) != 4 || path[0] != "A" || path[1] != "C" || path[2] != "D" || cost != 6898.06 {
		t.Errorf("Test case 3 failed: expected [A C D], 6898.06, but got %v, %f", path, cost)
	}
}
