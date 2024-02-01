// 2/1
package main

import "fmt"

func searchInsert(nums []int, target int) int {

	min := 0
	max := len(nums) - 1

	for {
		ans := (max + min) / 2
		n := nums[ans]
		if n == target {
			return ans
		} else if n > target {
			max = ans - 1
		} else if n < target {
			min = ans + 1
		}
		if max < 0 {
			return 0
		} else if min > max {
			return min
		}
	}
}

type TestSet struct {
	Nums     []int
	Target   int
	Expected int
}

func main() {
	inputs := []TestSet{
		TestSet{[]int{1, 3, 5, 6}, 5, 2},
		TestSet{[]int{1, 3, 5, 6}, 2, 1},
		TestSet{[]int{1, 3, 5, 6}, 7, 4},
		TestSet{[]int{0}, 1, 1},
		TestSet{[]int{0}, -1, 0},
	}

	for _, v := range inputs {
		ret := searchInsert(v.Nums, v.Target)
		fmt.Println("Expected:", v.Expected, "Your Output:", ret)
	}

	return
}
