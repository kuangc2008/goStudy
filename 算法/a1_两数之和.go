package main

import "fmt"

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)

	result = twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)

	result = twoSum([]int{3, 3}, 6)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, exist := m[target-v]; exist {
			//result := make([]int, 2, 2)
			//result[0] = index
			//result[1] = i
			//return result
			return []int {index, i}
		} else {
			m[v] = i
		}
	}
	return nil
}
