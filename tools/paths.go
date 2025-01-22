package tools

import (
	"fmt"
	"os"
)

func (g *Graph) BreadthFirstSearch() [][]string {
	var paths [][]string // To store completed paths only.
	var queue [][]*Node  // To keep track of current path.
	var startnode, endnode *Node

	// Find the start and end nodes.
	for _, node := range g.Nodes {
		switch {
		case node.Tag == "start":
			startnode = node
		case node.Tag == "end":
			endnode = node
		}
	}
	// Give queue start node.
	queue = append(queue, []*Node{startnode})

	// len(queue) will be 0 when at end node
	for len(queue) > 0 {
		// Get first path from queue and remove path from queue.
		// fmt.Println(queue)
		path := queue[0]
		queue = queue[1:]
		// fmt.Println(queue)

		// Get the last node from path.
		prevnode := path[len(path)-1]

		// If it is end, then create the completed path.
		if prevnode == endnode {
			var completepath []string
			for _, node := range path {
				completepath = append(completepath, node.ID)
			}
			paths = append(paths, completepath)
		}

		// Check the last node's neighbours, check if they are already in the path,
		// Or if they are already visited in a different path.
		for _, neighbour := range prevnode.Neighbours {
			if !ContainsNode(path, neighbour) {
				// fmt.Println(prevnode.ID, neighbour.ID, neighbour.Visited)
				// If unvisited, add path with additional node to queue.
				newpath := make([]*Node, len(path))
				copy(newpath, path)
				newpath = append(newpath, neighbour)
				queue = append(queue, newpath)
			}
		}
	}
	if len(paths) == 0 {
		fmt.Println("Error: invalid data format, no valid paths.")
		os.Exit(0)
	}
	return paths
}

func ContainsNode(path []*Node, node *Node) bool {
	for _, p := range path {
		if p == node {
			return true
		}
	}
	return false
}
