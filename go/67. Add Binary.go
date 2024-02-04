// 2/4
package main

import (
	"fmt"
	"runtime"
)

func addBinary(a string, b string) string {
	idxA := len(a) - 1
	idxB := len(b) - 1
	ans := ""
	carry := 0

	for idxA >= 0 || idxB >= 0 || carry == 1 {
		n := carry
		if idxA >= 0 {
			n += int(a[idxA] - '0')
			idxA--
		}
		if idxB >= 0 {
			n += int(b[idxB] - '0')
			idxB--
		}
		carry = n / 2
		ans = string(n%2+'0') + ans
	}

	return ans
}

type TestSet struct {
	A, B     string
	Expected string
}

func main() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	inputs := []TestSet{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"0", "1", "1"},
	}

	for _, v := range inputs {
		ret := addBinary(v.A, v.B)
		fmt.Println("Expected:", v.Expected, "You Output:", ret)
	}

	fmt.Println("TotalAlloc:", memStats.TotalAlloc)

	return
}
