// 1/27
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool)
	tmp_i := 0

	for i, v := range nums {
		set[v] = true
		if nums[i] != nums[tmp_i] {
			tmp_i += 1
			nums[tmp_i] = nums[i]
		}
	}

	return len(set)
}

func main() {
	inputs := []int{1, 1, 2}
	ret := removeDuplicates(inputs)
	fmt.Println(ret, ":", inputs)

	inputs = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ret = removeDuplicates(inputs)
	fmt.Println(ret, ":", inputs)

	return
}
