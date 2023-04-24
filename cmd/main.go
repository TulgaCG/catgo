package main

import (
	"fmt"
	"os"

	"github.com/TulgaCG/catgo/pkg/file"
)

const (
	bufferSize uint = 1024
)

func main() {
	fileNames := os.Args[1:]

	if err := file.Fprint(os.Stdout, bufferSize, fileNames...); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
