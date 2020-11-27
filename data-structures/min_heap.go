package main

import (
	"fmt"
	"heap"
)

type MinHeap struct {
	*heap.Heap
}

func initMinHeap(input []int) *MinHeap {
	h := &MinHeap{
		&heap.Heap{
			Items: input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMinHeap()
	}
	return h
}

func (h *MinHeap) ExtractMin() int {
	if len(h.Items) == 0 {
		fmt.Println("No items in the heap")
	}
	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	h.Items = h.Items[:len(h.Items)-1]
	h.minHeapifyDown(0)
	return minItem
}

func (h *MinHeap) Insert(item int) *MinHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.minHeapifyUp(lastElementIndex)

	return h
}

func (h *MinHeap) buildMinHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.minHeapifyDown(i)
	}
}

func (h *MinHeap) minHeapifyDown(index int) {
	for (h.HasLeft(index) && (h.Items[index] > h.Left(index))) ||
		(h.HasRight(index) && (h.Items[index] > h.Right(index))) {

		if (h.HasLeft(index) && (h.Items[index] > h.Left(index))) &&
			(h.HasRight(index) && (h.Items[index] > h.Right(index))) {
			if h.Left(index) < h.Right(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
				index = h.GetRightIndex(index)
			}
		} else if h.HasLeft(index) && (h.Items[index] > h.Left(index)) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

func (h *MinHeap) minHeapifyUp(index int) {
	for h.HasParent(index) && (h.Parent(index) > h.Items[index]) {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func main() {
	h := initMinHeap([]int{300, 500})
	h.Insert(3001).Insert(-3).Insert(0).Insert(-2000).Insert(-42).Insert(5).Insert(-55)
	fmt.Println("Expected min value:", h.ExtractMin())

	tests := []struct {
		initial     []int
		toAdd       []int
		minExpected int
	}{
		{[]int{}, []int{4, 9, 10, 0, -4, 7}, -4},
		{[]int{}, []int{1, 2, 3, 4, 5}, 1},
		{[]int{}, []int{300, 5, 77, -8, 0, 50}, -8},
		{[]int{}, []int{-1000, 1000}, -1000},
		{[]int{}, []int{1000, -1000}, -1000},
		{[]int{41, 9, 10, 0, 7}, []int{-41}, -41},
		{[]int{0, 7, 10}, []int{1, 2, 3, 4, 5}, 0},
		{[]int{100}, []int{300, 5, 77, -8, 0, 50}, -8},
		{[]int{-3000, 0, 800}, []int{-1000, 1000}, -3000},
		{[]int{5000, 10000}, []int{1000, -1000}, -1000},
	}

	for i, test := range tests {
		h := initMinHeap(test.initial)
		for _, n := range test.toAdd {
			h.Insert(n)
		}
		if test.minExpected == h.ExtractMin() {
			fmt.Printf("Min value: %d\n", h.ExtractMin())
		} else {
			fmt.Printf("Min value incorrect: %d, expected: %d, index: %d\n", h.ExtractMin(), test.minExpected, i)
		}
	}
}
