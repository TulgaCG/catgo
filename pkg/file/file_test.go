package file_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/TulgaCG/catgo/pkg/file"
	"github.com/stretchr/testify/require"
)

func TestFprint(t *testing.T) {
	for _, test := []struct {
		description string
		input []string
	}{
		{
			description: "single file",
			input: []string{"./test_data/test_file_1"},
		},
		{
			description: "multiple files",
			input: []string{"./test_data/test_file_1", "./test_data/test_file_2", "./test_data/test_file_3"},
		},
		{
			description: "non-existing file",
			input: []string{"./test_data/test_file_0"},
		},
		{
			description: "existing with non-existing file",
			input: []string{"./test_data/test_file_1", "./test_dat/test_file_0"},
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			actual := bytes.Buffer()
			file.Fprint(actual, 1024, test.input...)

			expected := bytes.Buffer()
			for _, fileName := range test.input {
				f, _ := os.ReadFile(fileName)
				fmt.Fprint(expected, f)
			}

			require.Equal(actual, expected)
		})
	}
}
