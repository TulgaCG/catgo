package file_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TulgaCG/catgo/pkg/file"
)

func TestFprint(t *testing.T) {
	tests := []struct {
		id          int
		description string
		input       []string
		wantError   bool
	}{
		{
			id:          0,
			description: "single file",
			input:       []string{"test_file_1"},
		},
		{
			id:          1,
			description: "multiple files",
			input:       []string{"test_file_1", "test_file_2", "test_file_3"},
		},
		{
			id:          2,
			description: "non-existing file",
			input:       []string{"test_file_0"},
			wantError:   true,
		},
		{
			id:          3,
			description: "existing with non-existing file",
			input:       []string{"test_file_1", "test_file_0"},
			wantError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			for idx, fileName := range test.input {
				test.input[idx] = filepath.Join("./testdata/", fileName)
			}
			actual := bytes.NewBufferString("")
			err := file.Fprint(actual, 1024, test.input...)
			if test.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			p := filepath.Join("testdata", "golden", fmt.Sprintf("%d.golden", test.id))
			expected, err := os.ReadFile(p)
			require.NoError(t, err)

			require.Equal(t, string(expected), actual.String())
		})
	}
}
