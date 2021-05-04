package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string

	// 数字的占位大小
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	//转成数字
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)

	an = an + 5
	// 转成字符串
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
