package main

import "fmt"

type Stack struct {
  elements []string
  size int
}

func initStack() *Stack {
  return &Stack{}
}

func (stack *Stack) Push(s string) {
  stack.elements = append(stack.elements, s)
  stack.size++
}

func (stack *Stack) Top() (string, error) {
  if stack.IsEmpty() {
    return "", fmt.Errorf("The stack is empty")
  }
  i := len(stack.elements)
  return stack.elements[i-1], nil
}

func (stack *Stack) Pop() error {
  if stack.IsEmpty() {
    return fmt.Errorf("The stack is empty")
  }
  i := len(stack.elements) - 1
  stack.elements[i] = ""
  stack.elements = stack.elements[0:i]
  stack.size--
  return nil
}

func (stack *Stack) IsEmpty() bool {
  return len(stack.elements) == 0
}

func (stack *Stack) Size() int {
  return stack.size
}

func main() {
  s := initStack()
  fmt.Printf("Stack empty: %v\n", s.IsEmpty())
  s.Push("A")
  s.Push("B")
  s.Push("C")
  s.Push("D")
  topElem, _ :=s.Top()
  fmt.Printf("Stack top element: %s\n", topElem)
  fmt.Printf("Size of stack: %d\n", s.Size())
  fmt.Printf("Stack empty: %v\n", s.IsEmpty())
  s.Pop()
  fmt.Printf("Size of stack: %d\n", s.Size())
}

