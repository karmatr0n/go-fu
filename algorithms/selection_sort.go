package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{1, 5, 3, 6, 2, 10, -1, -10, -45}
	selectionSort(arr)
	fmt.Println(arr)
}
