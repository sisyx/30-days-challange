package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		panic("Please pass the address of the two files")
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if !strings.HasPrefix(file1, "/") {
		file1 = path.Join(currentDir, file1)
	}

	if !strings.HasPrefix(file2, "/") {
		file2 = path.Join(currentDir, file2)
	}

	file1Content, err := os.ReadFile(file1)
	if err != nil {
		panic(err)
	}

	file2Content, err := os.ReadFile(file2)
	if err != nil {
		panic(err)
	}

	h1 := sha256.New()
	h1.Write([]byte(file1Content))
	bs1 := h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(file2Content))
	bs2 := h2.Sum(nil)

	isEqual := true
	for idx, b := range bs1 {
		if bs2[idx] != b {
			isEqual = false
			break
		}
	}

	if isEqual {
		fmt.Println("YAYY")
	} else {
		fmt.Println("NAHH")
	}
}
