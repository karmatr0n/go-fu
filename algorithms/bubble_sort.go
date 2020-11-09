package main

import "fmt"

func bubbleSort(a []int) {
	for {
		swapped := false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	a := []int{4, 4, 1, 2, 5, 6, 8, 0, 9, 0}
	bubbleSort(a[:])
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
