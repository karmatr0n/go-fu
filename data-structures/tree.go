package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

func initNode(data int) *Node {
  return &Node{left: nil, right: nil, data: data}
}

func (t *Node) Insert(data int) *Node {
  if data < t.data {
    if t.left != nil {
      t.left.Insert(data)
    } else {
      t.left = initNode(data)
    }
  } else {
    if t.right != nil {
      t.right.Insert(data)
    } else {
      t.right = initNode(data)
    }
  }
  return t
}

func (t *Node) Contains(data int) bool {
  if t.data == data {
    return true
  }

  if data < t.data {
    if t.left != nil {
      t.left.Contains(data)
    } else {
      return false
    }
  } else {
    if t.right != nil {
      t.right.Contains(data)
    } else {
      return false
    }
  }
  return false
}

func (t *Node) Traverse() {
  if t.left != nil {
    t.left.Traverse()
  }
  fmt.Println(t.data)

  if t.right != nil {
    t.right.Traverse()
  }
}


func main() {
  tree := initNode(5)
	tree.Insert(3).Insert(8).Insert(-5).Insert(200)
  fmt.Println(tree.Contains(5))
}

