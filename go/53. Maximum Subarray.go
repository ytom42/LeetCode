// 2/11
package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

type TestSet struct {
	Nums []int
	Exp  int
}

func main() {
	inputs := []TestSet{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, v := range inputs {
		ret := maxSubArray(v.Nums)
		fmt.Println("Your Output:", ret, "Exp:", v.Exp)
	}
	return
}
