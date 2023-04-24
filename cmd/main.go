package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileNames := os.Args[1:]

	readBuffer(fileNames, 1024)
}

func readBuffer(fileNames []string, bufferSize int) error {

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			return fmt.Errorf("file could not be opened: %w", err)
		}
		defer file.Close()

		buffer := make([]byte, bufferSize)

		for {
			bytesRead, err := file.Read(buffer)

			if err != nil {
				if err != io.EOF {
					return fmt.Errorf("error occured while reading file %s: %w", fileName, err)
				}

				break
			}

			fmt.Print(string(buffer[:bytesRead]))
		}
	}

	return nil
}
