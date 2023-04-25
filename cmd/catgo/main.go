package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TulgaCG/catgo/pkg/file"
)

func main() {
	bufferSize := flag.Uint("buffer-size", 32, "buffer size used to read files")
	flag.Parse()

	fileNames := flag.Args()
	if err := file.Fprint(os.Stdout, *bufferSize, fileNames...); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
