package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user2:pass@host.com:5432/path?k=v#f"
	//s = "https://www.baidu.com:8080/aaaa/sfddf?i=2&word=adfsdf#fsafa/"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host is ", host)
	fmt.Println("port is ", port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)

}
