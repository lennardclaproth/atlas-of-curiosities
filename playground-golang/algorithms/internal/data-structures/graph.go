package datastruct

import "fmt"

// An Undirected graph that contains vertices.
type Graph[T comparable] struct {
	Vertices map[T]*Vertex[T]
}

// A vertex that maintains it's own edges and data, this is
// a node in the graph
type Vertex[T comparable] struct {
	Data  T
	Edges map[T]*Edge[T]
}

// The relationships to other vertices in the graph.
type Edge[T comparable] struct {
	Vertex *Vertex[T]
}

// NewGraph creates a new graph and returns the created graph.
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		Vertices: make(map[T]*Vertex[T]),
	}
}

// AddVertex instantiates a new vertex with the data as a key
// and checks if the verted already exists.
func (g *Graph[T]) AddVertex(d T) error {
	// Throw error when vertex already exists
	if g.Vertices[d] != nil {
		return fmt.Errorf("vertex already exists in graph")
	}
	// Instantiate the new vertex
	g.Vertices[d] = &Vertex[T]{Edges: make(map[T]*Edge[T])}
	return nil
}

// AddEdge adds a new edge to both a and b, this is because it
// is an undirected graph therefore both should point to eachother.
func (g *Graph[T]) AddEdge(da, db T) error {
	// Get the vertices from the graph
	va := g.Vertices[da]
	vb := g.Vertices[db]
	// If one of the vertices is nil we cannot create an edge because there
	// there can be no relationship
	if va == nil || vb == nil {
		return fmt.Errorf("one or more vertices do not exist, cannot create an edge on a non existing vertex")
	}
	// We set the edge with the comparable data of b to a because
	// a points to b and vice versa
	va.Edges[db] = &Edge[T]{Vertex: vb}
	vb.Edges[da] = &Edge[T]{Vertex: va}
	return nil
}

// AdjacencyMatrix generates an adjacency matrix based
// off the graph. An adjacency matrix has a time complexity of
// O(n) and a space complexity of O(v^2) where v is the amount
// of vertices.
func (g *Graph[T]) AdjacencyMatrix() ([][]bool, []T) {
	indexMap := make(map[T]int)
	keys := make([]T, 0, len(g.Vertices))
	i := 0
	for k := range g.Vertices {
		indexMap[k] = i
		keys = append(keys, k)
		i++
	}

	n := len(g.Vertices)
	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}

	for aKey, aVertex := range g.Vertices {
		for bKey := range aVertex.Edges {
			i := indexMap[aKey]
			j := indexMap[bKey]
			matrix[i][j] = true
			matrix[j][i] = true
		}
	}

	return matrix, keys
}
