package main

import (
	trans2 "awesomeProject/server/the_way_to_go/chapter_4/what/trans"
	"fmt"
)

var twoPi = 2 * trans2.Pi // decl computes twoPi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
