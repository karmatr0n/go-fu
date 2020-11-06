package main

import (
  "fmt"
  "net"
  "sort"
  "os"
  "time"
)

func worker(proto string, dstHost string, ports, results chan int) {
  for p := range ports {
    hostWithPort := fmt.Sprintf("%s:%d", dstHost, p)
    timeout, _ := time.ParseDuration("5s")
    conn, err := net.DialTimeout(proto, hostWithPort, timeout)
    if err != nil {
      results <- 0
      continue
    }
    conn.Close()
    results <-p
  }
}

func main() {
  dstHost := os.Args[1]

  ports := make(chan int, 100)
  results := make(chan int)
  var openports []int
  for i := 0; i < cap(ports); i++ {
    go worker("tcp", dstHost, ports, results)
  }

  go func() {
    for i := 1; i <= 1024; i++ {
      ports <- i
    }
  }()

  for  i:= 0; i < 1024; i++ {
      port := <- results
      if port != 0 {
        openports = append(openports, port)
      }
  }

  close(ports)
  close(results)
  sort.Ints(openports)
  for _, port := range openports {
    fmt.Printf("%d open\n", port)
  }
}
