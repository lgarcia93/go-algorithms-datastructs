package graphs

import "fmt"

type Graph map[Node][]Node

var visitedNodes = map[Node]bool{}

func (g *Graph) HasPath(src Node, dest Node) bool {
	if src.Key == dest.Key {
		return true
	}

	for _, v := range (*g)[src] {
		if g.HasPath(v, dest) {
			return true
		}
	}

	return false
}

func (g *Graph) HasPathUndirected(src Node, dest Node) bool {
	if src.Key == dest.Key {
		return true
	}

	for _, v := range (*g)[src] {

		wasVisited, _ := visitedNodes[v]

		if wasVisited {
			continue
		}

		if g.HasPathUndirected(v, dest) {
			return true
		}

		visitedNodes[v] = true
	}

	return false
}

func (g *Graph) NodeExists(node Node) bool {
	_, ok := (*g)[node]

	return ok
}

func (g *Graph) DepthFirstSearch(sourceNode Node) {
	if !g.NodeExists(sourceNode) {
		return
	}

	var stack Stack[Node]
	stack.Push(sourceNode)

	for !stack.IsEmpty() {

		if current, ok := stack.Pop(); ok {

			fmt.Print(current.Key)

			for _, v := range (*g)[current] {
				stack.Push(v)
			}
		}
	}
}

func (g *Graph) BreadthFirstSearch(sourceNode Node) {
	if !g.NodeExists(sourceNode) {
		return
	}

	var queue Queue[Node]

	queue.Add(sourceNode)

	for !queue.IsEmpty() {
		if current, ok := queue.Deque(); ok {
			fmt.Println(current.Key)

			for _, v := range (*g)[current] {
				queue.Add(v)
			}
		}
	}
}

func BuildUndirectedGraphFromEdges(edges []Edge) Graph {
	graph := Graph{}

	for _, edge := range edges {
		graph[edge.NodeA] = append(graph[edge.NodeA], edge.NodeB)
		graph[edge.NodeB] = append(graph[edge.NodeB], edge.NodeA)
	}

	return graph
}

func (g *Graph) ConnectedCountAndBiggestComponent() (componentCount int, biggestComponentSize int) {
	visitedNodes = make(map[Node]bool)

	componentCount = 0

	biggestComponentSize = -1

	for k := range *g {

		if wasVisited, _ := visitedNodes[k]; wasVisited {
			continue
		}

		componentCount++

		stack := Stack[Node]{}

		stack.Push(k)

		size := 1

		visitedNodes[k] = true

		for !stack.IsEmpty() {
			current, _ := stack.Pop()

			for _, v := range (*g)[current] {
				if wasVisited, _ := visitedNodes[v]; wasVisited {
					continue
				}

				stack.Push(v)
				size++

				visitedNodes[v] = true
			}
		}

		if biggestComponentSize < size {
			biggestComponentSize = size
		}
	}

	return componentCount, biggestComponentSize
}
