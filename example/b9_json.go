package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcd := []string {"apple", "peach", "pear"}
	slcb, _ := json.Marshal(slcd)
	fmt.Println(string(slcb))

	mapD := map[string]int {"apple":5, "letuce":7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		page: 1,
		Fruits: []string {"apple", "peac", "pear"},
	}
	res1B , _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2d := &response2{
		Page: 2,
		Fruites: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2d)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs2 := dat["strs"].([]interface{})
	str1 := strs2[0].(string)
	fmt.Println(str1)


	str := `{"page": 1, "fruites": ["apple", "peach"]}`
	//str = `{"page":2,"fruites":["apple","peach","pear"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruites[0])

	// to http or file
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int {"apple": 7, "lettce": 5}
	enc.Encode(d)

}

type response1 struct {
	page int
	Fruits []string

}

type response2 struct {
	Page int `json:"page"`
	Fruites []string `json:"fruites"`
}
