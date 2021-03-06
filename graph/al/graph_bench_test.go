package al

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/sdboyer/gocheck"
	. "github.com/sdboyer/gogl"
)

// TODO reimplement with specs
//func SetUpBenchmarksFromBuilder(b GraphBuilder) bool {
//Suite(&GraphBenchSuite{b: b})

//return true
//}

//var _ = SetUpBenchmarksFromBuilder(BMBD)

type GraphBenchSuite struct {
	//b       GraphBuilder
	g10     Graph
	g100    Graph
	g1000   Graph
	g10000  Graph
	g100000 Graph
}

// An edge type specifically for benchmarking that encompasses all edge types.
type benchEdge struct {
	U Vertex
	V Vertex
	W float64
	L string
	P interface{}
}

func (e benchEdge) Source() Vertex {
	return e.U
}

func (e benchEdge) Target() Vertex {
	return e.V
}

func (e benchEdge) Both() (Vertex, Vertex) {
	return e.U, e.V
}

func (e benchEdge) Weight() float64 {
	return e.W
}

func (e benchEdge) Label() string {
	return e.L
}

func (e benchEdge) Data() interface{} {
	return e.P
}

func bernoulliDistributionGenerator(vertexCount uint, edgeProbability int, src rand.Source) GraphSource {
	if edgeProbability > 100 || edgeProbability < 1 {
		panic("Must designate an edge probability between 1 and 100")
	}

	if src == nil {
		src = rand.NewSource(time.Now().UnixNano())
	}

	r := rand.New(src)

	list := make([][]benchEdge, vertexCount, vertexCount)

	size := 0
	vc := int(vertexCount)
	for u := 0; u < vc; u++ {
		list[u] = make([]benchEdge, vertexCount, vertexCount)
		for v := 0; v < vc; v++ {
			// without this conditional, this loop would create a complete graph
			if v != u && // no loops
				r.Intn(100) <= edgeProbability { // create edge iff probability says so
				list[u][v] = benchEdge{U: u, V: v}
				size++
			}
		}
	}

	return &benchGraph{targetOrder: vertexCount, directed: true, list: list, size: size}
}

// A type of graph intended to serve as a controlled source of graph data for benchmarking.
type benchGraph struct {
	targetOrder   uint
	targetDensity float64
	maxDegree     uint
	minDegree     uint
	directed      bool
	list          [][]benchEdge
	size          int
}

func (g *benchGraph) Vertices(f VertexStep) {
	for v := range g.list {
		if f(v) {
			return
		}
	}
}

func (g *benchGraph) Edges(f EdgeStep) {
	for _, adj := range g.list {
		for _, e := range adj {
			if f(e) {
				return
			}
		}
	}
}

func (g *benchGraph) Size() int {
	return g.size
}

func (g *benchGraph) Order() int {
	return len(g.list)
}

// back to reality

func (s *GraphBenchSuite) SetUpSuite(c *C) {
	//src := rand.NewSource(time.Now().UnixNano())
	//s.g10 = s.b.Using(bernoulliDistributionGenerator(10, 50, src)).Graph()
	//s.g100 = s.b.Using(bernoulliDistributionGenerator(100, 50, src)).Graph()
	//s.g1000 = s.b.Using(bernoulliDistributionGenerator(1000, 50, src)).Graph()
	//s.g10000 = s.b.Using(bernoulliDistributionGenerator(10000, 50, src)).Graph()
	//	s.g100000 = s.b.Using(bernoulliDistributionGenerator(100000, 50, src)).Graph()
}

func (s *GraphBenchSuite) BenchmarkHasVertex10(c *C) {
	benchHasVertex(s.g10, c)
}

func (s *GraphBenchSuite) BenchmarkHasVertex100(c *C) {
	benchHasVertex(s.g100, c)
}

func (s *GraphBenchSuite) BenchmarkHasVertex1000(c *C) {
	benchHasVertex(s.g1000, c)
}

//func (s *GraphBenchSuite) BenchmarkHasVertex10000(c *C) {
//benchHasVertex(s.g10000, c)
//}

//func (s *GraphBenchSuite) BenchmarkHasVertex100000(c *C) {
//benchHasVertex(s.g100000, c)
//}

func benchHasVertex(g Graph, c *C) {
	for i := 0; i < c.N; i++ {
		g.HasVertex(50)
	}
}

var bgraph = Spec().Using(bernoulliDistributionGenerator(1000, 50, nil)).Create(G)

func BenchmarkHasVertex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bgraph.HasVertex(50)
	}
}

func BenchmarkVertices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bgraph.Vertices(func(v Vertex) (terminate bool) {
			return
		})

	}
}

func BenchmarkEdges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bgraph.Edges(func(e Edge) (terminate bool) {
			return
		})
	}
}
