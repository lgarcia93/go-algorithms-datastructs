package main

import "fmt"

func main() {

	//Create the nodes
	NodeA := Node{
		Key: "a",
	}

	NodeB := Node{
		Key: "b",
	}

	NodeC := Node{
		Key: "c",
	}

	NodeD := Node{
		Key: "d",
	}

	NodeE := Node{
		Key: "e",
	}

	NodeF := Node{
		Key: "f",
	}
	graph := Graph{
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
