package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Tag        string
	ID         string
	X, Y       int
	Visited    bool
	Neighbours []*Node
}

type Graph struct {
	Nodes map[string]*Node
}

var ants int
var nodes []Node
var edges []string

func FillGraph() (int, *Graph, string) {

	// "&x" is to get the address of the x -- > 0xc000192008
	// "*" is to get the content of the address, it undoes the &. "*&x" --> x

	if len(os.Args) < 2 {
		fmt.Println("Error: You need to give me a file please.")
		os.Exit(0)
	}
	contents := ReadFile(os.Args[1])
	GetNodes(contents)

	graph := MakeGraph()
	for nodeindex, node := range nodes {
		if nodeindex != 0 && (nodes[nodeindex-1]).X == node.X && (nodes[nodeindex-1]).Y == node.Y {
			fmt.Println("Error: Invalid data format, duplicate room location.")
			os.Exit(0)
		}
		graph.AddNode(node)
	}

	for _, edge := range edges {
		exit := strings.Split(edge, "-")[0]
		entrance := strings.Split(edge, "-")[1]
		graph.AddEdge(exit, entrance)
	}

	/* 	fmt.Println("GRAPH:", Graph)
	   	fmt.Println("GRAPHNODES:", Graph.Nodes)
	   	for _, node := range Graph.Nodes {
	   		fmt.Print("Tag:", node.Tag, " ID:", node.ID, " x,y:(", node.X, ",", node.Y, ") Neighbours: ")
	   		for _, neighbour := range node.Neighbours {
	   			fmt.Print(neighbour.ID, " ")
	   		}
	   		fmt.Println()
	   	} */

	return ants, graph, contents
}

func MakeGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(node Node) {
	if _, exists := g.Nodes[node.ID]; !exists {
		newnode := & /*  */ Node{
			Tag:        node.Tag,
			ID:         node.ID,
			X:          node.X,
			Y:          node.Y,
			Neighbours: []*Node{},
		}
		g.Nodes[node.ID] = newnode
	} else {
		fmt.Println("Error: Invalid data format, duplicate node ID.")
		os.Exit(0)
	}
}

func (g *Graph) AddEdge(nodeID1, nodeID2 string) {
	node1 := g.Nodes[nodeID1]
	node2 := g.Nodes[nodeID2]

	node1.Neighbours = append(node1.Neighbours, node2)
	node2.Neighbours = append(node2.Neighbours, node1)

}

func ReadFile(filename string) string {
	bytes, err := os.ReadFile("examples/" + filename)
	if err != nil {
		fmt.Println("Error: file does not exist.")
		os.Exit(0)
	}
	return string(bytes)
}

func GetNodes(contents string) {
	lines := strings.Split(contents, "\n")
	for lineindex, line := range lines {
		switch {
		case lineindex == 0:
			// its the number of ants
			ants = StringToInt(line)
			if ants == 0 {
				fmt.Println("Error: invalid data format, no ants.")
				os.Exit(0)
			}
		case strings.Contains(line, "-"):
			// its a tunnel
			edges = append(edges, line)
		case strings.Contains(lines[lineindex-1], "##start"):
			// its the start
			room := strings.Split(line, " ")
			var node Node
			node.Tag = "start"
			node.ID = room[0]
			node.X = StringToInt(room[1])
			node.Y = StringToInt(room[2])

			nodes = append(nodes, node)

		case strings.Contains(lines[lineindex-1], "##end"):
			// its the end
			room := strings.Split(line, " ")
			var node Node
			node.Tag = "end"
			node.ID = room[0]
			node.X = StringToInt(room[1])
			node.Y = StringToInt(room[2])

			nodes = append(nodes, node)

		case !(strings.Contains(line, "#")) && !(strings.Contains(line, "-")) && len(line) > 0 && !(strings.Contains(lines[lineindex-1], "##")):
			// its a room
			room := strings.Split(line, " ")
			var node Node
			node.Tag = "room"
			node.ID = room[0]
			node.X = StringToInt(room[1])
			node.Y = StringToInt(room[2])

			nodes = append(nodes, node)
		}
	}
	start := false
	end := false
	for _, node := range nodes {
		switch {
		case node.Tag == "start":
			start = true
		case node.Tag == "end":
			end = true
		}
	}
	if !start {
		fmt.Println("Error: invalid data format, no start room found.")
		os.Exit(0)
	}
	if !end {
		fmt.Println("Error: invalid data format, no end room found.")
		os.Exit(0)
	}
}

func StringToInt(str string) int {
	int, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return int
}
