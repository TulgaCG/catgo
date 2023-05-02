# Catgo
Catgo is a simple command-line interface (CLI) app written in Go. It is a project that I developed to practice my Golang skills. The aim of Catgo is to replicate the basic functionality of the standard Unix `cat` utility, which reads files sequentially and writes them to standard output (stdout).

## CLI Usage
To use catgo, you can download the binary from the releases page and run it directly in your terminal. Alternatively, you can install the catgo package and use it in your own Go projects.

Here's an example of how to use catgo from the command line:

```bash
$ catgo file1.txt file2.txt
```
This will output the contents of `file1.txt` and `file2.txt` to the terminal.

### Flags

```bash
Usage of catgo:
  -buffer-size uint
    	buffer size used to read files (default 32)
```

Here's an example of how to use the buffer-size flag:

```bash
$ catgo -buffer-size=64 file1.txt
```
This will output the contents of `file1.txt` to the terminal, using a buffer size of 64 bytes.

## Package Usage
To use the catgo package in your own Go projects, you can install it using the following command:

```bash
$ go get github.com/TulgaCG/catgo/pkg/file
```
Then, you can import the package in your Go code:

```go
import "github.com/TulgaCG/catgo/pkg/file"
```
Here's an example of how to use the catgo package in your Go code:

```go
package main

import (
	"fmt"
	"os"

	"github.com/TulgaCG/catgo/pkg/file"
)

func main(){
	err := file.Fprint(os.Stdout, 64, "file1.txt", "file2.txt")
	if err != nil {
		fmt.Println(err)
	}
}
```
This code will read the contents of `file1.txt` and `file2.txt` using a buffer size of 64 bytes, and output them to the terminal. If there is an error reading a file, it will be printed to the terminal instead.

## License
Catgo is licensed under the Apache-2.0 License. Please see the [LICENSE](https://github.com/TulgaCG/catgo/blob/main/LICENSEg) file for more information.

## Contributing
As this project is solely for my own learning purposes, I am not currently accepting contributions from external sources. However, if you have any suggestions or feedback, feel free to open an issue on the project's [GitHub page](https://github.com/TulgaCG/catgo).