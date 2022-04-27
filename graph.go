package main

import "fmt"

type Graph map[Node][]Node

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
