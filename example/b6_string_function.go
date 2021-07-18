package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println


func main() {
	p("Contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("HasPrefix", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "e"))
	p("Index", s.Index("test", "e"))
	p("Join", s.Join([]string{"a", "b"}, "-"))
	p("repeat", s.Repeat("a", 5))
	p("Replace1", s.Replace("foo", "o", "x", -1))
	p("Replace2", s.Replace("foo", "o", "x", 1))
	p("InSplitdex", s.Split("a-b-c-d-e", "-"))
	p("Split", s.ToLower("TESt"))
	p("ToUpper", s.ToUpper("test"))
	p("Len", len("Hello"))
	p("char", "Hello"[1])
}
