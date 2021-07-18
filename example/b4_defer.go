package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("tmp.txt")
	defer closeFile(f)
	writeFile(f)
}

func writeFile(f *os.File) {
	fmt.Println("write")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n ", err)
		os.Exit(1)
	}
}

func createFile(p string) *os.File {
	fmt.Println("ceateing")
	if f, err := os.Create(p); err != nil {
		panic(err)
	} else {
		return f
	}
}
