package gogl

type GraphSpec struct {
	Props GraphProperties
	Source GraphSource
}

//var b_ibd = b_ibu.Directed()

// Create a graph specification through a fluent builder-type interface.
func BuildGraph() GraphSpec {
	b := GraphSpec{Props: G_UNDIRECTED}
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
	b.Props |= G_UNDIRECTED
	return b
}

// Specify that the graph should have directed edges (be a digraph).
func (b GraphSpec) Directed() GraphSpec {
	b.Props |= G_DIRECTED
	return b
}

// Specify that the edges should be "basic" - no weights, labels, or data.
func (b GraphSpec) BasicEdges() GraphSpec {
	b.Props &^= G_LABELED | G_WEIGHTED | G_DATA
	return b
}

// Specify that the edges should be labeled. See LabeledEdge
func (b GraphSpec) LabeledEdges() GraphSpec {
	b.Props |= G_LABELED
	return b
}

// Specify that the edges should be weighted. See WeightedEdge
func (b GraphSpec) WeightedEdges() GraphSpec {
	b.Props |= G_WEIGHTED
	return b
}

// Specify that the edges should contain arbitrary data. See DataEdge
func (b GraphSpec) DataEdges() GraphSpec {
	b.Props |= G_DATA
	return b
}

// Specify that the graph should be simple - have no loops or multiple edges.
func (b GraphSpec) SimpleGraph() GraphSpec {
	b.Props &^= G_LOOPS | G_MULTI
	return b
}

// Specify that the graph is a multigraph - it allows multiple edges.
func (b GraphSpec) MultiGraph() GraphSpec {
	b.Props |= G_MULTI
	return b
}

// Specify that the graph allows loops - edges connecting a vertex to itself.
func (b GraphSpec) Loops() GraphSpec {
	b.Props |= G_LOOPS
	return b
}

// Creates a graph from the spec, using the provided creator function.
//
// This is just a convenience method; the creator function can always
// be called directly.
func (b GraphSpec) Create(f func(GraphSpec) Graph) Graph {
	return f(b)
}
