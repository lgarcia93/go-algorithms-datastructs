package test

import (
	"go-algorithms-datastructs/graphs"
	"reflect"
	"testing"
)

func TestHasPath(t *testing.T) {

	NodeA := graphs.Node{Key: "a"}
	NodeB := graphs.Node{Key: "b"}
	NodeC := graphs.Node{Key: "c"}
	NodeD := graphs.Node{Key: "d"}
	NodeE := graphs.Node{Key: "e"}
	NodeF := graphs.Node{Key: "f"}

	graph := graphs.Graph{
		NodeA: {NodeB, NodeC},
		NodeB: {NodeD},
		NodeC: {NodeE},
		NodeD: {NodeF},
		NodeE: {},
		NodeF: {},
	}

	actual := graph.HasPath(NodeA, NodeF)
	expect := true

	if actual != expect {
		t.Fail()
	}

	actual = graph.HasPath(NodeF, NodeA)
	expect = false

	if actual != expect {
		t.Fail()
	}
}

func TestBuildUndirectedGraphFromEdges(t *testing.T) {

	nodeI := graphs.Node{Key: "i"}
	nodeJ := graphs.Node{Key: "j"}
	nodeK := graphs.Node{Key: "k"}
	nodeM := graphs.Node{Key: "m"}
	nodeL := graphs.Node{Key: "l"}
	nodeO := graphs.Node{Key: "o"}
	nodeN := graphs.Node{Key: "n"}

	edges := []graphs.Edge{
		{NodeA: nodeI, NodeB: nodeJ},
		{NodeA: nodeK, NodeB: nodeI},
		{NodeA: nodeM, NodeB: nodeK},
		{NodeA: nodeK, NodeB: nodeL},
		{NodeA: nodeO, NodeB: nodeN},
	}

	expectedGraph := graphs.Graph{
		nodeI: {nodeJ, nodeK},
		nodeJ: {nodeI},
		nodeK: {nodeI, nodeM, nodeL},
		nodeM: {nodeK},
		nodeL: {nodeK},
		nodeO: {nodeN},
		nodeN: {nodeO},
	}

	actual := graphs.BuildUndirectedGraphFromEdges(edges)

	if !reflect.DeepEqual(actual, expectedGraph) {
		t.Fail()
	}
}

func TestHasPathUndirected(t *testing.T) {
	nodeI := graphs.Node{Key: "i"}
	nodeJ := graphs.Node{Key: "j"}
	nodeK := graphs.Node{Key: "k"}
	nodeM := graphs.Node{Key: "m"}
	nodeL := graphs.Node{Key: "l"}
	nodeO := graphs.Node{Key: "o"}
	nodeN := graphs.Node{Key: "n"}

	edges := []graphs.Edge{
		{NodeA: nodeI, NodeB: nodeJ},
		{NodeA: nodeK, NodeB: nodeI},
		{NodeA: nodeM, NodeB: nodeK},
		{NodeA: nodeK, NodeB: nodeL},
		{NodeA: nodeO, NodeB: nodeN},
	}
	graph := graphs.BuildUndirectedGraphFromEdges(edges)

	actual := graph.HasPathUndirected(nodeO, nodeN)

	expected := true

	if actual != expected {
		t.Fail()
	}
}

func TestConnectedCountAndBiggestComponent(t *testing.T) {
	node0 := graphs.Node{Key: "0"}
	node1 := graphs.Node{Key: "1"}
	node2 := graphs.Node{Key: "2"}
	node3 := graphs.Node{Key: "3"}
	node4 := graphs.Node{Key: "4"}
	node5 := graphs.Node{Key: "5"}
	node8 := graphs.Node{Key: "8"}

	graph := graphs.Graph{
		node0: {node8, node1, node5},
		node1: {node0},
		node5: {node0, node8},
		node8: {node0, node5},
		node2: {node3, node4},
		node3: {node3, node4},
		node4: {node3, node2},
	}

	expectedComponentCount := 2
	expectedBiggestComponent := 4

	componentCount, biggestComponentSize := graph.ConnectedCountAndBiggestComponent()

	if expectedComponentCount != componentCount {
		t.Fail()
	}

	if expectedBiggestComponent != biggestComponentSize {
		t.Fail()
	}
}
