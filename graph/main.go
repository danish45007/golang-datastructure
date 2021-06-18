package main

import "fmt"

// adjacency list(Directed)
type Graph struct {
	vertexs []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

//Add vertex
func (g *Graph) Add(k int) {
	if !Contains(g.vertexs, k) {
		g.vertexs = append(g.vertexs, &Vertex{key: k})
	} else {
		fmt.Printf("Key %v already exist's", k)
	}
}

//Add edge
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invaild edge from %v --> to %v", from, to)
		fmt.Printf(err.Error())
	} else if Contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Edge %v already exist;s in %v --> to %v", to, from)
		fmt.Printf(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

//get vertex get the vertex pointer from the key
func (g *Graph) GetVertex(k int) *Vertex {
	for _, v := range g.vertexs {
		if v.key == k {
			return v
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, vertex := range g.vertexs {
		fmt.Printf("\n Vertex %v : ", vertex.key)
		for _, v := range vertex.adjacent {
			fmt.Printf(" %v ", v.key)
		}
		fmt.Println()
	}
}

func Contains(s []*Vertex, key int) bool {
	for _, v := range s {
		if v.key == key {
			return true
		}
	}
	return false
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.Add(i)
	}
	test.AddEdge(0, 1)
	test.AddEdge(0, 2)
	test.AddEdge(1, 4)
	test.AddEdge(0, 4)
	test.AddEdge(4, 1)
	test.Print()
}
