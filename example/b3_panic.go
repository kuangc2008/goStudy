package main

import (
	"fmt"
	"os"
)

func main() {
	panic("a problem")

	_, err := os.Create("~/tmp/file121212121")
	if err != nil {
		fmt.Println("aaaaa")
		panic(err)
	} else {
		fmt.Println("bbbbb")
	}
}
