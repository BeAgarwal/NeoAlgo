// Depth First Search (Graph) in Golang

package main

import (
	"fmt"
	"strconv"
)

//creating a node
type node struct {
	v    *Vertex
	next *node
}

func DFS(g *Graph, root *Vertex, visit func(int)) {
	
	//maintaining the map of all visited node
	visited := map[int]bool{}

	if root == nil {
		return
	}

	visited[root.info] = true
	visit(root.info)

	//calling the same DFS for all its adjacent
	for _, v := range root.Vertices {
		if visited[v.info] {
			continue
		}
		DFS(g, v, visit)
	}
}

type Vertex struct {
	// the data
	info int
	//vertices connecting to other vertex
	Vertices map[int]*Vertex
}

//creating a new vertex
func NewVertex(info int) *Vertex {
	return &Vertex{
		info:      info,
		Vertices: map[int]*Vertex{},
	}
}

func (v *Vertex) String() string {
	s := strconv.Itoa(v.info) + ":"

	for _, next := range v.Vertices {
		s += " " + strconv.Itoa(next.info)
	}

	return s
}

type Graph struct {
	//map for all vertices connected to each other
	Vertices map[int]*Vertex
	// This will decide if it's a directed or undirected graph
	directed bool
}

//creating a new graph
func NewGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

//function for adding vertex
func (g *Graph) AddVertex(info int) {
	v := NewVertex(info)
	g.Vertices[info] = v
}

// function for adding edge
func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		return
	}

	if _, ok := v1.Vertices[v2.info]; ok {
		return
	}

	//adding a directed edge
	v1.Vertices[v2.info] = v2
	if !g.directed && v1.info != v2.info {
		v2.Vertices[v1.info] = v1
	}

	// adding the vertices
	g.Vertices[v1.info] = v1
	g.Vertices[v2.info] = v2
}

func (g *Graph) String() string {
	st := ""
	i := 0
	for _, vertex := range g.Vertices {
		if i != 0 {
			st += "\n"
		}
		st += vertex.String()
		i++
	}
	return st
}

func main() {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)

	g.AddEdge(1, 9)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)

	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	DFS(g, g.Vertices[1], cb)

	fmt.Print("Depth First Search: ")
	fmt.Println(visitedOrder)
}

//OUTPUT :

// Depth First Search: [1 9 10 5 6 7 8 2]
