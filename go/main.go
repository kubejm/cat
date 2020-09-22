package main

import (
	"fmt"
	"io"
	"os"
)

func cat(file *os.File) {
	buffer := make([]byte, 8192)
	var err error

	for {
		_, err = file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading: %s\n", err)
			os.Exit(1)
		}

		_, err = os.Stdout.Write(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to stdout: %s\n", err)
			os.Exit(1)
		}
	}
}

func main() {
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening: %s\n", err)
			os.Exit(1)
		}

		cat(file)
		file.Close()
	}
}
