package main

import "fmt"

func sum(list []int) (int) {
  list_size := len(list)
  if list_size == 0 {
    return 0
  } else if list_size == 1 {
    return list[0]
  } else {
    return list[list_size-1] + sum(list[:list_size-1])
  }
}

func main() {
  list := []int{1,2,4,5}
  fmt.Println(sum(list))
}
