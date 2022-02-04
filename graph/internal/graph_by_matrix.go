package internal

type GraphByAdjacencyMatrix struct {
	E int
	V int
	Vertices []string
	WeightedEdges []WeightedEdge
	Edges [][]int
}

type WeightedEdge struct{
	V1 int
	V0 int
	Weight int
}

func NewGraphByAdjacencyMatrix()*GraphByAdjacencyMatrix{
	vertices:=[]string{"A","B","C","D"}
	edges:=make([][]int,4)
	edges[0]=[]int{0,1,1,1}
	edges[1]=[]int{1,0,1,1}
	edges[2]=[]int{1,1,0,1}
	edges[3]=[]int{1,1,1,0}
	return &GraphByAdjacencyMatrix{
		E:6,
		V:4,
		Vertices: vertices,
		Edges: edges,
	}
}