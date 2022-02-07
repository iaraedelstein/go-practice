package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Invalid number of arguments:", len(args))
		os.Exit(1)
	}
	filename := args[1]
	f, e := os.Open(filename)
	if e != nil {
		fmt.Println("Error opening file", e)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}
