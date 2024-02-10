// 2/10
package main

import (
	"fmt"
)

func jump(nums []int) int {
	lenNum := len(nums)
	ansAry := make([]int, lenNum)
	for i := 0; i < len(ansAry); i++ {
		ansAry[i] = i
	}

	for i, v := range nums {
		maxI := min(lenNum-1, i+v)
		for j := maxI; j > 0; j-- {
			if ansAry[j] > ansAry[i]+1 {
				ansAry[j] = ansAry[i] + 1
			}
		}
	}

	return ansAry[lenNum-1]
}

type TestSet struct {
	Nums []int
	Exp  int
}

func main() {
	inputs := []TestSet{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{2, 1}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{4, 1, 1, 3, 1, 1, 1}, 2},
	}

	for _, v := range inputs {
		ret := jump(v.Nums)
		fmt.Println("Your Output:", ret, "Exp:", v.Exp)
	}
	return
}
