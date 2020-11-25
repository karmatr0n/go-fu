package main
import "fmt"

type Node struct {
  id int
  edges map[int]int
}

type Graph struct {
  nodes []*Node
}

func initGraph() *Graph {
  return &Graph{
      nodes: []*Node{},
  }
}

func (g *Graph) AddNode() (id int) {
  id = len(g.nodes)
  g.nodes = append(g.nodes, &Node{ id: id, edges: make(map[int]int) })
  return
}

func (g *Graph) AddEdge(n1, n2 int, w int) {
  g.nodes[n1].edges[n2] = w
}

func (g *Graph) Neighbors(id int) []int {
  neighbors := [] int{}
  for _, node := range g.nodes {
    for edge := range node.edges {
      if node.id == id {
        neighbors = append(neighbors, edge)
      }
      if edge == id {
        neighbors = append(neighbors, node.id)
      }
    }
  }
  return neighbors
}

func (g *Graph) Nodes() []int {
  nodes := make([]int, len(g.nodes))
  for i := range g.nodes {
    nodes[i] = i
  }
  return nodes
}

func (g *Graph) Edges() [][3]int {
  edges := make([][3]int, 0, len(g.nodes))
  for i := range g.nodes {
    for k, v := range g.nodes[i].edges {
      edges = append(edges, [3]int{i, k, int(v)})
    }
  }
  return edges
}

func main() {
  g := initGraph()
  n0 := g.AddNode()
  n1 := g.AddNode()
  n2 := g.AddNode()
  g.AddEdge(n0, n1, 1)
  fmt.Println(n0)
  fmt.Println(n1)
  fmt.Println(n1)
  fmt.Println(g.Nodes())
  g.AddEdge(n0, n2, 5)
  g.Neighbors(n0)
}
