// 2/3
package main

import "fmt"

func plusOne(digits []int) []int {
	i := len(digits) - 1
	carry := 1

	for i >= 0 && carry != 0 {
		tmp := digits[i] + carry
		digits[i] = tmp % 10
		carry = tmp / 10
		i--
	}

	if i < 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

type TestSet struct {
	Digits   []int
	Expected []int
}

func main() {
	inputs := []TestSet{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1}, []int{2}},
	}

	for _, v := range inputs {
		ret := plusOne(v.Digits)
		fmt.Println("Your Output:", ret, "Expected:", v.Expected)
	}

	return
}
