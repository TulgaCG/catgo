package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TulgaCG/catgo/pkg/file"
)

const (
	defaultBufferSize uint = 32
)

func main() {
	bufferSize := flag.Uint("buffer-size", defaultBufferSize, "buffer size used to read files")
	flag.Parse()

	fileNames := flag.Args()
	if err := file.Fprint(os.Stdout, *bufferSize, fileNames...); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
