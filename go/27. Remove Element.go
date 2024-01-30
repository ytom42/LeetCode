// 1/29
package main

import "fmt"

type IntSet struct {
	Nums []int
	Val  int
}

func removeElement(nums []int, val int) int {
	var k int

	for _, v := range nums {
		if val != v {
			nums[k] = v
			k++
		}
	}

	return k
}

func main() {
	inputs := []IntSet{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
	}

	for _, v := range inputs {
		fmt.Println(v)
		ret := removeElement(v.Nums, v.Val)
		fmt.Println(v, ret)
		fmt.Println("---")
	}

	return
}
