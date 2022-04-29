package main

import (
	"fmt"
	"go-algorithms-datastructs/graphs"
)

func main() {

	//Create the nodes
	NodeA := graphs.Node{
		Key: "a",
	}

	NodeB := graphs.Node{
		Key: "b",
	}

	NodeC := graphs.Node{
		Key: "c",
	}

	NodeD := graphs.Node{
		Key: "d",
	}

	NodeE := graphs.Node{
		Key: "e",
	}

	NodeF := graphs.Node{
		Key: "f",
	}
	graph := graphs.Graph{
		NodeA: {
			NodeB,
			NodeC,
		},
		NodeB: {
			NodeD,
		},
		NodeC: {
			NodeE,
		},
		NodeD: {
			NodeF,
		},
		NodeE: {},
		NodeF: {},
	}

	fmt.Println("Printing Depth First Search...")
	graph.DepthFirstSearch(NodeA)

	fmt.Println("-------------------")

	fmt.Println("Printing Breadth First Search...")
	graph.BreadthFirstSearch(NodeA)

	fmt.Println("-------------------")

}
