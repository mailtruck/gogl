package gogl

// A graph's behaviors are primarily a product of the constraints and
// capabilities it places on its edges. These constraints and capabilities
// determine whether certain types of operations are possible on the graph, as
// well as the efficiencies for various operations.

// gogl aims to provide a range of graph implementations that can meet
// the varying constraints and implementation needs, but still achieve optimal
// performance given those constraints.

// The Edge interface describes a connection between two vertices.
//
// Edge does not have an intrinsic opinion about directionality; gogl treats
// that as a property of the overall Graph object in which the Edge appears
// rather than a property of any individual Edge.
type Edge interface {
	Source() Vertex
	Target() Vertex
	Both() (Vertex, Vertex)
}

// A WeightedEdge is an Edge that also has a numerical weight.
type WeightedEdge interface {
	Edge
	Weight() float64
}

// A LabeledEdge is an Edge that also has a string label.
type LabeledEdge interface {
	Edge
	Label() string
}

// A PropertyEdge is an Edge that also has arbitrary property data.
type PropertyEdge interface {
	Edge
	Property() interface{}
}

// BaseEdge is a struct used to represent edges and meet the Edge interface
// requirements. It uses the standard graph notation, (U,V), for its
// contained vertex pair.
type BaseEdge struct {
	U Vertex
	V Vertex
}

func (e BaseEdge) Source() Vertex {
	return e.U
}

func (e BaseEdge) Target() Vertex {
	return e.V
}

func (e BaseEdge) Both() (Vertex, Vertex) {
	return e.U, e.V
}

// BaseWeightedEdge extends BaseEdge with weight data.
type BaseWeightedEdge struct {
	BaseEdge
	W float64
}

func (e BaseWeightedEdge) Weight() float64 {
	return e.W
}

// BaseLabeledEdge extends BaseEdge with label data.
type BaseLabeledEdge struct {
	BaseEdge
	L string
}

func (e BaseLabeledEdge) Label() string {
	return e.L
}

// BasePropertyEdge extends BaseEdge with arbitrary data.
type BasePropertyEdge struct {
	BaseEdge
	P interface{}
}

func (e BasePropertyEdge) Property() interface{} {
	return e.P
}
