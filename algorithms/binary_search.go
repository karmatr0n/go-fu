package main

import "fmt"

func binarySearch(list []int, item int) *int {
	low := list[0]
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == item {
			return &mid
		}

		if list[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func main() {
	list := []int{0, 3, 4, 6, 8, 9, 10, 13, 15, 18, 20}
	fmt.Println(*binarySearch(list, 18))
	fmt.Println(binarySearch(list, -1))
}
