package gogl

/* Graph type constants. Used primarily for specs. */

// Describes the properties of a graph as a bitfield.
type GraphProperties uint16

const (
	// Edge directedness. Flags are provided for both, though gogl does not really support
	// a hybrid graph containing both directed and undirected edges. Algorithms would have
	// undefined results.
	G_UNDIRECTED = 1 << iota
	G_DIRECTED

	// Edge type. Basic (untyped edges, represented solely by the Edge interface) is the implied zero-value.
	G_BASIC
	G_LABELED
	G_WEIGHTED
	G_DATA

	// Multiplicity. Simple (no loops or multiple edges) is the implied zero-value.
	G_SIMPLE
	G_LOOPS
	G_PARALLEL

	// Mutability. Immutable is the implied zero-value.
	G_IMMUTABLE
	G_MUTABLE
	G_PERSISTENT = 1<<iota | G_MUTABLE // Persistent graphs are, kinda weirdly, both.
)

/*
TODO go back to using zero vals; see if the following can be made to work
const (
	G_UNDIRECTED = ^G_DIRECTED
	G_BASIC = ^(G_LABELED | G_WEIGHTED | G_DATA)
	G_SIMPLE = ^(G_LOOPS | G_PARALLEL)
	G_IMMUTABLE = ^G_MUTABLE
)
*/

type GraphSpec struct {
	Props  GraphProperties
	Source GraphSource
}

// Create a graph spec, which allows specification and creation of a graph through
// a fluent builder-style interface.
func Spec() GraphSpec {
	b := GraphSpec{Props: G_UNDIRECTED | G_SIMPLE | G_BASIC | G_MUTABLE}
	return b
}

// Specify that the graph should be populated from the provided source "graph".
//
// The GraphSource interface is used here because this is an ideal place
// at which to load in, for example, graph data exported into a flat file;
// a GraphSource can represent that data and only implement the minimal interface.
func (b GraphSpec) Using(g GraphSource) GraphSpec {
	b.Source = g
	return b
}

// Specify that the graph should have undirected edges.
func (b GraphSpec) Undirected() GraphSpec {
	b.Props &^= G_DIRECTED
	b.Props |= G_UNDIRECTED
	return b
}

// Specify that the graph should have directed edges (be a digraph).
func (b GraphSpec) Directed() GraphSpec {
	b.Props &^= G_UNDIRECTED
	b.Props |= G_DIRECTED
	return b
}

// Specify that the edges should be "basic" - no weights, labels, or data.
func (b GraphSpec) Basic() GraphSpec {
	b.Props &^= G_LABELED | G_WEIGHTED | G_DATA
	b.Props |= G_BASIC
	return b
}

// Specify that the edges should be labeled. See LabeledEdge
func (b GraphSpec) Labeled() GraphSpec {
	b.Props &^= G_BASIC
	b.Props |= G_LABELED
	return b
}

// Specify that the edges should be weighted. See WeightedEdge
func (b GraphSpec) Weighted() GraphSpec {
	b.Props &^= G_BASIC
	b.Props |= G_WEIGHTED
	return b
}

// Specify that the edges should contain arbitrary data. See DataEdge
func (b GraphSpec) DataEdges() GraphSpec {
	b.Props &^= G_BASIC
	b.Props |= G_DATA
	return b
}

// Specify that the graph should be simple - have no loops or multiple edges.
func (b GraphSpec) SimpleGraph() GraphSpec {
	b.Props &^= G_LOOPS | G_PARALLEL
	b.Props |= G_SIMPLE
	return b
}

// Specify that the graph is a multigraph - allows parallel edges, but no loops.
func (b GraphSpec) MultiGraph() GraphSpec {
	b.Props &^= G_SIMPLE | G_LOOPS
	b.Props |= G_PARALLEL
	return b
}

// Specify that the graph is a pseudograph - allows both loops and parallel edges.
func (b GraphSpec) PseudoGraph() GraphSpec {
	b.Props &^= G_SIMPLE
	b.Props |= G_LOOPS | G_PARALLEL
	return b
}

// Specify that the graph allows parallel edges.
func (b GraphSpec) Parallel() GraphSpec {
	b.Props &^= G_SIMPLE
	b.Props |= G_PARALLEL
	return b
}

// Specify that the graph allows loops - edges connecting a vertex to itself.
func (b GraphSpec) Loop() GraphSpec {
	b.Props &^= G_SIMPLE
	b.Props |= G_LOOPS
	return b
}

// Specify that the graph is mutable.
func (b GraphSpec) Mutable() GraphSpec {
	b.Props &^= G_IMMUTABLE | G_PERSISTENT
	b.Props |= G_MUTABLE
	return b
}

// Specify that the graph is immutable.
func (b GraphSpec) Immutable() GraphSpec {
	b.Props &^= G_PERSISTENT | G_MUTABLE // redundant, but being thorough
	b.Props |= G_IMMUTABLE
	return b
}

// Specify that the graph is persistent.
// TODO Commented out until this actually gets implemented
//func (b GraphSpec) Persistent() GraphSpec {
//b.Props &^= G_IMMUTABLE
//b.Props |= G_PERSISTENT
//return b
//}

// Creates a graph from the spec, using the provided creator function.
//
// This is just a convenience method; the creator function can always
// be called directly.
func (b GraphSpec) Create(f func(GraphSpec) Graph) Graph {
	return f(b)
}
