package main

import (
	"fmt"

	"github.com/nosarthur/graph"
)

func main() {
	var n1, n2 graph.Node
	n1.Nbs = []*graph.Node{&n2}

	var g graph.Graph
	g.AddNodes(&n1, &n2)

	e1 := graph.Edge{N1: &n1, N2: &n2}

	fmt.Println(e1)
}
