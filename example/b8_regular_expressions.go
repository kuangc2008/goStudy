package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	if match , err := regexp.MatchString("p([a-z]+)ch", "peach"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("match::", match)
	}

	compile, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println( compile.MatchString("peach"))
		fmt.Println( compile.FindString("peach punch peach"))
		fmt.Println( compile.FindStringSubmatch("peach punch peach"))
		fmt.Println( compile.FindStringSubmatchIndex("peach punch peach"))
		fmt.Println( compile.FindStringIndex("peach punch peach"))

		fmt.Println("------------")
		fmt.Println( compile.FindAllString("peach punch peach", 2))
		fmt.Println( compile.FindAllStringSubmatchIndex("peach punch peach", -1))
		fmt.Println(compile.Match( []byte("peach")))
		fmt.Println("------------")

		r := regexp.MustCompile("p([a-z]+)ch")
		fmt.Println(r)
		fmt.Println(r.ReplaceAllString("a peach", "<fruite>"))

		in := []byte("a peach")
		out := r.ReplaceAllFunc(in, bytes.ToUpper)
		fmt.Println(string(out))
	}




}
