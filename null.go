package gogl

import (
	"math"
)

// The null graph is a graph without any edges or vertices. It implements all possible (non-mutable) graph interfaces.
//
// In effect, it is the zero-value of all possible graph types.
const NullGraph = nullGraph(false)

type nullGraph bool

var _ Graph = nullGraph(false)
var _ Digraph = nullGraph(false)
var _ SimpleGraph = nullGraph(false)
var _ WeightedGraph = nullGraph(false)
var _ LabeledGraph = nullGraph(false)
var _ DataGraph = nullGraph(false)

func (g nullGraph) Vertices(f VertexStep)                   {}
func (g nullGraph) Edges(f EdgeStep)                       {}
func (g nullGraph) Arcs(f ArcStep)                         {}
func (g nullGraph) IncidentTo(Vertex, EdgeStep)       {}
func (g nullGraph) ArcsFrom(Vertex, ArcStep)               {}
func (g nullGraph) PredecessorsOf(Vertex, VertexStep)      {}
func (g nullGraph) ArcsTo(Vertex, ArcStep)                 {}
func (g nullGraph) SuccessorsOf(Vertex, VertexStep)        {}
func (g nullGraph) AdjacentTo(start Vertex, f VertexStep) {}

func (g nullGraph) HasVertex(v Vertex) bool {
	return false
}

func (g nullGraph) InDegreeOf(Vertex) (degree int, exists bool) {
	return 0, false
}

func (g nullGraph) OutDegreeOf(Vertex) (degree int, exists bool) {
	return 0, false
}

func (g nullGraph) DegreeOf(Vertex) (degree int, exists bool) {
	return 0, false
}

func (g nullGraph) HasEdge(e Edge) bool {
	return false
}

func (g nullGraph) HasArc(e Arc) bool {
	return false
}

func (g nullGraph) HasWeightedEdge(e WeightedEdge) bool {
	return false
}

func (g nullGraph) HasLabeledEdge(e LabeledEdge) bool {
	return false
}

func (g nullGraph) HasDataEdge(e DataEdge) bool {
	return false
}

func (g nullGraph) Density() float64 {
	return math.NaN()
}

func (g nullGraph) Transpose() Digraph {
	return g
}
