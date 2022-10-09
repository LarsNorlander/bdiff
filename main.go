package main

import (
	"fmt"
	"os"
)

const (
	ErrCodeInvalidArguments  = 0
	ErrCodeFailedToReadFiles = 1
)

type Diff struct {
	Offset      int
	Left, Right byte
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("please specify two files")
		os.Exit(ErrCodeInvalidArguments)
	}
	fileNames := args[1:]

	// read file contents
	dataA, err := os.ReadFile(fileNames[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(ErrCodeFailedToReadFiles)
	}
	dataB, err := os.ReadFile(fileNames[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(ErrCodeFailedToReadFiles)
	}

	var limit int
	if len(dataA) <= len(dataB) {
		limit = len(dataA)
	} else {
		limit = len(dataB)
	}

	var diffs []Diff
	for i := 0; i < limit; i++ {
		if dataA[i] != dataB[i] {
			diffs = append(diffs, Diff{
				Offset: i,
				Left:   dataA[i],
				Right:  dataB[i],
			})
		}
	}

	for i := 0; i < len(diffs); i++ {
		fmt.Printf("%07x %02x %02x\n", diffs[i].Offset, diffs[i].Left, diffs[i].Right)
	}
}
