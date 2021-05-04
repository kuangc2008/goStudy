package main

import (
	"fmt"
)

// 常量
const c = "C"

// 变量
var v int = 5

// struct类型
type T struct{}

// init函数
func init() {
	// initialization of package
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
	fmt.Println("Method1")
}

func Func1() { // exported function Func1
	//...
}
