package main

// 申明
var a string

func main() {
	a = "G"
	print(a)
	f1()
}
func f1() {
	// 申明并赋值
	a := "O"
	print(a)
	f2()
}
func f2() {
	print(a)
}
