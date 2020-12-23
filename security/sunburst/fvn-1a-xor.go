package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
)

func getHash(s string) uint64 {
	data := []byte(s)
	// 64-bit FNV-1a
	// https://github.com/golang/go/blob/master/src/hash/fnv/fnv.go#L62
	hash := fnv.New64a()
	hash.Write(data)
	// https://github.com/golang/go/blob/master/src/hash/fnv/fnv.go#L119
	num := hash.Sum64()
	val := uint64(6605813339339102567)
	return num ^ val
}

func main() {
	// https://docs.google.com/spreadsheets/d/1u0_Df5OMsdzZcTkBDiaAtObbIOkMa5xbeXdKk_k0vWs/edit#gid=0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prog := scanner.Text()
		hash := getHash(prog)
		fmt.Printf("%s, %d\n", prog, hash)
	}
}
