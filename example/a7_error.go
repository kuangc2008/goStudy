package main

import (
	"errors"
	"fmt"
)

// error一般是最后一个参数
func f1(arg int) (int, error) {
	if arg == 42 {
		// 新建一个error
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// 实现error接口的方式，自定义error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 fail:", e)
		} else {
			fmt.Println("worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 fail:", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work "}
	}
	return arg + 3, nil
}
