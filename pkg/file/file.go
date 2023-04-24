package file

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Fprint(w io.Writer, bufferSize uint, names ...string) error {
	var failedFiles []string
	for _, name := range names {
		if err := readBuffer(w, name, bufferSize); err != nil {
			failedFiles = append(failedFiles, name)
		}
	}

	if len(failedFiles) != 0 {
		return fmt.Errorf("failed to read buffer from file(s): %s", strings.Join(failedFiles, ", "))
	}

	return nil
}

func readBuffer(w io.Writer, name string, bufferSize uint) error {
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", name, err)
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("failed to read file %s: %w", name, err)
			}

			break
		}

		fmt.Fprint(w, string(buffer[:bytesRead]))
	}

	return nil
}
