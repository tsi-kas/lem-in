package main

import (
	"fmt"
	"lem-in/tools"
)

func main() {

	ants, graph, problem := tools.FillGraph()
	// fmt.Println("ants:", ants)
	// fmt.Println("graph:", graph)

	paths := graph.BreadthFirstSearch()
	// fmt.Println("paths:", paths)

	flowmap := tools.Flow(paths)
	// fmt.Println("flowmap", flowmap)

	optimalpath, steps := tools.Optimal(paths, flowmap, ants)
	// fmt.Println("optimalpath", optimalpath)

	solution := tools.AntAssignToPath(optimalpath, ants, steps)

	fmt.Println(problem)
	fmt.Println(solution)

}
