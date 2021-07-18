package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("wirte ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it is weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(" before noon")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("no %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
