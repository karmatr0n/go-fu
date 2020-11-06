package main

import (
  "fmt"
)

func hasUniqueChars(word string) bool {
  if len(word) == 0 {
    return false 
  }
  charSet := map[uint32]bool{}
  for _, r := range word {
    v := uint32(r)
    if charSet[v] {
        return false
    }
    charSet[v] = true
  }
  return true
}

func main() {
  words := []string{ "abcde", "hello", "apple", "kite", "padle" }
  for _, word := range words {
    fmt.Println(word, ":", hasUniqueChars(word))
  }
}
