package main

import (
	"container/heap"
	"fmt"
	"math"
)

type City struct {
	name     string
	lat, lng float64
}

type Edge struct {
	to     *City
	weight float64
}

type Graph struct {
	cities map[string]*City
	edges  map[string][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		cities: make(map[string]*City),
		edges:  make(map[string][]*Edge),
	}
}

func (g *Graph) AddCity(name string, lat, lng float64) {
	city := &City{name, lat, lng}
	g.cities[name] = city
	g.edges[name] = []*Edge{}
}

func (g *Graph) AddEdge(from, to string, weight float64) {
	edge := &Edge{g.cities[to], weight}
	g.edges[from] = append(g.edges[from], edge)
}

func haversineDistance(city1, city2 *City) float64 {
	const R = 6371 // Earth's radius in km
	dLat := (city2.lat - city1.lat) * math.Pi / 180
	dLng := (city2.lng - city1.lng) * math.Pi / 180

	lat1 := city1.lat * math.Pi / 180
	lat2 := city2.lat * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

type PQItem struct {
	city      *City
	cost      float64
	heuristic float64
	index     int
}

type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost+pq[i].heuristic < pq[j].cost+pq[j].heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (g *Graph) AStar(start, goal string) (path []string, cost float64) {
	startCity := g.cities[start]
	goalCity := g.cities[goal]

	costs := make(map[string]float64)
	costs[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)

	item := &PQItem{
		city:      startCity,
		cost:      0,
		heuristic: haversineDistance(startCity, goalCity),
	}
	heap.Push(pq, item)

	cameFrom := make(map[string]string)

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*PQItem)
		currentCity := currentItem.city

		if currentCity.name == goal {
			path = []string{goal}
			for currentCity.name != start {
				path = append([]string{cameFrom[currentCity.name]}, path...)
				currentCity = g.cities[cameFrom[currentCity.name]]
			}
			return path, costs[goal]
		}

		for _, edge := range g.edges[currentCity.name] {
			newCost := costs[currentCity.name] + edge.weight
			if _, ok := costs[edge.to.name]; !ok || newCost < costs[edge.to.name] {
				costs[edge.to.name] = newCost
				cameFrom[edge.to.name] = currentCity.name

				item := &PQItem{
					city:      edge.to,
					cost:      newCost,
					heuristic: haversineDistance(edge.to, goalCity),
				}
				heap.Push(pq, item)
			}
		}
	}

	return nil, math.Inf(1)
}

func main() {
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

	path, cost := graph.AStar("A", "D")
	fmt.Printf("Path: %v, Cost: %.2f\n", path, cost)
}
