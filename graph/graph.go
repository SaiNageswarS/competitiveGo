package graph

import (
	"container/list"
)

type GraphNode interface{}

type Graph struct {
	Edges map[GraphNode]*list.List
}

func (g *Graph) AddEdge(u, v GraphNode) {
	if g.Edges == nil {
		g.Edges = make(map[GraphNode]*list.List)
	}

	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = list.New()
	}

	g.Edges[u].PushFront(v)
}

func (g *Graph) RemoveEdgeNode(u GraphNode, v *list.Element) {
	if _, ok := g.Edges[u]; ok && v != nil {
		g.Edges[u].Remove(v)
	}
}

func (g *Graph) RemoveEdge(u, v GraphNode) {
	for vNode := g.Edges[u].Front(); vNode != nil; vNode = vNode.Next() {
		if vNode.Value.(GraphNode) == v {
			g.RemoveEdgeNode(u, vNode)
			return
		}
	}
}
