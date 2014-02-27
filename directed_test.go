package gogl

import (
	"fmt"
	. "launchpad.net/gocheck"
	"math"
	"testing"
)

var _ = fmt.Println

var d_fact = &GraphFactory{
	CreateMutableGraph: func() MutableGraph {
		return NewDirected()
	},
	CreateGraph: func(edges []Edge) Graph {
		return NewDirectedFromEdgeSet(edges)
	},
}

var _ = Suite(&MutableGraphSuite{
	Factory: d_fact,
})

var _ = Suite(&GraphSuite{
	Factory: d_fact,
})

func TestDensity(t *testing.T) {
	g := NewDirected()
	var density float64

	if !math.IsNaN(g.Density()) {
		t.Error("On graph initialize, Density should be NaN (divides zero by zero)).")
	}

	g.AddEdges(&BaseEdge{"foo", "bar"})

	density = g.Density()
	if density != 1 {
		t.Error("In undirected graph of V = 2 and E = 1, density should be 1; was", density)
	}

	g.AddEdges(&BaseEdge{"baz", "qux"})

	density = g.Density()
	if density != float64(1)/float64(3) {
		t.Error("In undirected graph of V = 4 and E = 2, density should be 0.3333; was", density)
	}
}