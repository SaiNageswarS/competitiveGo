package graph

import (
	"fmt"
	"testing"
)

func TestUnweightedGraph(t *testing.T) {
	g := Graph{}
	g.AddEdge(GraphNode("1"), GraphNode("5"))

	for _, v := range []string{"3", "9", "12"} {
		g.AddEdge(GraphNode("2"), GraphNode(v))
	}

	//displaying graph
	for k := range g.Edges {
		fmt.Printf("%s ---> ", k)

		for n := g.Edges[k].Front(); n != nil; n = n.Next() {
			fmt.Printf("%s, ", n.Value.(string))
		}
		fmt.Println()
	}
}
