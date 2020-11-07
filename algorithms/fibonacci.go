package main

import "fmt"

func fibonacci(n int) int {
   a, b := 0, 1
   for i := 0; i < n; i++ {
     a, b = b, b+a
   }
   return a
}

func main() {
  max := 15
  for n := 0; n < max; n++ {
    result := fibonacci(n)
    fmt.Printf("Fibonacci %d = %d\n", n, result)
  }
}
