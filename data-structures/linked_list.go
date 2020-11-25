package main

import "fmt"

type node struct {
	data int
	next *node
}

type singlyList struct {
	len  int
	head *node
}

func initList() *singlyList {
	return &singlyList{}
}

func (s *singlyList) AddFront(data int) {
	n := &node{data: data}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.len++
}

func (s *singlyList) AddBack(data int) {
	n := &node{data: data}
	if s.head == nil {
		s.head = n
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	s.len++
}

func (s *singlyList) PrintNodes() {
	if s.head == nil {
		fmt.Errorf("The list is empty")
	}
	current := s.head
	for current != nil {
		fmt.Printf("%d\n", current.data)
		current = current.next
	}
}

func (s *singlyList) Size() int {
	return s.len
}

func (s *singlyList) Head() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("The list is empty")
	}
	return s.head.data, nil
}

func (s *singlyList) DelFront() {
	if s.head == nil {
		fmt.Errorf("The list is empty")
	}
	s.head = s.head.next
	s.len--
}

func (s *singlyList) DelBack() {
	if s.head == nil {
		fmt.Errorf("The list is empty")
	}
	var prev *node
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
}

func main() {
	singlyList := initList()
	singlyList.AddFront(1)
	singlyList.AddFront(2)
	singlyList.AddFront(3)
	singlyList.AddBack(0)
	singlyList.AddBack(-1)
	singlyList.AddBack(-2)
	singlyList.AddBack(-3)
	singlyList.DelFront()
	singlyList.DelBack()
	head, err := singlyList.Head()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Head node: %d\n", head)
	fmt.Printf("Total of nodes: %d\n", singlyList.Size())
	fmt.Printf("Total of nodes: %d\n", singlyList.Size())
	singlyList.PrintNodes()
}
