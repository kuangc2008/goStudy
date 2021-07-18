package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2021-07-18T21:26:49+08:00")
	if e != nil {
		panic(e)
	} else {
		p(t1)
	}


	p(t.Format("3:04PM"))
	p(t.Format("Mon 01 _2 15:04:05 2006"))
	p(t.Format("2006-01-02 15:04:05 -07:00"))

	form := "3 04 pm"
	t2, e := time.Parse(form, "8 41 pm")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
