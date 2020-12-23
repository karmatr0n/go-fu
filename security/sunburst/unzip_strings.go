package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	path := "./OrionImprovementBusinessLayer.cs"

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	unzipPattern := regexp.MustCompile(`.*OrionImprovementBusinessLayer.ZipHelper.Unzip\("(?P<EncString>(\w|=|\/)*)"\)`)
	gzHeader := []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	for scanner.Scan() {
		unzipMatch := unzipPattern.FindStringSubmatch(scanner.Text())
		if len(unzipMatch) == 3 {
			gzContent, err := base64.StdEncoding.DecodeString(unzipMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			gzedBytes := append(gzHeader[:], gzContent[:]...)
			buff := bytes.NewBuffer(gzedBytes)
			gzReader, err := gzip.NewReader(buff)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(os.Stdout, gzReader)
			fmt.Printf("\n")
			if err := gzReader.Close(); err != nil {
				log.Fatal(err)
			}
		}

	}
}
