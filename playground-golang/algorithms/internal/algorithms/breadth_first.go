package algorithms

import datastruct "github.com/lennardclaproth/golang-playground/algorithms/internal/data-structures"

func BFS[T comparable](g *datastruct.Graph[T], source T) []T {
	// Make a map that keeps track of what vertexes
	// have already been visited
	visited := make(map[T]bool)
	// Tracks the traversal path through the breadth
	// first search
	result := []T{}

	// Initialize a queue to track what vertexes to
	// process next.
	q := datastruct.NewQueue[T]()
	// start with the source and set it to visited,
	// so that we never visit the same vertex multiple
	// times
	q.Enqueue(source)
	visited[source] = true
	// while the queue is not empty we need to keep
	// processing
	for !q.IsEmpty() {
		current := q.Dequeue()
		// Get the vertex of the current T and add
		// it to the result
		vertex := g.Vertices[current]
		result = append(result, vertex.Data)
		// Get the neighbours of the vertex, we do
		// this bt selecting the edges, each edge has
		// a vertex connected to it. We check if the
		// edge has already been visited, if it has not
		// been visited we add it to the queue and mark
		// it as visited.
		for neighborKey := range vertex.Edges {
			if !visited[neighborKey] {
				visited[neighborKey] = true
				q.Enqueue(neighborKey)
			}
		}
	}

	return result
}
