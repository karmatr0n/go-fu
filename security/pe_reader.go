package main

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ioReader(file string) io.ReaderAt {
	r, err := os.Open(file)
	check(err)
	return r
}

func hasPECharacteristics(fc uint16, c uint16) bool {
	if (fc & c) == c {
		return true
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pe_parser <pe_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file := ioReader(filePath)
	f, err := pe.NewFile(file)
	check(err)

	var sizeofOptionalHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
	var sizeofOptionalHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))
	var dosheader [96]byte
	var sign [4]byte
	var base int64

	file.ReadAt(dosheader[0:], 0)
	if dosheader[0] == 'M' && dosheader[1] == 'Z' {
		signoff := int64(binary.LittleEndian.Uint32(dosheader[0x3c:]))
		file.ReadAt(sign[:], signoff)
		if !(sign[0] == 'P' && sign[1] == 'E' && sign[2] == 0 && sign[3] == 0) {
			fmt.Printf("Invalid PE File Format.\n")
		}
		base = signoff + 4
	} else {
		base = int64(0)
	}

	sr := io.NewSectionReader(file, 0, 1<<63-1)
	sr.Seek(base, os.SEEK_SET)
	binary.Read(sr, binary.LittleEndian, &f.FileHeader)

	var optHeader32 pe.OptionalHeader32
	var optHeader64 pe.OptionalHeader64
	var x86_x64 string
	var magicNumber uint16

	switch f.FileHeader.SizeOfOptionalHeader {
	case sizeofOptionalHeader32:
		binary.Read(sr, binary.LittleEndian, &optHeader32)
		if optHeader32.Magic != 0x10b {
			fmt.Printf("pe32 optional header has unexpected Magic of 0x%x", optHeader32.Magic)
		}
		magicNumber = optHeader32.Magic
		x86_x64 = "x86"
	case sizeofOptionalHeader64:
		binary.Read(sr, binary.LittleEndian, &optHeader64)
		if optHeader64.Magic != 0x20b {
			fmt.Printf("pe32 optional header has unexpected Magic of 0x%x", optHeader64.Magic)
		}
		magicNumber = optHeader32.Magic
		x86_x64 = "x64"
	}

	isDLL := hasPECharacteristics(f.Characteristics, uint16(0x2000))
	isSYS := hasPECharacteristics(f.Characteristics, uint16(0x1000))
	f.Close()

	fmt.Printf("OptionalHeader: %#x\n", f.OptionalHeader)
	fmt.Printf("DLL File: %t\n", isDLL)
	fmt.Printf("SYS File: %t\n", isSYS)
	fmt.Printf("Base: %d\n", base)
	fmt.Printf("File type: %c%c\n", sign[0], sign[1])
	fmt.Printf("dosheader[0]: %c\n", dosheader[0])
	fmt.Printf("dosheader[1]: %c\n", dosheader[1])
	fmt.Printf("MagicNumber: %#x (%s)\n", magicNumber, x86_x64)
}
