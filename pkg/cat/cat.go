package cat

import (
	"fmt"
	"os"
)

func Cat(names ...string) {
	for _, name := range names {
		f, err := os.ReadFile(name)
		if err != nil {
			fmt.Printf("No such file or directory: %v\n", err)
		} else {
			fmt.Printf("%s\n", string(f))
		}
	}
}
