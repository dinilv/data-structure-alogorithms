package internal

type GraphByAdjacencyList struct {
	V int
	Vertices []*Vertex
}
type Vertex struct {
	Data int
	Next * AdjacencyNode
	Last * AdjacencyNode
}

type AdjacencyNode struct {
	Vertex int
	Next *AdjacencyNode
}

func (graph *GraphByAdjacencyList)Print (){

}
func (graph *GraphByAdjacencyList)Connect (){

}

func NewGraphByAdjacencyList(){
}


