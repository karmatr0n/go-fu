package main

import "fmt"

type Queue struct {
  items []string
  size int
}

func initQueue() *Queue {
  return &Queue{}
}

func (q *Queue) Enqueue(s string) {
  q.items = append(q.items, s)
  q.size++
}

func (q *Queue) IsEmpty() bool {
  return len(q.items) == 0
}

func (q *Queue) Size() int {
  return q.size
}

func (q *Queue) Front() (string, error) {
  if q.IsEmpty() {
    return "", fmt.Errorf("The queue is empty")
  }
  return q.items[0], nil
}

func (q *Queue) Dequeue() error {
  if q.IsEmpty() {
    return fmt.Errorf("The queue is empty")
  }
  q.items = append(q.items[:0], q.items[1:]...)
  q.size--
  return nil
}

func main() {
  q := initQueue()
  fmt.Printf("Queue is empty: %v\n", q.IsEmpty())
  q.Enqueue("A")
  q.Enqueue("B")
  q.Enqueue("C")
  q.Enqueue("D")
  frontItem1, _ := q.Front()
  fmt.Printf("Item at the front of the queue: %s\n", frontItem1)
  q.Dequeue()
  frontItem2, _ := q.Front()
  fmt.Printf("Item at the front of the queue: %s\n", frontItem2)
  fmt.Printf("Queue size: %d\n", q.Size())
}
