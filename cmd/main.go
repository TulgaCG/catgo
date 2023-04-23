package main

import (
	"fmt"
	"os"
)

func main() {
	fileNames := os.Args[1:]

	for _, name := range fileNames {
		f, _ := os.ReadFile(name)
		fmt.Println(string(f))
	}
}
