package rand

import (
	"github.com/sdboyer/gogl"
	stdrand "math/rand"
)

// Generates a random graph of vertex count n with probability ρ of an edge existing between any two vertices.
//
// This produces simple graphs only - no loops, no multiple edges. Graphs can be either directed or undirected, governed
// by the appropriately named parameter.
//
// ρ must be a float64 in the range [0.0,1.0) - that is, 0.0 <= ρ < 1.0 - else, panic.
//
// If a stable graph is requested (stable == true), then the edge set presented by calling EachEdge() on the returned graph
// will be the same on every call. To provide stability, however, a memory allocation of n^2 * (int width) bytes
// is required to store the generated graph.
//
// Unstable graphs will create a new probabilistic edge set on the fly each time EachEdge(). It thus makes only minimal
// allocations, but is still CPU intensive for successive runs (and produces a different edge set). Given these
// characteristics, unstable graphs should always be used for single-use random graphs.
//
// Note that calling the Size() method on an unstable graph will create a prediction based on the Bernoulli number, but
// is not guaranteed to be exactly the same as the number of edges traversed through EachEdge().
//
// Binomial trials require a rand source. If none is provided, the builtin math lib's global rand source is used.
func BernoulliDistribution(n uint, ρ float64, directed bool, stable bool, src stdrand.Source) gogl.GraphEnumerator {
	if ρ < 0.0 || ρ >= 1.0 {
		panic("ρ must be in the range [0.0,1.0).")
	}

	var f bTrial

	if src == nil {
		f = func(ρ float64) bool {
			return stdrand.Float64() < ρ
		}
	} else {
		r := stdrand.New(src)
		f = func(ρ float64) bool {
			return r.Float64() < ρ
		}
	}

	if stable {
		return &stableBernoulliGraph{order: n, ρ: ρ, trial: f, directed: directed}
	} else {
		return unstableBernoulliGraph{order: n, ρ: ρ, trial: f, directed: directed}
	}
}

type bTrial func(ρ float64) bool

type stableBernoulliGraph struct {
	order uint
	ρ float64
	trial bTrial
	size int
	directed bool
	list [][]struct{}
}

func (g *stableBernoulliGraph) EachVertex(f gogl.VertexLambda) {
	o := int(g.order)
	for i := 0; i < o; i++ {
		if f(i) {
			return
		}
	}
}

func (g *stableBernoulliGraph) EachEdge(f gogl.EdgeLambda) {
	if g.list == nil {
		g.list = make([][]struct{}, g.order, g.order)

		// Wrapping edge lambda; records edges into the adjacency list, then passes edge along
		ff := func (e gogl.Edge) bool {
			if g.list[e.Source().(int)] == nil {
				g.list[e.Source().(int)] = make([]struct{}, g.order)
			}
			g.list[e.Source().(int)][e.Target().(int)] = struct{}{}
			g.size++
			return f(e)
		}

		if g.directed {
			bernoulliArcCreator(ff, int(g.order), g.ρ, g.trial)
		} else {
			bernoulliEdgeCreator(ff, int(g.order), g.ρ, g.trial)
		}
	} else {
		var e gogl.BaseEdge
		for u, adj := range g.list {
			for v, _ := range adj {
				e.U, e.V = u, v
				if f(e) {
					return
				}
			}
		}
	}
}

func (g *stableBernoulliGraph) Order() int {
	return int(g.order)
}

func (g *stableBernoulliGraph) Size() int {
	g.EachEdge(func (e gogl.Edge) (terminate bool) {
		return
	})
	return g.size
}

type unstableBernoulliGraph struct {
	order uint
	ρ float64
	trial bTrial
	directed bool
}

func (g unstableBernoulliGraph) EachVertex(f gogl.VertexLambda) {
	o := int(g.order)
	for i := 0; i < o; i++ {
		if f(i) {
			return
		}
	}
}

func (g unstableBernoulliGraph) EachEdge(f gogl.EdgeLambda) {
	if g.directed {
		bernoulliArcCreator(f, int(g.order), g.ρ, g.trial)
	} else {
		bernoulliEdgeCreator(f, int(g.order), g.ρ, g.trial)
	}
}

func (g unstableBernoulliGraph) Order() int {
	return int(g.order)
}

// The return value here is hogwash; as the generator is rerun with each passthrough, there is
// no guarantee the size will actually be exactly the same as the size produced by iterating EachEdge().
// It should be reasonably close...but rarely exactly correct, with the likelihood inversely proportional
// to the order of the graph.
func (g unstableBernoulliGraph) Size() int {
	var cs int

	cs = int(g.order) * (int(g.order) - 1)
	if !g.directed {
		cs = cs/2
	}

	return int(float64(cs) * (g.ρ))
}

var bernoulliEdgeCreator = func(el gogl.EdgeLambda, order int, ρ float64, cmp bTrial) {
	var e gogl.BaseEdge
	for u := 0; u < order; u++ {
		// Set target vertex to one more than current source vertex. This guarantees
		// we only evaluate each unique edge pair once, as gogl's implicit contract requires.
		for v := u + 1; v < order; v++ {
			if cmp(ρ) {
				e.U, e.V = u, v
				if el(e) {
					return
				}
			}
		}
	}
}

var bernoulliArcCreator = func(el gogl.EdgeLambda, order int, ρ float64, cmp bTrial) {
	var e gogl.BaseEdge
	for u := 0; u < order; u++ {
		for v := 0; v < order; v++ {
			if u != v && cmp(ρ) {
				e.U, e.V = u, v
				if el(e) {
					return
				}
			}
		}
	}
}